package structures

import (
	"encoding/json"
	"errors"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestURLUnmarshalJSON(t *testing.T) {
	validURL, _ := url.Parse("http://google.com")

	examples := []struct {
		input    string
		expected URL
		err      error
	}{
		{"0", URL{}, errors.New("json: cannot unmarshal number into Go value of type string")},
		{"\":abracadabra\"", URL{}, errors.New("missing protocol scheme")},
		{"\"http://google.com\"", URL(*validURL), nil},
	}

	for _, testCase := range examples {
		t.Run(testCase.input, func(t *testing.T) {
			var actual URL
			err := json.Unmarshal([]byte(testCase.input), &actual)

			if testCase.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected, actual)
			} else {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), testCase.err.Error())
			}
		})
	}
}

func TestTimeMarshalJSON(t *testing.T) {
	testTime, _ := time.Parse("2006-01-02 15:04:05", "2020-06-10 08:37:22")

	testData := struct {
		A *Time `json:"a"`
		B *Time `json:"b"`
		C Time  `json:"c"`
		D Time  `json:"d"`
	}{
		B: (*Time)(&testTime),
		D: Time(testTime),
	}

	expectedJSON := `{"a":null,"b":1591778242,"c":null,"d":1591778242}`

	actual, err := json.Marshal(testData)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(actual))
}

func TestTimeUnmarshalJSON(t *testing.T) {
	validTime, _ := time.Parse("2006-01-02 15:04:05", "2020-06-08 14:00:07")

	examples := []struct {
		input    string
		expected Time
		err      error
	}{
		{"0", Time{}, errors.New("json: cannot unmarshal number into Go value of type string")},
		{"\"aaa\"", Time{}, errors.New("parsing time \"aaa\" as \"2006-01-02 15:04:05\": cannot parse \"aaa\" as \"2006\"")},
		{"\"2020-06-08 14:00:07\"", Time(validTime), nil},
	}

	for _, testCase := range examples {
		t.Run(testCase.input, func(t *testing.T) {
			var actual Time
			err := json.Unmarshal([]byte(testCase.input), &actual)

			if testCase.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected, actual)
			} else {
				assert.EqualError(t, err, testCase.err.Error())
			}
		})
	}
}
