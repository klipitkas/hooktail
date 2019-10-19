package request_test

import (
	"reflect"
	"testing"

	request "github.com/klipitkas/hooktail/request"
)

func TestRequestParse(t *testing.T) {

	type args struct {
		body string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Parse valid request body",
			args{
				body: `{"zen":"It's not fully shipped until it's fast.","hook_id":150331826,"hook":{"type":"Repository","id":150331826,"name":"web","active":true,"events":["push"],"config":{"content_type":"json","insecure_ssl":"1","secret":"********","url":"https://cb488464.ngrok.io"},"updated_at":"2019-10-19T12:20:55Z","created_at":"2019-10-19T12:20:55Z","url":"https://api.github.com/repos/klipitkas/hooktail/hooks/150331826","test_url":"https://api.github.com/repos/klipitkas/hooktail/hooks/150331826/test","ping_url":"https://api.github.com/repos/klipitkas/hooktail/hooks/150331826/pings","last_response":{"code":null,"status":"unused","message":null}},"repository":{"id":215826756,"node_id":"MDEwOlJlcG9zaXRvcnkyMTU4MjY3NTY=","name":"hooktail","full_name":"klipitkas/hooktail","private":true,"owner":{"login":"klipitkas","id":3259834,"node_id":"MDQ6VXNlcjMyNTk4MzQ=","avatar_url":"https://avatars2.githubusercontent.com/u/3259834?v=4","gravatar_id":"","url":"https://api.github.com/users/klipitkas","html_url":"https://github.com/klipitkas","followers_url":"https://api.github.com/users/klipitkas/followers","following_url":"https://api.github.com/users/klipitkas/following{/other_user}","gists_url":"https://api.github.com/users/klipitkas/gists{/gist_id}","starred_url":"https://api.github.com/users/klipitkas/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/klipitkas/subscriptions","organizations_url":"https://api.github.com/users/klipitkas/orgs","repos_url":"https://api.github.com/users/klipitkas/repos","events_url":"https://api.github.com/users/klipitkas/events{/privacy}","received_events_url":"https://api.github.com/users/klipitkas/received_events","type":"User","site_admin":false},"html_url":"https://github.com/klipitkas/hooktail","description":"A golang server that manages github webhook deployments.","fork":false,"url":"https://api.github.com/repos/klipitkas/hooktail","forks_url":"https://api.github.com/repos/klipitkas/hooktail/forks","keys_url":"https://api.github.com/repos/klipitkas/hooktail/keys{/key_id}","collaborators_url":"https://api.github.com/repos/klipitkas/hooktail/collaborators{/collaborator}","teams_url":"https://api.github.com/repos/klipitkas/hooktail/teams","hooks_url":"https://api.github.com/repos/klipitkas/hooktail/hooks","issue_events_url":"https://api.github.com/repos/klipitkas/hooktail/issues/events{/number}","events_url":"https://api.github.com/repos/klipitkas/hooktail/events","assignees_url":"https://api.github.com/repos/klipitkas/hooktail/assignees{/user}","branches_url":"https://api.github.com/repos/klipitkas/hooktail/branches{/branch}","tags_url":"https://api.github.com/repos/klipitkas/hooktail/tags","blobs_url":"https://api.github.com/repos/klipitkas/hooktail/git/blobs{/sha}","git_tags_url":"https://api.github.com/repos/klipitkas/hooktail/git/tags{/sha}","git_refs_url":"https://api.github.com/repos/klipitkas/hooktail/git/refs{/sha}","trees_url":"https://api.github.com/repos/klipitkas/hooktail/git/trees{/sha}","statuses_url":"https://api.github.com/repos/klipitkas/hooktail/statuses/{sha}","languages_url":"https://api.github.com/repos/klipitkas/hooktail/languages","stargazers_url":"https://api.github.com/repos/klipitkas/hooktail/stargazers","contributors_url":"https://api.github.com/repos/klipitkas/hooktail/contributors","subscribers_url":"https://api.github.com/repos/klipitkas/hooktail/subscribers","subscription_url":"https://api.github.com/repos/klipitkas/hooktail/subscription","commits_url":"https://api.github.com/repos/klipitkas/hooktail/commits{/sha}","git_commits_url":"https://api.github.com/repos/klipitkas/hooktail/git/commits{/sha}","comments_url":"https://api.github.com/repos/klipitkas/hooktail/comments{/number}","issue_comment_url":"https://api.github.com/repos/klipitkas/hooktail/issues/comments{/number}","contents_url":"https://api.github.com/repos/klipitkas/hooktail/contents/{+path}","compare_url":"https://api.github.com/repos/klipitkas/hooktail/compare/{base}...{head}","merges_url":"https://api.github.com/repos/klipitkas/hooktail/merges","archive_url":"https://api.github.com/repos/klipitkas/hooktail/{archive_format}{/ref}","downloads_url":"https://api.github.com/repos/klipitkas/hooktail/downloads","issues_url":"https://api.github.com/repos/klipitkas/hooktail/issues{/number}","pulls_url":"https://api.github.com/repos/klipitkas/hooktail/pulls{/number}","milestones_url":"https://api.github.com/repos/klipitkas/hooktail/milestones{/number}","notifications_url":"https://api.github.com/repos/klipitkas/hooktail/notifications{?since,all,participating}","labels_url":"https://api.github.com/repos/klipitkas/hooktail/labels{/name}","releases_url":"https://api.github.com/repos/klipitkas/hooktail/releases{/id}","deployments_url":"https://api.github.com/repos/klipitkas/hooktail/deployments","created_at":"2019-10-17T15:32:00Z","updated_at":"2019-10-19T09:48:20Z","pushed_at":"2019-10-19T09:39:03Z","git_url":"git://github.com/klipitkas/hooktail.git","ssh_url":"git@github.com:klipitkas/hooktail.git","clone_url":"https://github.com/klipitkas/hooktail.git","svn_url":"https://github.com/klipitkas/hooktail","homepage":null,"size":14,"stargazers_count":0,"watchers_count":0,"language":"Go","has_issues":true,"has_projects":true,"has_downloads":true,"has_wiki":true,"has_pages":false,"forks_count":0,"mirror_url":null,"archived":false,"disabled":false,"open_issues_count":0,"license":{"key":"unlicense","name":"The Unlicense","spdx_id":"Unlicense","url":"https://api.github.com/licenses/unlicense","node_id":"MDc6TGljZW5zZTE1"},"forks":0,"open_issues":0,"watchers":0,"default_branch":"master"},"sender":{"login":"klipitkas","id":3259834,"node_id":"MDQ6VXNlcjMyNTk4MzQ=","avatar_url":"https://avatars2.githubusercontent.com/u/3259834?v=4","gravatar_id":"","url":"https://api.github.com/users/klipitkas","html_url":"https://github.com/klipitkas","followers_url":"https://api.github.com/users/klipitkas/followers","following_url":"https://api.github.com/users/klipitkas/following{/other_user}","gists_url":"https://api.github.com/users/klipitkas/gists{/gist_id}","starred_url":"https://api.github.com/users/klipitkas/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/klipitkas/subscriptions","organizations_url":"https://api.github.com/users/klipitkas/orgs","repos_url":"https://api.github.com/users/klipitkas/repos","events_url":"https://api.github.com/users/klipitkas/events{/privacy}","received_events_url":"https://api.github.com/users/klipitkas/received_events","type":"User","site_admin":false}}`,
			},
			"git@github.com:klipitkas/hooktail.git",
			false,
		},
		{
			"Parse invalid request body should fail",
			args{
				body: `this is a non json request`,
			},
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := request.Request{}
			err := req.Parse(tt.args.body)
			got := req.Body.Repository.SSHURL
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %+v (%T), want = %+v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}

func TestRequestHasValidSignature(t *testing.T) {

	type args struct {
		body   string
		secret string
		hash   string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test that the check for valid signature works correctly",
			args{
				body:   `{"zen":"It's not fully shipped until it's fast."}`,
				secret: "love",
				hash:   "77ca6ab111eac1d56346565bf3cdf6cdb0d2a890",
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := request.Request{}
			got := req.HasValidSignature(tt.args.secret, tt.args.body, tt.args.hash)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %+v (%T), want = %+v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}
