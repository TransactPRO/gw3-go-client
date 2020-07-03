package structures

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"strconv"
	"strings"
	"time"
)

type (
	// QOP is a custom type for Authorization header "qop" field
	QOP uint
	// Algorithm is a custom type for Authorization header "algorithm" field
	Algorithm uint

	digest struct {
		Username  string
		URI       string
		QOP       QOP
		Algorithm Algorithm
		Cnonce    []byte
		Response  string
		Body      []byte
	}

	// RequestDigest is a functional structure for request Authorization header value creation
	RequestDigest struct {
		digest

		Secret string
	}

	// ResponseDigest is a functional structure for response Authorization header value validation
	ResponseDigest struct {
		digest

		Timestamp      int
		Snonce         []byte
		OriginalURI    string
		OriginalCnonce []byte
	}
)

// Authorization header's "qop" values
const (
	QopUnknown QOP = iota
	QopAuth
	QopAuthInt
)

var qop2string = map[QOP]string{
	QopAuth:    "auth",
	QopAuthInt: "auth-int",
}

var string2qop = map[string]QOP{
	"auth":     QopAuth,
	"auth-int": QopAuthInt,
}

// NewQOP creates QOP value from raw string or returns an error if input is unknown
func NewQOP(in string) (QOP, error) {
	if result, ok := string2qop[in]; ok {
		return result, nil
	}

	return QopUnknown, fmt.Errorf("unknown QOP value %s", in)
}

func (o QOP) String() string {
	if result, ok := qop2string[o]; ok {
		return result
	}

	return "unknown"
}

// Authorization header's "algorithm" values
const (
	AlgorithmUnknown Algorithm = iota
	AlgorithmSHA256
)

var algorithm2string = map[Algorithm]string{
	AlgorithmSHA256: "SHA-256",
}

var string2algorithm = map[string]Algorithm{
	"SHA-256": AlgorithmSHA256,
}

// NewAlgorithm creates Algorithm value from raw string or returns an error if input is unknown
func NewAlgorithm(in string) (Algorithm, error) {
	if result, ok := string2algorithm[in]; ok {
		return result, nil
	}

	return AlgorithmUnknown, fmt.Errorf("unknown algorithm %s", in)
}

func (o Algorithm) String() string {
	if result, ok := algorithm2string[o]; ok {
		return result
	}

	return "unknown"
}

// Hash returns corresponding hash.Hash creator function
func (o Algorithm) Hash() (hashCreator func() hash.Hash, err error) {
	switch o {
	case AlgorithmSHA256:
		hashCreator = sha256.New
	default:
		err = fmt.Errorf("unsupported algorithm %s", o)
	}

	return
}

func calcNonce() (nonce []byte, err error) {
	nonceRand := make([]byte, 32)
	if _, err := rand.Read(nonceRand); err != nil {
		return nil, fmt.Errorf("cannot create cnonce: %s", err)
	}

	buf := bytes.NewBuffer([]byte(fmt.Sprintf("%d:", time.Now().UTC().Unix())))
	buf.Write(nonceRand)
	nonce = buf.Bytes()

	return
}

// NewRequestDigest creates new RequestDigest structure
func NewRequestDigest(ObjectGUID, secret, uri string, body []byte) (result *RequestDigest, err error) {
	var cnonce []byte
	if cnonce, err = calcNonce(); err != nil {
		return
	}

	result = &RequestDigest{
		digest: digest{
			Username:  ObjectGUID,
			Algorithm: AlgorithmSHA256,
			QOP:       QopAuthInt,
			URI:       uri,
			Cnonce:    cnonce,
			Body:      body,
		},
		Secret: secret,
	}
	return
}

// CreateHeader creates  a string for Authorization header for a Gateway request
func (o *RequestDigest) CreateHeader() (digest string, err error) {
	var hashFunc func() hash.Hash
	if hashFunc, err = o.Algorithm.Hash(); err != nil {
		return
	}

	mac := hmac.New(hashFunc, []byte(o.Secret))
	mac.Write([]byte(o.Username))
	mac.Write(o.Cnonce)
	mac.Write([]byte(o.QOP.String()))
	mac.Write([]byte(o.URI))
	if o.QOP == QopAuthInt {
		mac.Write(o.Body)
	}
	o.Response = hex.EncodeToString(mac.Sum(nil))

	digest = fmt.Sprintf(
		"Digest username=%s, uri=\"%s\", algorithm=%s, cnonce=\"%s\", qop=%s, response=\"%s\"",
		o.Username,
		o.URI,
		o.Algorithm,
		base64.StdEncoding.EncodeToString(o.Cnonce),
		o.QOP,
		o.Response,
	)

	return
}

// NewResponseDigest parse Authorization header's content and fill ResponseDigest structure
func NewResponseDigest(authorizationHeader string) (result *ResponseDigest, err error) {
	if len(authorizationHeader) == 0 {
		err = errors.New("authorization header is missing")
		return
	}

	values := map[string]string{
		"username":  "",
		"uri":       "",
		"response":  "",
		"algorithm": "",
		"qop":       "",
		"cnonce":    "",
		"snonce":    "",
	}

	elements := strings.Split(
		strings.TrimSpace(
			strings.ReplaceAll(
				strings.ReplaceAll(authorizationHeader, "\r", ""),
				"\n", "",
			),
		),
		",",
	)
	for i := range elements {
		keyValue := strings.TrimSpace(elements[i])
		delimiterPos := strings.Index(keyValue, "=")
		if delimiterPos != -1 {
			key := keyValue[:delimiterPos]
			if i == 0 {
				key = strings.Replace(strings.ToLower(key), "digest ", "", 1)
			}

			if _, ok := values[key]; ok {
				value := keyValue[delimiterPos+1:]
				values[key] = strings.TrimSpace(strings.Trim(value, "\""))
			}
		}
	}

	result = &ResponseDigest{}
	for key, value := range values {
		if len(value) == 0 {
			err = fmt.Errorf("digest mismatch: empty value for %s", key)
			return
		}

		switch key {
		case "username":
			result.Username = value
		case "uri":
			result.URI = value
		case "response":
			result.Response = value
		case "algorithm":
			if result.Algorithm, err = NewAlgorithm(value); err != nil {
				err = fmt.Errorf("digest mismatch: format error: %s", err)
				return
			}
		case "qop":
			if result.QOP, err = NewQOP(value); err != nil {
				err = fmt.Errorf("digest mismatch: format error: %s", err)
				return
			}
		case "cnonce":
			if result.Cnonce, err = base64.StdEncoding.DecodeString(value); err != nil {
				err = fmt.Errorf("digest mismatch: corrupted value for cnonce (%s)", err)
				return
			}
		case "snonce":
			if result.Snonce, err = base64.StdEncoding.DecodeString(value); err != nil {
				err = fmt.Errorf("digest mismatch: corrupted value for snonce (%s)", err)
				return
			}

			timestampDelimiterPos := bytes.Index(result.Snonce, []byte(":"))
			if timestampDelimiterPos == -1 {
				err = errors.New("digest mismatch: corrupted value for snonce (missing timestamp)")
				return
			}

			if result.Timestamp, err = strconv.Atoi(string(result.Snonce[:timestampDelimiterPos])); err != nil {
				err = fmt.Errorf("digest mismatch: corrupted value for snonce (unexpected timestamp value: %s)", err)
				return
			}
		}
	}

	return
}

// Verify verifies that parsed response digest was made using  given GUID/secret pair.
// In addition, if set, original request's GUID, URI and cnonce will be compared to parsed values.
func (o *ResponseDigest) Verify(objectGUID, secret string) (err error) {
	if strings.ToLower(objectGUID) != strings.ToLower(o.Username) {
		return errors.New("digest mismatch: username mismatch")
	}

	if len(o.OriginalURI) > 0 && o.OriginalURI != o.URI {
		return errors.New("digest mismatch: uri mismatch")
	}

	if len(o.OriginalCnonce) > 0 && !bytes.Equal(o.OriginalCnonce, o.Cnonce) {
		return errors.New("digest mismatch: cnonce mismatch")
	}

	var hashFunc func() hash.Hash
	if hashFunc, err = o.Algorithm.Hash(); err != nil {
		return
	}

	mac := hmac.New(hashFunc, []byte(secret))
	mac.Write([]byte(o.Username))
	mac.Write(o.Cnonce)
	mac.Write(o.Snonce)
	mac.Write([]byte(o.QOP.String()))
	mac.Write([]byte(o.URI))
	if o.QOP == QopAuthInt {
		mac.Write(o.Body)
	}
	expectedDigest := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(o.Response), []byte(expectedDigest)) {
		return errors.New("digest mismatch")
	}

	return nil
}
