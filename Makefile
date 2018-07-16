include local.env

setup:
	- go get github.com/BeepBoopHQ/go-slackbot
	- go get github.com/nlopes/slack

compile:
	- go install github.com/ajt89/septa-slackbot

run:
	- export SLACK_TOKEN=$(SLACK_TOKEN); \
	septa-slackbot
