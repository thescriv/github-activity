package github

import "time"

type User struct {
	Login       string    `json:"login"`
	ID          int64     `json:"id"`
	NodeID      string    `json:"node_id"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	Company     string    `json:"company"`
	Email       string    `json:"email"`
	Hireable    string    `json:"hireable"`
	Bio         string    `json:"bio"`
	PublicRepos int64     `json:"public_repos"`
	PublicGists int64     `json:"public_gists"`
	Follower    int64     `json:"followers"`
	Following   int64     `json:"following"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type EventsActor struct {
	ID           int64  `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
}

type EventsRepo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EventsPayload struct {
	Action       string `json:"action"`
	RepositoryID int64  `json:"repository_id"`
	PushID       int64  `json:"push_id"`
	Size         int64  `json:"size"`
}

type EventType string

const (
	EventWatch                    EventType = "WatchEvent"
	EventCreate                   EventType = "CreateEvent"
	EventPublic                   EventType = "PublicEvent"
	EventPush                     EventType = "PushEvent"
	EventFork                     EventType = "ForkEvent"
	EventDelete                   EventType = "DeleteEvent"
	EventGollum                   EventType = "GollumEvent"
	EventIssueComment             EventType = "IssueCommentEvent"
	EventIssues                   EventType = "IssuesEvent"
	EventMember                   EventType = "MemberEvent"
	EventPullRequest              EventType = "PullRequestEvent"
	EventPullRequestReview        EventType = "PullRequestReviewEvent"
	EventPullRequestReviewComment EventType = "PullRequestReviewCommentEvent"
	EventPullRequestReviewThread  EventType = "PullRequestReviewThreadEvent"
	EventRelease                  EventType = "PullReleaseEvent"
	EventSponsorship              EventType = "SponsorshipEvent"
)

type Events struct {
	ID     string        `json:"id"`
	Type   EventType     `json:"type"`
	Actor  EventsActor   `json:"actor"`
	Repo   EventsRepo    `json:"repo"`
	Action EventsPayload `json:"payload"`
	Public bool          `json:"public"`
}
