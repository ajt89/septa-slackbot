include local.env

setup:
	- go get -d -v

compile:
	- go install github.com/ajt89/septa-slackbot

run:
	- export SLACK_TOKEN=$(SLACK_TOKEN); \
	septa-slackbot

build-docker:
	- docker build -t ajt89/septa-slackbot .

run-docker:
	- docker run --rm -e SLACK_TOKEN=$(SLACK_TOKEN) ajt89/septa-slackbot:latest
