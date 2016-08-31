package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("Insert Token Here") // required
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Pretext: "Results from Assessment Run",
		Text:    "Summary Text of AR",
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}
	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := api.PostMessage("CHANNELID", "wakka wakka", params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
