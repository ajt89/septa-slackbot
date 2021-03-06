package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"

	"github.com/ajt89/septa-slackbot/septa"
)

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Subrouter()
	toMe.Hear(".*(train view).*").MessageHandler(TrainViewHandler)
	toMe.Hear("(train status) .*").MessageHandler(TrainNumberHandler)
	toMe.Hear("(get trains next to arrive at) .*").MessageHandler(NextToArriveHandler)
	toMe.Hear(".*(get all trains).*").MessageHandler(GetAllTrainNumbers)
	bot.Run()
}

// TrainViewHandler returns the raw response from the TrainView endpoint
func TrainViewHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	data := septa.GetTrainView()
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

// TrainNumberHandler will retrieve data on a specific train number
func TrainNumberHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	msg := slackbot.MessageFromContext(ctx)
	text := slackbot.StripDirectMention(msg.Text)

	re := regexp.MustCompile("train status (?P<TrainNo>.+)")
	trainStrArr := re.FindAllStringSubmatch(text, -1)[0]
	trainNo := trainStrArr[1]
	bot.Reply(evt, fmt.Sprintf("Ok, looking for data on %s", trainNo), slackbot.WithTyping)

	getTrainNoResponse := septa.GetTrainNo(trainNo)
	var returnText string

	if getTrainNoResponse.Status == 1 {
		returnText = getTrainNoResponse.ErrorMsg
	} else {
		nextStop := getTrainNoResponse.Data.NextStop
		late := getTrainNoResponse.Data.Late
		dest := getTrainNoResponse.Data.Dest
		returnText = fmt.Sprintf("The next stop for train %s (%s) is %s and is %s minute(s) late", trainNo, dest, nextStop, late)
	}

	attachment := slack.Attachment{
		Title:     fmt.Sprintf("Train %s", trainNo),
		TitleLink: "http://www3.septa.org/hackathon/TrainView/",
		Text:      returnText,
		Fallback:  returnText,
		Color:     "#7CD197",
	}

	attachments := []slack.Attachment{attachment}

	bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
}

// NextToArriveHandler returns all trains next to arrive at a given station
func NextToArriveHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	msg := slackbot.MessageFromContext(ctx)
	text := slackbot.StripDirectMention(msg.Text)

	re := regexp.MustCompile("get trains next to arrive at (?P<StationName>.+)")
	stationStrArr := re.FindAllStringSubmatch(text, -1)[0]
	stationName := stationStrArr[1]
	bot.Reply(evt, fmt.Sprintf("Ok, looking for trains next to arrive at %s", stationName), slackbot.WithTyping)

	getTrainNoResponse := septa.GetAllTrainsNextToArrive(stationName)
	var returnText string
	var returnTitle string

	if getTrainNoResponse.Status == 1 {
		returnText = getTrainNoResponse.ErrorMsg
	} else if len(getTrainNoResponse.Data) > 0 {
		returnTitle = fmt.Sprintf("Trains Next to arrive at %s", stationName)
		trainNumberArray := getTrainNoResponse.Data
		returnText = fmt.Sprintf(strings.Join(trainNumberArray, "\n "))
	} else {
		returnTitle = fmt.Sprintf("No trains next to arrive at %s", stationName)
	}

	attachment := slack.Attachment{
		Title:     returnTitle,
		TitleLink: "http://www3.septa.org/hackathon/TrainView/",
		Text:      returnText,
		Fallback:  returnText,
		Color:     "#7CD197",
	}

	attachments := []slack.Attachment{attachment}

	bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
}

// GetAllTrainNumbers returns all train numbers
func GetAllTrainNumbers(ct context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "Ok, looking for all train numbers", slackbot.WithTyping)
	getAllTrainNosResponse := septa.GetAllTrainNos()
	var returnText string

	if getAllTrainNosResponse.Status == 1 {
		returnText = getAllTrainNosResponse.ErrorMsg
	} else {
		trainNumberArray := getAllTrainNosResponse.Data
		returnText = fmt.Sprintf(strings.Join(trainNumberArray, ", "))
	}

	attachment := slack.Attachment{
		Title:     "All current train numbers",
		TitleLink: "http://www3.septa.org/hackathon/TrainView/",
		Text:      returnText,
		Fallback:  returnText,
		Color:     "#7CD197",
	}

	attachments := []slack.Attachment{attachment}

	bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
}
