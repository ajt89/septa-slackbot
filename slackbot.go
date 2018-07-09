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
	toMe.Hear("(train status) .*").MessageHandler(TrainNumberHandler)
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
	dataJson := GetTrainNo(trainNo)
	var returnText string
	if dataJson.status == 1 {
		returnText = dataJson.errorMsg
	} else {
		nextStop := dataJson.data[0]
		late := dataJson.data[1]
		dest := dataJson.data[2]
		returnText = fmt.Sprintf("The next stop for train %s(%s) is %s and is %s minute(s) late", trainNo, dest, nextStop, late)
	}
	fmt.Println(returnText)

	attachment := slack.Attachment{
		Pretext:   "Here is the train json",
		Title:     "Train View",
		TitleLink: "http://www3.septa.org/hackathon/TrainView/",
		Text:      returnText,
		Fallback:  returnText,
		Color:     "#7CD197",
	}

	attachments := []slack.Attachment{attachment}

	bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
}
