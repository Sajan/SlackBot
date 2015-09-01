// slackbot is a simple command line program that posts messages in a slack channel.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/bluele/slack"
)

var fToken = flag.String("token", "", "slack authentication token")
var fChannel = flag.String("channel", "", "slack channel name")
var fMessage = flag.String("message", "", "message")

// Error Codes that are returned via os.Exit
const (
	ErrorNone = iota
	ErrorUnknown
	ErrorTokenEmpty
	ErrorChannelEmpty
)

func errorExitParm(message string, returnCode int) {
	fmt.Println(message)
	os.Exit(returnCode)
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
		panic(err)
	}

	if len(*fMessage) > 0 {
		err = api.ChatPostMessage(channel.Id, *fMessage, nil)
		if err != nil {
			panic(err)
		}
	}

	if !termutil.Isatty(os.Stdin.Fd()) {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		err = api.ChatPostMessage(channel.Id, string(b), nil)
		if err != nil {
			panic(err)
		}
	}
}
