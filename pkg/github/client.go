package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const githubApiUserURL = "https://api.github.com/users/%s"
const githubApiEventsURL = "https://api.github.com/users/%s/events"

type Client struct {
	http http.Client
}

func NewClient(httpClient http.Client) Client {
	return Client{http: httpClient}
}

func (c Client) get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("User-Agent", "thescriv/github-activity")

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("user does not exist, check your input")
		}

		return nil, fmt.Errorf("error fetching data %d", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

func (c Client) GetUser(ghUser string) (User, error) {
	getUserUrl := fmt.Sprintf(githubApiUserURL, ghUser)

	body, err := c.get(getUserUrl)
	if err != nil {
		return User{}, fmt.Errorf("get: %w", err)
	}

	u := User{}
	err = json.Unmarshal(body, &u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (c Client) GetEventsFromUser(user User) ([]Events, error) {
	body, err := c.get(fmt.Sprintf(githubApiEventsURL, user.Login))
	if err != nil {
		return []Events{}, fmt.Errorf("get: %w", err)
	}

	e := []Events{}
	err = json.Unmarshal(body, &e)
	if err != nil {
		return []Events{}, err
	}

	return e, nil
}
