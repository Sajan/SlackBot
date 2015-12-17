// slackbot is a simple command line program that posts messages in a slack channel.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/nlopes/slack"
)

var fChannel = flag.String("channel", "", "slack channel name (required)")
var fMessage = flag.String("message", "", "message")
var fToken = flag.String("token", "", "slack authentication token (required)")
var fUsername = flag.String("username", slack.DEFAULT_MESSAGE_USERNAME, "username to use when posting")

// Error Codes that are returned via os.Exit
const (
	ErrorNone = iota
	ErrorUnknown
	ErrorTokenEmpty
	ErrorChannelEmpty
)

func errorExitParm(message string, returnCode int) {
	flag.Usage()

	fmt.Println()
	fmt.Println(message)

	os.Exit(returnCode)
}

func initFlags() {
	flag.Usage = func() {
		fmt.Println("Full Source Available Here: https://github.com/Sajan/slackbot")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *fToken == "" {
		errorExitParm("-token is a required parameter, but it is missing.", ErrorTokenEmpty)
	} else if *fChannel == "" {
		errorExitParm("-channel is a required parameter, but it is missing.", ErrorChannelEmpty)
	}
}

func slackMessage(api *slack.Client, message string) {
	if len(message) > 0 {
		parms := slack.PostMessageParameters{
			Username: *fUsername,
		}
		_, _, err := api.PostMessage(*fChannel, message, parms)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	initFlags()

	api := slack.New(*fToken)

	// If the user filled in the -message parameter, send that message out
	slackMessage(api, *fMessage)

	// If the user passed in input via stdin, send that out too
	if !termutil.Isatty(os.Stdin.Fd()) {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		slackMessage(api, string(b))
	}
}
