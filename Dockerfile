FROM golang:1.10.3
WORKDIR /go/src/github.com/ajt89/septa-slackbot
COPY . .
RUN go get -d -v
RUN go install github.com/ajt89/septa-slackbot
CMD septa-slackbot
