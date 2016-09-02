package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("bah")
	channels, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, channel := range channels {
		fmt.Println(channel.Name + " : " + channel.ID)
	}
}
