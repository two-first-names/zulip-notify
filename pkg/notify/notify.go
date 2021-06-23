package notify

import "github.com/two-first-names/zulip-notify/pkg/zulip"

type Notifier interface {
	SendMessage(content string) error
}

type notifier struct {
	stream string
	topic string
	zulipClient zulip.Client
}

func NewNotifier(stream string, topic string, zulipClient zulip.Client) Notifier {
	return &notifier{
		stream: stream,
		topic: topic,
		zulipClient: zulipClient,
	}
}

func (n *notifier) SendMessage(content string) (err error) {
	err = n.zulipClient.SendStreamMessage(content, n.stream, n.topic)
	if err != nil {
		return
	}
	return
}
