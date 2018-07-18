FROM golang:1.10.3
RUN mkdir -p /go/src/github.com/ajt89/septa-slackbot
WORKDIR /go/src/github.com/ajt89/septa-slackbot
ADD . /go/src/github.com/ajt89/septa-slackbot
ENV GOPATH /go
RUN make setup
RUN make compile
RUN make run
