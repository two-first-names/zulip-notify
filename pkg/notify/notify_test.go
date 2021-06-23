package notify_test

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)
import . "github.com/two-first-names/zulip-notify/pkg/notify"
import zulip "github.com/two-first-names/zulip-notify/mocks/pkg/zulip"

func TestNotifier_SendMessage(t *testing.T) {
	zulipClient := &zulip.Client{}
	notifier := NewNotifier("test-stream", "test-topic", zulipClient)

	zulipClient.
		On("SendStreamMessage", "This is a test message", "test-stream", "test-topic").
		Return(nil)

	err := notifier.SendMessage("This is a test message")

	require.Nil(t, err, "error should be nil")

	zulipClient.AssertExpectations(t)
}

func TestNotifier_SendMessage_Error(t *testing.T) {
	zulipClient := &zulip.Client{}
	notifier := NewNotifier("test-stream", "test-topic", zulipClient)

	e := errors.New("this is a test error")

	zulipClient.
		On("SendStreamMessage", "This is a test message", "test-stream", "test-topic").
		Return(e)

	err := notifier.SendMessage("This is a test message")

	require.Equal(t, e, err)

	zulipClient.AssertExpectations(t)
}