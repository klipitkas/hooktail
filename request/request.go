package request

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/klipitkas/hooktail/common"
)

// Request contains any needed request information.
type Request struct {
	Headers  map[string][]string
	JSONBody string
	Body     struct {
		Ref        string `json:"ref"`
		Before     string `json:"before"`
		After      string `json:"after"`
		Repository struct {
			ID       int    `json:"id"`
			NodeID   string `json:"node_id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			Private  bool   `json:"private"`
			Owner    struct {
				Name              string `json:"name"`
				Email             string `json:"email"`
				Login             string `json:"login"`
				ID                int    `json:"id"`
				NodeID            string `json:"node_id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"owner"`
			HTMLURL     string `json:"html_url"`
			Description string `json:"description"`
			Fork        bool   `json:"fork"`
			URL         string `json:"url"`
			GitURL      string `json:"git_url"`
			SSHURL      string `json:"ssh_url"`
			CloneURL    string `json:"clone_url"`
			SvnURL      string `json:"svn_url"`
			License     struct {
				Key    string `json:"key"`
				Name   string `json:"name"`
				SpdxID string `json:"spdx_id"`
				URL    string `json:"url"`
				NodeID string `json:"node_id"`
			} `json:"license"`
			Forks         int    `json:"forks"`
			OpenIssues    int    `json:"open_issues"`
			Watchers      int    `json:"watchers"`
			DefaultBranch string `json:"default_branch"`
			Stargazers    int    `json:"stargazers"`
			MasterBranch  string `json:"master_branch"`
		} `json:"repository"`
		Pusher struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"pusher"`
	}
}

// Parse the request from the body to the struct.
func (r *Request) Parse(body []byte) error {
	if err := json.Unmarshal(body, &r.Body); err != nil {
		return fmt.Errorf("unmarshal request body to json: %v", err)
	}
	return nil
}

// Hash returns the sha1 hash from the headers of the request.
func (r *Request) Hash() string {
	if r.Headers == nil ||
		r.Headers["X-Hub-Signature"] == nil ||
		r.Headers["X-Hub-Signature"][0] == "" {
		return ""
	}
	return strings.ReplaceAll(r.Headers["X-Hub-Signature"][0], "sha1=", "")
}

// HasValidSignature checks if the an HMAC hash has a valid
// signature given a key "secret".
func (r *Request) HasValidSignature(secret string) bool {
	return common.Sha1Hmac(r.JSONBody, secret) == r.Hash()
}
