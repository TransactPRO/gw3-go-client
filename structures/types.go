package structures

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type (
	// URL is a standard url.URL with custom UnmarshalJSON
	URL url.URL

	// Time is a standard time.Time with custom UnmarshalJSON
	Time time.Time
)

// UnmarshalJSON is a custom unmarshal function for URL type
func (o *URL) UnmarshalJSON(raw []byte) error {
	var urlString string
	if err := json.Unmarshal(raw, &urlString); err != nil {
		return err
	}

	parsed, err := url.Parse(urlString)
	if err != nil {
		return err
	}

	*o = URL(*parsed)
	return nil
}

// MarshalJSON is a custom marshal function for Time type.
// Is used to convert Time to Unix timestamp.
func (o Time) MarshalJSON() ([]byte, error) {
	timestamp := time.Time(o).UTC().Unix()
	if timestamp <= 0 {
		return []byte("null"), nil
	}

	timestampStr := fmt.Sprintf("%d", time.Time(o).UTC().Unix())
	return []byte(timestampStr), nil
}

// UnmarshalJSON is a custom unmarshal function for Time type.
// Is used to create Time from Unix timestamp.
func (o *Time) UnmarshalJSON(raw []byte) error {
	var timeString string
	if err := json.Unmarshal(raw, &timeString); err != nil {
		return err
	}

	parsed, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		return err
	}

	*o = Time(parsed)
	return nil
}
