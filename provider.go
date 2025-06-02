package switch_payment_gateway_golang

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	TestHost = "eu-test.oppwa.com"
	LiveHost = ""

	DefaultHost = TestHost
)

const (
	PrepareTheCheckout = "/v1/checkouts"
)

type Provider struct {
	Host string

	HTTPClient interface {
		Do(req *http.Request) (*http.Response, error)
	}

	Authorization interface {
		Set(req *http.Request)
	}
}

func (p *Provider) PrepareCheckout(ctx context.Context, tx Transaction) (c Checkout, err error) {
	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost,
		fmt.Sprintf("https://%s%s", p.Host, PrepareTheCheckout),
		strings.NewReader(tx.URLValues().Encode()),
	)
	if err != nil {
		return c, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	p.Authorization.Set(req)

	res, err := p.HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return c, fmt.Errorf("send request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c, fmt.Errorf("read response: %w", err)
	}

	if err := json.Unmarshal(body, &c); err != nil {
		return c, fmt.Errorf("unmarshal response: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return c, fmt.Errorf("status code %d", res.StatusCode)
	}

	return c, nil
}

func (p *Provider) GetCheckout(ctx context.Context, id, entityID string) (c Checkout, err error) {
	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet,
		fmt.Sprintf("https://%s%s/%s/payment?entityId=%s", p.Host, PrepareTheCheckout, id, entityID),
		nil)
	if err != nil {
		return c, fmt.Errorf("create request: %w", err)
	}

	p.Authorization.Set(req)

	res, err := p.HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return c, fmt.Errorf("send request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c, fmt.Errorf("read response: %w", err)
	}

	if err := json.Unmarshal(body, &c); err != nil {
		return c, fmt.Errorf("unmarshal response: %w", err)
	}

	//if res.StatusCode != http.StatusOK {
	//	return c, fmt.Errorf("status code %d", res.StatusCode)
	//}

	return c, nil
}
