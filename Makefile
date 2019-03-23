include local.env
TAG = active
DOCKER_HOST = ajt89

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
	-docker build -t $(DOCKER_HOST)/septa-slackbot:$(TAG) .

push-tag-docker:
	-docker puish $(DOCKER_HOST)/septa-slackbot:$(TAG)

run-docker:
	- docker run --rm -e SLACK_TOKEN=$(SLACK_TOKEN) $(DOCKER_HOST)/septa-slackbot:$(TAG)
