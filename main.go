package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func botpost(pt string, t string, c string, m string) (string, error) {
	var title string
	var link string

	title = "Blah Title"
	link = "https://console.aws.amazon.com/inspector/home?region=us-east-1#/run"

	api := slack.New("blah") // required
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Color:     "good",
		Title:     title,
		TitleLink: link,
		//	Pretext:   pt,
		//	Text:      t,
	}
	params.AsUser = true
	params.Username = "Penny Gadget"
	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := api.PostMessage(c, m, params)
	if err != nil {
		return "failed to send message", err
	}
	return fmt.Sprintf("Message successfully sent to channel %s at %s", channelID, timestamp), nil
}

func main() {
	// test-out : C0PJW9L14
	// site-reliability-bots : C1X6UUFCY
	channel := "C0PJW9L14"
	_, err := botpost("Pretext to Attachment", "Summary Attached", channel, "")
	if err != nil {
		fmt.Println(err)
	}

}
