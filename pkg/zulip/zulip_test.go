package zulip_test

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	mocks "github.com/two-first-names/zulip-notify/mocks/pkg/zulip"
	. "github.com/two-first-names/zulip-notify/pkg/zulip"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_SendStreamMessage(t *testing.T) {
	httpClient := &mocks.HttpClient{}
	client := NewClient("http://localhost:8080", "test@foo.com", "password", httpClient)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/messages", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := url.Values{}
	q.Add("type", "stream")
	q.Add("to", "test-stream")
	q.Add("subject", "test-topic")
	q.Add("content", "This is a test message")
	req.URL.RawQuery = q.Encode()

	auth := "test@foo.com:password"
	req.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(auth)))

	json := `{"id": 1, "msg": "", "result": "success"}`
	res := &http.Response{
		StatusCode: 200,
		Body: io.NopCloser(bytes.NewReader([]byte(json))),
	}

	httpClient.On("Do", req).Return(res, nil)

	err = client.SendStreamMessage("This is a test message", "test-stream", "test-topic")

	require.Nil(t, err, "SendTopicMessage should not return an error")

	httpClient.AssertExpectations(t)
}

func TestClient_SendStreamMessage_HttpClientError(t *testing.T) {
	httpClient := &mocks.HttpClient{}
	client := NewClient("http://localhost:8080", "test@foo.com", "password", httpClient)

	e := errors.New("this is a http client error")

	httpClient.On("Do", mock.Anything).Return(nil, e)

	err := client.SendStreamMessage("This is a test message", "test-stream", "test-topic")

	require.Equal(t, e, err)

	httpClient.AssertExpectations(t)
}

func TestClient_SendStreamMessage_ServiceError(t *testing.T) {
	httpClient := &mocks.HttpClient{}
	client := NewClient("http://localhost:8080", "test@foo.com", "password", httpClient)

	e := errors.New("zulip error: Stream 'nonexistent_stream' does not exist")

	json := `{"code": "STREAM_DOES_NOT_EXIST", "msg": "Stream 'nonexistent_stream' does not exist", "result": "error", "stream": "nonexistent_stream"}`
	res := &http.Response{
		StatusCode: 200,
		Body: io.NopCloser(bytes.NewReader([]byte(json))),
	}

	httpClient.On("Do", mock.Anything).Return(res, nil)

	err := client.SendStreamMessage("This is a test message", "nonexistent_stream", "test-topic")

	require.Equal(t, e, err)

	httpClient.AssertExpectations(t)
}