package switch_payment_gateway_golang

import (
	"fmt"
	"net/http"
)

const TestJWT = "OGE4Mjk0MTc0ZDA1OTViYjAxNGQwNWQ4MjllNzAxZDF8OVRuSlBjMm45aA=="

var DefaultJWTAuthorization = NewJWTAuthorization(TestJWT)

type JWTAuthorization struct {
	Token string
}

func NewJWTAuthorization(token string) *JWTAuthorization {
	return &JWTAuthorization{token}
}

func (j *JWTAuthorization) Set(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", j.Token))
}
