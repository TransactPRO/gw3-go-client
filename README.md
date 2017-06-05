# Transact Pro Gateway v3 Go package

This package provide ability to make requests to Transact Pro Gateway API v3.

## Documentation
This `README` provide introduction to the library usage.

Read more about all gateway possibilities in the [Official documentation](http://transactpro.lv/docs/gw3-api-doc.pdf)

### Supported operations
- Transactions
  - SMS
  - DMS HOLD
  - DMS CHARGE
  - CANCEL
  - MOTO SMS
  - MOTO DMS
  - RECURRENT DMS
  - RECURRENT SMS
  - REFUND
  - REVERSAL

- Information
  - HISTORY
  - RECURRENTS
  - REFUNDS
  - RESULT
  - STATUS

#### Basic usage
```go
    // Setup your credentials for authorized requests
    AccID := 42 // Your account ID form Transact Pro
    SecKey := "someSecretKey" // Your API secret key

    // Setup new Gateway Client
    gateCli, gateCliErr := tprogateway.NewGatewayClient(AccID, SecKey)
    if gateCliErr != nil {
        log.Fatal(gateCliErr)
    }

    // Prepare operation builder to handle your operations
    specOpsBuilder :=  gateCli.OperationBuilder()

    // Now, define your special operation for processing
    sms := specOpsBuilder.NewSms()

    // Set transaction data
    sms.GeneralData.OrderData.OrderDescription = "Operation Single-Message Transactions"
    sms.GeneralData.CustomerData.Email = "some@email.com"
    sms.PaymentMethod.Pan = "1111111111111111"
    sms.PaymentMethod.ExpMmYy = "10/60"
    sms.PaymentMethod.Cvv = "123"
    sms.Money.Amount = 1500
    sms.Money.Currency = "USD"
    sms.System.UserIP = "199.99.99.1"
    sms.System.XForwardedFor = "199.99.99.1"

    // Now process the operation
    opResp, opErr := gc.NewRequest(sms)
    if opErr != nil {
        log.Printf(opErr)
    }
```

### Submit bugs and feature requests
Bugs and feature request are tracked on [GitHub](https://github.com/TransactPRO/gw3-go-client/issues)


### How to run unit tests by executing command in terminal:
```bash
$: go test ./...
```

### License
This library is licensed under the MIT License - see the `LICENSE` file for details.