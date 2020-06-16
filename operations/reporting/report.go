package reporting

import (
	"net/http"

	"github.com/TransactPRO/gw3-go-client/structures"
)

// ReportAssembly is default structure for transactions CSV report loading
type ReportAssembly struct {
	opHTTPData structures.OperationRequestHTTPData

	DateCreatedFrom  structures.Time `json:"dt-created-from,omitempty"`
	DateCreatedTo    structures.Time `json:"dt-created-to,omitempty"`
	DateFinishedFrom structures.Time `json:"dt-finished-from,omitempty"`
	DateFinishedTo   structures.Time `json:"dt-finished-to,omitempty"`
}

// NewReportAssembly returns new instance with prepared HTTP request data ReportAssembly
func NewReportAssembly() *ReportAssembly {
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod(http.MethodPost)
	opd.SetOperationType(structures.Report)

	return &ReportAssembly{
		opHTTPData: opd,
	}
}

// GetHTTPMethod return HTTP method which will be used for send request
func (op *ReportAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *ReportAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}

// ParseResponse parses Gateway response into corresponding data structure
func (op *ReportAssembly) ParseResponse(response *structures.GatewayResponse) (result *structures.CsvReport, err error) {
	return structures.NewCsvReport(response)
}
