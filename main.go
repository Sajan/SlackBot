package main

import (
	"flag"
	"fmt"

	"github.com/bluele/slack"
)

var fToken = flag.String("token", "", "slack authentication token")
var fChannel = flag.String("channel", "", "slack channel name")
var fMessage = flag.String("message", "", "message")

func initFlags() bool {
	flag.Parse()
	if *fToken == "" {
		fmt.Println("-token is a required parameter.")
		return false
	}
	if *fChannel == "" {
		fmt.Println("-channel is a required parameter.")
		return false
	}
	if *fMessage == "" {
		fmt.Println("-message is a required parameter.")
		return false
	}
	return true
}

func main() {
	if !initFlags() {
		return
	}

	api := slack.New(*fToken)
	channel, err := api.FindChannelByName(*fChannel)
	if err != nil {
		panic(err)
	}
	err = api.ChatPostMessage(channel.Id, *fMessage, nil)
	if err != nil {
		panic(err)
	}
}
