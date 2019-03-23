include local.env
tag = active

setup:
	- go get -d -v

compile:
	- go install github.com/ajt89/septa-slackbot

run:
	- export SLACK_TOKEN=$(SLACK_TOKEN); \
	septa-slackbot

test:
	- go test github.com/ajt89/septa-slackbot/septa -v

coverage:
	- go test github.com/ajt89/septa-slackbot/septa -cover

build-docker:
	-docker build -t ajt89/septa-slackbot:$(tag) .

push-tag-docker:
	-docker puish ajt89/septa-slackbot:$(tag)

run-docker:
	- docker run --rm -e SLACK_TOKEN=$(SLACK_TOKEN) ajt89/septa-slackbot:$(tag)
