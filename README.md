# Zulip Notify

A Go program to send a message to a Zulip channel.

The intention is to have this run on a cronjob and send a message every Monday morning.

## Usage

To get the application:
```shell
go install github.com/two-first-names/zulip-notify/cmd/zulip-notify
```

To run:
```shell
ZULIP_ENDPOINT=https://zulip.example.com \
ZULIP_STREAM=general \
ZULIP_TOPIC="Monday Morning Message" \
ZULIP_EMAIL=bot@zulip.example.com \
ZULIP_TOKEN=securetokenhere \
CONTENT_FILENAME=monday.md \
zulip-notify
```