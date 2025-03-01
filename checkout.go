package switch_payment_gateway_golang

type Checkout struct {
	ID          string `json:"id"`
	BuildNumber string `json:"buildNumber"`
	NDC         string `json:"ndc"`
	Integrity   string `json:"integrity"`
	Timestamp   string `json:"timestamp"`

	Result struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"result"`
}
