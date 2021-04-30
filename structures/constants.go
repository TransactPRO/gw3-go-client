package structures

type (
	// Status represents transaction status in the Gateway
	Status int
	// ErrorCode represents Gateway error
	ErrorCode int
	// Enrollment represents card's 3-D Secure enrollment
	Enrollment int
)

// Gateway transaction statuses
const (
	StatusInit Status = iota + 1
	StatusSent2Bank
	StatusDmsHoldOK
	StatusDmsHoldFailed
	StatusSmsFailed
	StatusDmsChargeFailed
	StatusSuccess
	StatusExpired
	StatusHoldExpired
	StatusRefundFailed Status = iota + 2
	StatusRefundPending
	StatusRefundSuccess
	StatusCardholderOnSite
	StatusDmsCanceled
	StatusDmsCancelFailed
	StatusReversed
	StatusInputValidationFailed
	StatusBusinessRulesValidationFailed
	StatusTerminalGroupSelectFailed
	StatusTerminalSelectFailed
	StatusInitParamsInvalid
	StatusDeclinedByBusinessRulesAction
	StatusCallbackURLGenerated
	StatusWaitingCardFormFill
	StatusMpiURLGenerated
	StatusWaitingMpi
	StatusMpiFailed
	StatusMpiNotReachable
	StatusCardFormURLSent
	StatusMpiAuthError
	StatusAcquirerNotReachable
	StatusReversalFailed
	StatusCreditFailed
	StatusP2PFailed
	StatusB2PFailed
	StatusTokenCreated
	StatusTokenCreateFailed
)

// Gateway transaction errors
const (
	EecGeneralError ErrorCode = 1000

	EecDisabledAccount     ErrorCode = 1001
	EecDisabledTerminal    ErrorCode = 1002
	EecDisabledLegalPerson ErrorCode = 1003

	EecTimeoutTransaction ErrorCode = 1004
	EecTimeoutRedirect    ErrorCode = 1005
	EecTimeout3D          ErrorCode = 1006
	EecTimeoutAcquirer    ErrorCode = 1007
	EecTimeoutInternal    ErrorCode = 1008

	EecHsmEncode ErrorCode = 1009
	EecHsmDecode ErrorCode = 1010

	EecMerchantCountersExceeded      ErrorCode = 1011
	EecAccountCountersExceeded       ErrorCode = 1012
	EecTerminalGroupCountersExceeded ErrorCode = 1013
	EecTerminalCountersExceeded      ErrorCode = 1014
	EecHsmToken                      ErrorCode = 1015

	EecInputValidationFailed ErrorCode = 1100
	EecFailedBusinessRules   ErrorCode = 1101

	EecCardBadNumber       ErrorCode = 1102
	EecCardBadExpire       ErrorCode = 1103
	EecCardNoCVV           ErrorCode = 1104
	EecCardBadCVV          ErrorCode = 1105
	EecCardExpired         ErrorCode = 1106
	EecCardUnknownCardType ErrorCode = 1107
	Eec3DErrorMdStatus     ErrorCode = 1108
	Eec3DErrorAuth         ErrorCode = 1109
	EecCardLiabilityShift  ErrorCode = 1110
	EecCardNotValidated    ErrorCode = 1111
	Eec3DDataCorrupted     ErrorCode = 1112

	EecWrongGwUniqID               ErrorCode = 1151
	EecUnacceptableGwUniqID        ErrorCode = 1152
	EecGwUniqIDConflict            ErrorCode = 1153
	EecTransactionTypeInvalid      ErrorCode = 1154
	EecTransactionStateInvalid     ErrorCode = 1155
	EecTransactionAlreadyFinished  ErrorCode = 1156
	EecNoParentTransactionProvided ErrorCode = 1157
	EecDynamicDescriptorError      ErrorCode = 1158
	EecUcofError                   ErrorCode = 1159
	EecSuspectedFraud              ErrorCode = 1160

	EecTerminalNotFound                 ErrorCode = 1200
	EecAllTerminalCountersExceeded      ErrorCode = 1201
	EecTerminalGroupNotFound            ErrorCode = 1202
	EecAllTerminalGroupCountersExceeded ErrorCode = 1203

	EecTerminalNotSupportingMOTO                  ErrorCode = 1204
	EecTerminalNotSupportingRecurringTransactions ErrorCode = 1205

	EecDeclinedByAcquirer  ErrorCode = 1301
	EecAcquirerError       ErrorCode = 1302
	EecAcquirerSoftDecline ErrorCode = 1303

	EecInvalidFormID   ErrorCode = 1400
	EecFormUnavailable ErrorCode = 1401

	EecCardVerificationNoCardData      ErrorCode = 1500
	EecCardVerificationAlreadyVerified ErrorCode = 1501

	EecRbsInvalidOrderNumber ErrorCode = 2000
	EecRbsInvalidDescription ErrorCode = 2001
)

// 3-D Secure enrollment statuses
const (
	EnrollmentUnknown Enrollment = iota
	EnrollmentNo
	EnrollmentYes
)

var enrollment2string = map[Enrollment]string{
	EnrollmentNo:  "no",
	EnrollmentYes: "yes",
}

func (o Enrollment) String() string {
	if result, ok := enrollment2string[o]; ok {
		return result
	}

	return "unknown"
}

// UnmarshalJSON is a custom unmarshal function for 3-D Secure enrollment representation string
func (o *Enrollment) UnmarshalJSON(raw []byte) error {
	switch string(raw) {
	case "\"n\"":
		*o = EnrollmentNo
	case "\"y\"":
		*o = EnrollmentYes
	default:
		*o = EnrollmentUnknown
	}

	return nil
}
