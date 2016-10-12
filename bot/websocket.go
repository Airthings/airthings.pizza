package main

import (
	"fmt"
	"log"
	"os"
  "strings"

  "github.com/nlopes/slack"
  "github.com/blankoslo/blank.pizza/common"
)

func main() {
	var api = slack.New(os.Getenv("SLACK_TOKEN"))
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				fmt.Printf("Message received: %v\n", ev)

        invitedSlackIDs := common.GetInvitedUsers()
        userIsInvited := common.StringInSlice(ev.User, invitedSlackIDs)

        if(userIsInvited) {
          if(strings.ToLower(ev.Text) == "ja") {
            common.Rsvp(ev.User, "attending")
            rtm.SendMessage(rtm.NewOutgoingMessage("Sweet! 👍", ev.Channel))
            common.FinalizeInvitationIfComplete()
          } else if (strings.ToLower(ev.Text) == "nei") {
            common.Rsvp(ev.User, "not attending")
            rtm.SendMessage(rtm.NewOutgoingMessage("Ok 😏", ev.Channel))
            common.InviteIfNeeded()
          } else {
            rtm.SendMessage(rtm.NewOutgoingMessage("Hehe jeg er litt dum, jeg. Skjønner jeg ikke helt hva du mener 😳. Kan du svare ja eller nei?", ev.Channel))
          }
        }

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:

				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}
}
