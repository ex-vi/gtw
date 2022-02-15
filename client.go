package gtw

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	twitterAPIBaseURL = "https://api.twitter.com"
)

type Client struct {
	restClient *resty.Client
}

type Parameters struct {
	Body        interface{}
	PathParams  map[string]string
	QueryParams map[string]string
}

func NewClient(apiToken string) *Client {
	c := resty.New()

	c.SetBaseURL(twitterAPIBaseURL)
	c.SetAuthToken(apiToken)

	return &Client{
		restClient: c,
	}
}

func (c *Client) Request(ctx context.Context, method, path string, params Parameters, resp interface{}) error {
	req := c.restClient.R().
		SetBody(params.Body).
		SetPathParams(params.PathParams).
		SetQueryParams(params.QueryParams)

	res, err := req.Execute(method, path)
	if err != nil {
		return errors.Wrap(err, "could not perform request to the twitter api")
	}

	switch res.StatusCode() {
	case http.StatusOK:
	case http.StatusCreated:
	default:
		return errors.Errorf("request to the twitter api failed. code: %d", res.StatusCode())
	}

	return json.Unmarshal(res.Body(), resp)
}
