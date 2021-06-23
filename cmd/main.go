package main

import (
	"github.com/two-first-names/zulip-notify/pkg/notify"
	"github.com/two-first-names/zulip-notify/pkg/zulip"
	"log"
	"net/http"
	"os"
)

func main() {
	httpClient := &http.Client{}
	zulipClient := zulip.NewClient(
		os.Getenv("ZULIP_ENDPOINT"),
		os.Getenv("ZULIP_EMAIL"),
		os.Getenv("ZULIP_TOKEN"),
		httpClient)

	notifier := notify.NewNotifier(
		os.Getenv("ZULIP_STREAM"),
		os.Getenv("ZULIP_TOPIC"),
		zulipClient)

	err := notifier.SendMessage("This is a test message!")
	if err != nil {
		log.Fatalln(err)
	}
}
