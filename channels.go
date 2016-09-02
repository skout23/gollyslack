package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("xoxb-74956207364-J9fqHu2xlOEJ2q6LgPWH5t7b")
	channels, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, channel := range channels {
		fmt.Println(channel.Name + " : " + channel.ID)
	}
}
