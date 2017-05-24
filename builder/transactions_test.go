package builder

import (
	"testing"
)

func TestRequestHTTPData_GetHTTPMethod(t *testing.T) {
	var req requestHTTPData

	req.method = "POST"

	methodHTTP := req.GetHTTPMethod()
	if methodHTTP == "" {
		t.Error("In HTTP request methods is empty")
	}
}

func TestRequestHTTPData_GetRoutePath(t *testing.T) {
	var req requestHTTPData

	req.routePath = "/some/path"
	path := req.GetRoutePath()
	if path == "" {
		t.Error("In HTTP request path is empty")
	}
}

func TestBuildSmsDataSet(t *testing.T) {
	var ob *OperationBuilder
	var smsDataSet SMSAssembly

	dataSet := ob.SMS()

	if dataSet != smsDataSet {
		t.Error("OperationBuilder method SMS returned wrong structure must be SMSAssembly")
	}
}
