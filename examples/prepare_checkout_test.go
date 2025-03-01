package examples_test

import (
	"net/http"
	"testing"

	pswitch "github.com/esiteltd/switch-payment-gateway-golang"
)

func TestPrepareCheckout(t *testing.T) {
	p := pswitch.Provider{
		Host:          pswitch.DefaultHost,
		HTTPClient:    http.DefaultClient,
		Authorization: pswitch.DefaultJWTAuthorization,
	}

	tx := pswitch.Transaction{
		EntityID:    "8a8294174d0595bb014d05d829cb01cd",
		Amount:      92.00,
		Currency:    pswitch.IraqiDinar,
		PaymentType: "DB",
		Integrity:   true,
	}

	checkout, err := p.PrepareCheckout(t.Context(), tx)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("checkout %+v", checkout)
}
