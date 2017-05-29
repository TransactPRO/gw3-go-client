package builder

import (
	"testing"
)

func TestBuildSmsDataSet(t *testing.T) {
	var ob *OperationBuilder
	var smsDataSet SMSAssembly

	smsDataSet.method = "POST"
	smsDataSet.operationType = SMS

	dataSet := ob.SMS()

	if dataSet != nil {
		t.Error("OperationBuilder method SMS returned wrong structure must be SMSAssembly")
	}
}
