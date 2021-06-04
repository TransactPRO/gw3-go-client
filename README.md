# Transact Pro Gateway v3 Go package

This package provide ability to make requests to Transact Pro Gateway API v3.

## Installation

```bash
go get github.com/TransactPRO/gw3-go-client
```

## Documentation
This `README` provide introduction to the library usage.

### Supported operations
- Transactions
  - SMS
  - DMS HOLD
  - DMS CHARGE
  - CANCEL
  - MOTO SMS
  - MOTO DMS
  - CREDIT
  - P2P
  - B2P
  - INIT RECURRENT DMS
  - RECURRENT DMS
  - INIT RECURRENT SMS
  - RECURRENT SMS
  - REFUND
  - REVERSAL

- Information
  - HISTORY
  - RECURRENTS
  - REFUNDS
  - RESULT
  - STATUS
  - LIMITS

- Verifications
  - Verify card 3-D Secure enrollment
  - Complete card verification

- Tokenization
  - Create payment data token

- Callback processing
  - verify callback data sign

- Reporting
  - Get transactions report in CSV format

### Basic usage
```go
    // Setup your credentials for authorized requests
    ObjectGUID := "someObjectGUID" // Your GUID from Transact Pro
    SecKey := "someSecretKey" // Your API secret key

    // Setup new Gateway Client
    gateCli, gateCliErr := tprogateway.NewGatewayClient(ObjectGUID, SecKey)
    if gateCliErr != nil {
        log.Fatal(gateCliErr)
    }
    gateCli.API.BaseURI = "https://<Gateway URL>"

    // Prepare operation builder to handle your operations
    specOpsBuilder :=  gateCli.OperationBuilder()

    // Now, define your special operation for processing
    order := specOpsBuilder.NewSms()

    // Set transaction data
    order.GeneralData.OrderData.OrderDescription = "Operation Single-Message Transactions"
    order.GeneralData.CustomerData.Email = "some@email.com"
    order.PaymentMethod.Pan = "1111111111111111"
    order.PaymentMethod.ExpMmYy = "10/60"
    order.PaymentMethod.Cvv = "123"
    order.Money.Amount = 1500
    order.Money.Currency = "USD"
    order.System.UserIP = "199.99.99.1"
    order.System.XForwardedFor = "199.99.99.1"

    // Now process the operation
    opResp, opErr := gateCli.NewRequest(order)
    if opErr != nil {
        log.Fatal(opErr)
    }

    parsedResponse, parsingError := order.ParseResponse(opResp)
    if parsingError != nil {
        log.Fatal(parsingError)
    }

    if parsedResponse.Error.Code != structures.ErrorCode(0) {
        log.Println(parsedResponse.Error.Message)
    } else if parsedResponse.Gateway.RedirectURL != nil {
        // Redirect a user to received URL
    }
```

### Card verification

```go
// set card verification init mode for a payment
payment.CommandData.CardVerificationMode = structures.CardVerificationModeInit

// complete card verification
request := specOpsBuilder.NewVerifyCard()
request.VerifyCardData.GWTransactionID = initialPaymentGatewayId
response := gateCli.NewRequest(request)
if response.StatusCode == http.StatusOK {
    log.Println("SUCCESS")
} else {
    log.Println("FAILURE")
}

// set card verification verify mode for subsequent payments
newPayment.CommandData.CardVerificationMode = structures.CardVerificationModeVerify
```

### Payment data tokenization

```go
// option 1: create a payment with flag to save payment data
payment.CommandData.PaymentMethodDataSource = structures.DataSourceSaveToGateway

// option 2: send "create token" request with payment data
operation = specOpsBuilder.NewCreateToken();
operation.PaymentMethod.Pan = "1111111111111111"
operation.PaymentMethod.ExpMmYy = "10/60"
operation.PaymentMethod.CardholderName = "John Doe"
operation.Money.Currency = "EUR"
gateCli.NewRequest(operation)

// send a payment with flag to load payment data by token
newPayment.CommandData.PaymentMethodDataSource = structures.DataSourceUseGatewaySavedCardholderInitiated
newPayment.CommandData.PaymentMethodDataToken = "<initial gateway-transaction-id>"

// execute the request and parse the response
if parsedResponse.Error.Code == structures.EecAcquirerSoftDecline && parsedResponse.Gateway.RedirectURL != nil {
    // Redirect a user to received URL
}
```

### Callback validation

```go
// verify data digest
responseDigest, err := NewResponseDigest(signFromPost)
responseDigest.OriginalURI = paymentResponse.Digest.URI
responseDigest.OriginalCnonce = paymentResponse.Digest.Cnonce
responseDigest.Body = []byte(jsonFromPost)
verifyErr := responseDigest.Verify("object-guid", "secret-key")

// parse callback data as a payment response
var parsedResult CallbackResult
parsingErr := json.Unmarshal(responseDigest.Body, &parsedResult)
```

### Transactions report loading

```go
operation := specOpsBuilder.NewReport()
operation.DateCreatedFrom = structures.Time(time.Now().UTC().Add(-86400 * time.Second))
operation.DateFinishedTo = structures.Time(time.Now().UTC())

opResp, opErr := gateCli.NewRequest(operation)
if opErr != nil {
    log.Fatal(opErr)
}

report, parsingErr := operation.ParseResponse(opResp)
if parsingErr != nil {
    log.Fatal(parsingErr)
}

report := report(gateCli, specOpsBuilder)
log.Println(report.Headers)
iterationErr := report.Iterate(func(row map[string]string) bool {
    log.Println(row)
    return true
})

if iterationErr != nil {
    log.Fatal(iterationErr)
}
```

### Customization

If you need to load an HTML form from Gateway instead of cardholder browser redirect, a special operation type may be used:

```go
operation, err := specOpsBuilder.NewRetrieveForm(parsedPaymentResponse)
if err != nil {
    log.Fatal(iterationErr)
}

opResp, opErr := gateCli.NewRequest(operation)
if opErr != nil {
    log.Fatal(opErr)
}

log.Println(string(opResp.Payload))
```

## About

### Requirements

- This library works with Go 1.12 or above.

### Submit bugs and feature requests
Bugs and feature request are tracked on [GitHub](https://github.com/TransactPRO/gw3-go-client/issues)

### How to run unit tests by executing command in terminal:
```bash
$: go test ./...
```

### License
This library is licensed under the MIT License - see the `LICENSE` file for details.
