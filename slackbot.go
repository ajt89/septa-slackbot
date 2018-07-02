package main

import (
	"os"

	"golang.org/x/net/context"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
)

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Subrouter()
	toMe.Hear("(?i)(hi|hello).*").MessageHandler(HelloHandler)
	bot.Hear("(?)attachment").MessageHandler(AttachmentsHandler)
	toMe.Hear(".*(train view).*").MessageHandler(TrainViewHandler)
	bot.Run()
}

func HelloHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "Oh hello!", slackbot.WithTyping)
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

func AttachmentsHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	txt := "Beep Beep Boop is a ridiculously simple hosting platform for your Slackbots."
	attachment := slack.Attachment{
		Pretext:   "We bring bots to life. :sunglasses: :thumbsup:",
		Title:     "Host, deploy and share your bot in seconds.",
		TitleLink: "https://beepboophq.com/",
		Text:      txt,
		Fallback:  txt,
		ImageURL:  "https://storage.googleapis.com/beepboophq/_assets/bot-1.22f6fb.png",
		Color:     "#7CD197",
	}

	// supports multiple attachments
	attachments := []slack.Attachment{attachment}

	bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
}
