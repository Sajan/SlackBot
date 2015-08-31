package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bluele/slack"
)

var fQuiet = flag.Bool("quiet", false, "Exit and set an error code, but do not print an error message on the console.")
var fToken = flag.String("token", "", "slack authentication token")
var fChannel = flag.String("channel", "", "slack channel name")
var fMessage = flag.String("message", "", "message")

const (
	ErrorNone = iota
	ErrorPanic
	ErrorTokenEmpty
	ErrorChannelEmpty
	ErrorMessageEmpty
)

func ErrorExitParm(message string, returnCode int) {
	if !*fQuiet {
		fmt.Println(message)
	}
	os.Exit(returnCode)
}

func ErrorExitPanic(err error) {
	if !*fQuiet {
		panic(err)
	} else {
		os.Exit(ErrorPanic)
	}
}

func initFlags() {
	flag.Parse()

	if *fToken == "" {
		ErrorExitParm("-token is a required parameter.", ErrorTokenEmpty)
	} else if *fChannel == "" {
		ErrorExitParm("-channel is a required parameter.", ErrorChannelEmpty)
	} else if *fMessage == "" {
		ErrorExitParm("-message is a required parameter.", ErrorMessageEmpty)
	}
}

func main() {
	initFlags()

	api := slack.New(*fToken)

	channel, err := api.FindChannelByName(*fChannel)
	if err != nil {
		ErrorExitPanic(err)
	}

	err = api.ChatPostMessage(channel.Id, *fMessage, nil)
	if err != nil {
		ErrorExitPanic(err)
	}
}
