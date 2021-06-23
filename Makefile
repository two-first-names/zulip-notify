all: test build
build:
	go build github.com/two-first-names/zulip-notify/cmd/zulip-notify

test:
	go test github.com/two-first-names/zulip-notify/...

install:
	go install github.com/two-first-names/zulip-notify/cmd/zulip-notify

run:
	go run github.com/two-first-names/zulip-notify/cmd/zulip-notify