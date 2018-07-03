package main

import (
	"fmt"
	"os"
	"regexp"

	"golang.org/x/net/context"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
)

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Subrouter()
	toMe.Hear(".*(train view).*").MessageHandler(TrainViewHandler)
	toMe.Hear("train status .*").MessageHandler(TrainNumberHandler)
	bot.Run()
}

func TrainViewHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	data := GetTrainView()
	attachment := slack.Attachment{
		Pretext:   "Here is the train data",
		Title:     "Train View",
		TitleLink: "http://www3.septa.org/hackathon/TrainView/",
		Text:      data,
		Fallback:  data,
		Color:     "#7CD197",
	}

	// supports multiple attachments
	attachments := []slack.Attachment{attachment}

	bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
}

func TrainNumberHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	msg := slackbot.MessageFromContext(ctx)
	text := slackbot.StripDirectMention(msg.Text)
	re := regexp.MustCompile("train status (?P<TrainNo>.+)")
	trainStrArr := re.FindAllStringSubmatch(text, -1)[0]
	trainNo := trainStrArr[1]
	fmt.Printf(trainNo)
	dataJson := GetTrainNo()
	dataStr := string(dataJson)
	fmt.Printf(dataStr)

	attachment := slack.Attachment{
		Pretext:   "Here is the train json",
		Title:     "Train View",
		TitleLink: "http://www3.septa.org/hackathon/TrainView/",
		Text:      dataStr,
		Fallback:  dataStr,
		Color:     "#7CD197",
	}

	attachments := []slack.Attachment{attachment}

	bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
}
