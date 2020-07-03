package structures

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnrollmentUnmarshalJSON(t *testing.T) {
	examples := map[string]struct {
		value Enrollment
		str   string
	}{
		"n":   {EnrollmentNo, "no"},
		"y":   {EnrollmentYes, "yes"},
		"aaa": {EnrollmentUnknown, "unknown"},
	}

	for value, expected := range examples {
		t.Run(expected.str, func(t *testing.T) {
			jsonValue := "\"" + value + "\""

			var actual Enrollment
			err := json.Unmarshal([]byte(jsonValue), &actual)

			assert.NoError(t, err)
			assert.Equal(t, expected.value, actual)
			assert.Equal(t, expected.str, actual.String())
		})
	}
}
