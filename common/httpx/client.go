package httpx

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cdumange/gocess"
)

type Doer interface {
	Do(ctx context.Context, req *http.Request) (*http.Response, error)
}

type Client struct {
	http     Doer
	presteps gocess.Process[*http.Request]
}

func NewClient(timeout time.Duration) *Client {
	c := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       timeout,
	}

	return &Client{
		http:     &httpclient{&c},
		presteps: gocess.NewGoProcess[*http.Request](),
	}
}

// Do executes configured presteps before Doing the request
func (c *Client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	r, err := c.presteps.Execute(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error while executing presteps : %w", err)
	}
	return c.http.Do(ctx, r)
}

type httpclient struct {
	c *http.Client
}

func (h *httpclient) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	return h.c.Do(req.WithContext(ctx))
}
