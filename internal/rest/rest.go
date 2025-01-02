package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// LinkResponse object returning links to next
type LinkResponse struct {
	Next *string `json:"next,omitempty"`
}

// ListAccountsResponse object
type ListAccountsResponse struct {
	Accounts []Account    `json:"accounts"`
	Links    LinkResponse `json:"links"`
}

// Client is an interface of HTTP client
type Client interface {
	// ListAccounts returns a list of accounts object
	ListAccounts(ctx context.Context) (ListAccountsResponse, error)
}

type client struct {
	timeout time.Duration
	url     string
}

func (c client) ListAccounts(ctx context.Context) (ListAccountsResponse, error) {
	url := fmt.Sprintf("%s/accounts", c.url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ListAccountsResponse{}, err
	}
	req = req.WithContext(ctx)

	client := http.Client{
		Timeout: c.timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return ListAccountsResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ListAccountsResponse{}, err
	}
	var listAccts ListAccountsResponse
	if err := json.Unmarshal(body, &listAccts); err != nil {
		return ListAccountsResponse{}, err
	}
	return listAccts, nil

}

func NewDefaultClient(url string) Client {
	return client{
		timeout: 10 * time.Second,
		url:     url,
	}
}
