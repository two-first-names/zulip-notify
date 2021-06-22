package zulip

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
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

type sendMessageResponse struct {
	Result  string `json:"result"`
	Message string `json:"msg"`
	Id      *int   `json:"id"`
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

	res, err := c.httpClient.Do(req)
	if err != nil {
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = res.Body.Close()
	if err != nil {
		return
	}

	response := sendMessageResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	if response.Result == "error" {
		return errors.New("zulip error: " + response.Message)
	}

	return
}