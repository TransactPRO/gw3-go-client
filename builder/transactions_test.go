package builder

import (
	"testing"
)

func TestBuildSmsDataSet(t *testing.T)  {
	var ob *OperationBuilder
	var smsDataSet SMSDataSet

	dataSet := ob.SMS()

	if dataSet != smsDataSet {
		t.Error("OperationBuilder method SMS returned wrong structure must be SMSDataSet")
	}
}