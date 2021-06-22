package zulip_test

import (
	"encoding/base64"
	"fmt"
	mocks "github.com/two-first-names/zulip-notify/mocks/pkg/zulip"
	. "github.com/two-first-names/zulip-notify/pkg/zulip"
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

	call := httpClient.On("Do", req).Return(nil, nil)

	_ = client.SendStreamMessage("This is a test message", "test-stream", "test-topic")

	fmt.Printf("%+v\n", call)

	httpClient.AssertExpectations(t)
}
