# Switch Payment Client

This library is a simple client for [SwitchPayment](https://hyperpay.docs.oppwa.com/). It's currently under development,
but key features
have been implemented.

```bash
go get -u github.com/esiteltd/switch-payment-gateway-golang
```

## Usage

The following example creates a checkout.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pswitch "github.com/esiteltd/switch-payment-gateway-golang"
)

func main() {
	p := pswitch.Provider{
		Host:          pswitch.DefaultHost,
		HTTPClient:    http.DefaultClient,
		Authorization: pswitch.DefaultJWTAuthorization,
	}

	tx := pswitch.Transaction{
		EntityID:    pswitch.TestEntityID,
		Amount:      92.00,
		Currency:    pswitch.IraqiDinar,
		PaymentType: "DB",
		Integrity:   true,
	}

	checkout, err := p.PrepareCheckout(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(checkout)
}
```
