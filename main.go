package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/thescriv/github-activity/pkg/github"
)

var formattedOutput = map[github.EventType]string{
	github.EventWatch:  "Watch %s",
	github.EventPush:   "Pushed %d commit(s) to %s",
	github.EventCreate: "Created a new repo %s",
	github.EventFork:   "Forked %s",
	github.EventDelete: "Deleted %s",

	github.EventPublic:                   "PublicEvent %s",
	github.EventGollum:                   "GollumEvent %s",
	github.EventIssueComment:             "IssueCommentEvent %s",
	github.EventIssues:                   "IssuesEvent %s",
	github.EventMember:                   "MemberEvent %s",
	github.EventPullRequest:              "PullRequestEvent %s",
	github.EventPullRequestReview:        "PullRequestReviewEvent %s",
	github.EventPullRequestReviewComment: "PullRequestReviewCommentEvent %s",
	github.EventPullRequestReviewThread:  "PullRequestReviewThreadEvent %s",
	github.EventRelease:                  "PullReleaseEvent %s",
	github.EventSponsorship:              "SponsorshipEvent %s",
}

func DisplayEvents(events []github.Events) {
	if len(events) == 0 {
		fmt.Println("No recent activity found")
		return
	}

	for _, event := range events {
		if event.Type == github.EventPush {
			fmt.Printf("- %s\n", fmt.Sprintf(formattedOutput[github.EventPush], event.Action.Size, event.Repo.Name))
			continue
		}

		fmt.Printf("- %s\n", fmt.Sprintf(formattedOutput[event.Type], event.Repo.Name))
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("you should give a valid username")

		return
	}

	ghUser := os.Args[1]

	client := http.Client{
		Timeout: time.Second * 5,
	}

	githubClient := github.NewClient(client)

	user, err := githubClient.GetUser(ghUser)
	if err != nil {
		fmt.Printf("getUser: %s\n", err)
		return
	}

	events, err := githubClient.GetEventsFromUser(user)
	if err != nil {
		fmt.Printf("getEventsFromUser: %s\n", err)

		return
	}

	DisplayEvents(events)
}
