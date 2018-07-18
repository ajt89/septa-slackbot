FROM golang:1.10.3
RUN mkdir -p /go/src/github.com/ajt89/septa-slackbot
WORKDIR /go/src/github.com/ajt89/septa-slackbot
ADD . /go/src/github.com/ajt89/septa-slackbot
ENV GOPATH /go
RUN go get github.com/BeepBoopHQ/go-slackbot
RUN go get github.com/nlopes/slack
RUN go install .
RUN make run
