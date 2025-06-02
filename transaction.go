package switch_payment_gateway_golang

import (
	"fmt"
	"net/url"
)

type Transaction struct {
	EntityID    string  `json:"entityId,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Currency    string  `json:"currency,omitempty"`
	PaymentType string  `json:"paymentType,omitempty"`
	Integrity   bool    `json:"integrity,omitempty"`
}

func (t *Transaction) URLValues() (data url.Values) {
	data = url.Values{}
	data.Set("entityId", url.QueryEscape(t.EntityID))
	data.Set("amount", url.QueryEscape(fmt.Sprintf("%2.f", t.Amount)))
	data.Set("currency", url.QueryEscape(t.Currency))
	data.Set("paymentType", url.QueryEscape(t.PaymentType))
	data.Set("integrity", url.QueryEscape(fmt.Sprintf("%t", t.Integrity)))
	return
}
