package main

import (
	"flag"
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

func GetTypeChan(api *slack.Client, id string) string {
	var err error

	_, err = api.GetChannelInfo(id)
	if err == nil {
		return "channel"
	}

	_, err = api.GetGroupInfo(id)
	if err == nil {
		return "group"
	}

	_, err = api.GetConversationInfo(id, false)
	if err == nil {
		return "conversation"
	}
	return ""

}

func main() {
	chanID := flag.String("c", "", "channel/groups ID where you want to invite everyone")

	slackToken := os.Getenv("SLACK_TOKEN")
	if len(slackToken) == 0 {
		fmt.Println("SLACK_TOKEN environment variable not set")
		return
	}

	flag.Parse()
	if len(*chanID) == 0 {
		flag.PrintDefaults()
		return
	}

	api := slack.New(slackToken)

	users, err := api.GetUsers()
	if err != nil {
		return
	}

	switch GetTypeChan(api, *chanID) {
	case "channel":
		for _, user := range users {
			if !user.Deleted {
				api.InviteUserToChannel(*chanID, user.ID)
				fmt.Printf("Inviting %s into the channel\n", user.Name)
			}
		}
	case "group":
		for _, user := range users {
			if !user.Deleted {
				api.InviteUserToGroup(*chanID, user.ID)
				fmt.Printf("Inviting %s into the group\n", user.Name)
			}
		}
	case "conversation":
		for _, user := range users {
			if !user.Deleted {
				api.InviteUsersToConversation(*chanID, user.ID)
				fmt.Printf("Inviting %s into the conversation\n", user.Name)
			}
		}
	}
}
