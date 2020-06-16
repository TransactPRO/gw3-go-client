package reporting

import (
	"testing"

	"github.com/TransactPRO/gw3-go-client/structures"
	"github.com/stretchr/testify/assert"
)

func TestParseCsvResponseFullRead(t *testing.T) {
	expected := []map[string]string{
		{
			"aaa": "1",
			"bbb": "2",
			"ccc": "3",
		},
		{
			"aaa": "xxx",
			"bbb": "yyyy",
			"ccc": "zzz",
		},
	}

	body := "aaa,bbb,ccc\n" +
		"1,2,3\n" +
		"xxx,yyyy,zzz\n" +
		"\n"

	response := structures.NewGatewayResponse(nil, []byte(body))
	parsedResponse, err := NewReportAssembly().ParseResponse(response)
	assert.NoError(t, err)

	var cnt int
	iterErr := parsedResponse.Iterate(func(row map[string]string) bool {
		assert.Truef(t, cnt < len(expected), "Extra value for %d in row %#v", cnt+1, row)
		assert.Equalf(t, expected[cnt], row, "Wrong value for %d in row %#v", cnt+1, row)

		cnt++
		return true
	})

	assert.NoError(t, iterErr)
	assert.Equal(t, len(expected), cnt, "Expected values count differ from actual")
}

func TestParseCsvResponsePartialRead(t *testing.T) {
	allowedRecords := 1

	expected := []map[string]string{
		{
			"aaa": "1",
			"bbb": "2",
			"ccc": "3",
		},
	}

	body := "aaa,bbb,ccc\n" +
		"1,2,3\n" +
		"xxx,yyyy,zzz\n" +
		"\n"

	response := structures.NewGatewayResponse(nil, []byte(body))
	parsedResponse, err := NewReportAssembly().ParseResponse(response)
	assert.NoError(t, err)

	var cnt int
	iterErr := parsedResponse.Iterate(func(row map[string]string) bool {
		assert.Truef(t, cnt < len(expected), "Extra value for %d in row %#v", cnt+1, row)
		assert.Equalf(t, expected[cnt], row, "Wrong value for %d in row %#v", cnt+1, row)

		cnt++
		return cnt < allowedRecords
	})

	assert.NoError(t, iterErr)
	assert.Equal(t, allowedRecords, cnt, "Expected values count differ from actual")
}
