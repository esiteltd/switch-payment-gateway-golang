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
		EntityID:    pswitch.TestEntityID,
		Amount:      92.00,
		Currency:    pswitch.IraqiDinar,
		PaymentType: "DB",
		Integrity:   true,
	}

	checkout, err := p.PrepareCheckout(t.Context(), tx)
	if err != nil {
		t.Fatal(err, checkout)
	}

	t.Logf("checkout %+v", checkout)

	checkout, err = p.GetCheckout(t.Context(), checkout.ID, tx.EntityID)
	if err != nil {
		t.Fatal(err, checkout)
	}

	t.Logf("checkout %+v", checkout)
}
