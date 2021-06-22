package zulip

import (
	"encoding/base64"
	"net/http"
	"net/url"
)

type HttpClient interface {
	Do(r *http.Request) (*http.Response, error)
}

type Client interface {
	SendStreamMessage(content string, stream string, topic string) error
}

type client struct {
	baseUrl string
	token string
	email string

	httpClient HttpClient
}

func NewClient(baseUrl string, email string, token string, httpClient HttpClient) Client {
	return &client{
		baseUrl: baseUrl,
		email: email,
		token: token,
		httpClient: httpClient,
	}
}

func (c *client) auth() string {
	a := c.email + ":" + c.token
	return base64.StdEncoding.EncodeToString([]byte(a))
}

func (c *client) SendStreamMessage(content string, stream string, topic string) (err error) {
	req, err := http.NewRequest(http.MethodPost, c.baseUrl + "/api/v1/messages", nil)
	if err != nil {
		return
	}

	q := url.Values{}
	q.Add("type", "stream")
	q.Add("to", stream)
	q.Add("subject", topic)
	q.Add("content", content)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Basic " + c.auth())

	_, err = c.httpClient.Do(req)
	if err != nil {
		return
	}

	return
}