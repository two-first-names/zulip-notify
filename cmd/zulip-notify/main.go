package main

import (
	"crypto/tls"
	"github.com/two-first-names/zulip-notify/pkg/notify"
	"github.com/two-first-names/zulip-notify/pkg/zulip"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	zulipClient := zulip.NewClient(
		os.Getenv("ZULIP_ENDPOINT"),
		os.Getenv("ZULIP_EMAIL"),
		os.Getenv("ZULIP_TOKEN"),
		httpClient)

	notifier := notify.NewNotifier(
		os.Getenv("ZULIP_STREAM"),
		os.Getenv("ZULIP_TOPIC"),
		zulipClient)

	filename := os.Getenv("CONTENT_FILE")
	f, err := os.Open(filename)
	if err != nil {
		log.Panicln(err)
	}
	content, err := io.ReadAll(f)
	if err != nil {
		log.Panicln(err)
	}

	err = notifier.SendMessage(string(content))
	if err != nil {
		log.Panicln(err)
	}
}
