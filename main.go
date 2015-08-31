// slackbot is a simple command line program that posts messages in a slack channel.
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

// Error Codes that are returned via os.Exit
const (
	ErrorNone = iota
	ErrorPanic
	ErrorTokenEmpty
	ErrorChannelEmpty
)

func errorExitParm(message string, returnCode int) {
	if !*fQuiet {
		fmt.Println(message)
	}
	os.Exit(returnCode)
}

func errorExitPanic(err error) {
	if !*fQuiet {
		panic(err)
	} else {
		os.Exit(ErrorPanic)
	}
}

func initFlags() {
	flag.Parse()

	if *fToken == "" {
		errorExitParm("-token is a required parameter.", ErrorTokenEmpty)
	} else if *fChannel == "" {
		errorExitParm("-channel is a required parameter.", ErrorChannelEmpty)
	}
}

func main() {
	initFlags()

	api := slack.New(*fToken)

	channel, err := api.FindChannelByName(*fChannel)
	if err != nil {
		errorExitPanic(err)
	}

	err = api.ChatPostMessage(channel.Id, *fMessage, nil)
	if err != nil {
		errorExitPanic(err)
	}
}
