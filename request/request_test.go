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
				body: `{"ref":"refs/heads/master","before":"69b3f38550378518a4d79983f9a8d041aa6c458e","after":"5979ddf50f80eece2af7ccaca21fcb776cbade3b","repository":{"id":215826756,"node_id":"MDEwOlJlcG9zaXRvcnkyMTU4MjY3NTY=","name":"hooktail","full_name":"klipitkas/hooktail","private":false,"owner":{"name":"klipitkas","email":"klipitkas@users.noreply.github.com","login":"klipitkas","id":3259834,"node_id":"MDQ6VXNlcjMyNTk4MzQ=","avatar_url":"https://avatars2.githubusercontent.com/u/3259834?v=4","gravatar_id":"","url":"https://api.github.com/users/klipitkas","html_url":"https://github.com/klipitkas","followers_url":"https://api.github.com/users/klipitkas/followers","following_url":"https://api.github.com/users/klipitkas/following{/other_user}","gists_url":"https://api.github.com/users/klipitkas/gists{/gist_id}","starred_url":"https://api.github.com/users/klipitkas/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/klipitkas/subscriptions","organizations_url":"https://api.github.com/users/klipitkas/orgs","repos_url":"https://api.github.com/users/klipitkas/repos","events_url":"https://api.github.com/users/klipitkas/events{/privacy}","received_events_url":"https://api.github.com/users/klipitkas/received_events","type":"User","site_admin":false},"html_url":"https://github.com/klipitkas/hooktail","description":"A golang server that manages github webhook deployments.","fork":false,"url":"https://github.com/klipitkas/hooktail","forks_url":"https://api.github.com/repos/klipitkas/hooktail/forks","keys_url":"https://api.github.com/repos/klipitkas/hooktail/keys{/key_id}","collaborators_url":"https://api.github.com/repos/klipitkas/hooktail/collaborators{/collaborator}","teams_url":"https://api.github.com/repos/klipitkas/hooktail/teams","hooks_url":"https://api.github.com/repos/klipitkas/hooktail/hooks","issue_events_url":"https://api.github.com/repos/klipitkas/hooktail/issues/events{/number}","events_url":"https://api.github.com/repos/klipitkas/hooktail/events","assignees_url":"https://api.github.com/repos/klipitkas/hooktail/assignees{/user}","branches_url":"https://api.github.com/repos/klipitkas/hooktail/branches{/branch}","tags_url":"https://api.github.com/repos/klipitkas/hooktail/tags","blobs_url":"https://api.github.com/repos/klipitkas/hooktail/git/blobs{/sha}","git_tags_url":"https://api.github.com/repos/klipitkas/hooktail/git/tags{/sha}","git_refs_url":"https://api.github.com/repos/klipitkas/hooktail/git/refs{/sha}","trees_url":"https://api.github.com/repos/klipitkas/hooktail/git/trees{/sha}","statuses_url":"https://api.github.com/repos/klipitkas/hooktail/statuses/{sha}","languages_url":"https://api.github.com/repos/klipitkas/hooktail/languages","stargazers_url":"https://api.github.com/repos/klipitkas/hooktail/stargazers","contributors_url":"https://api.github.com/repos/klipitkas/hooktail/contributors","subscribers_url":"https://api.github.com/repos/klipitkas/hooktail/subscribers","subscription_url":"https://api.github.com/repos/klipitkas/hooktail/subscription","commits_url":"https://api.github.com/repos/klipitkas/hooktail/commits{/sha}","git_commits_url":"https://api.github.com/repos/klipitkas/hooktail/git/commits{/sha}","comments_url":"https://api.github.com/repos/klipitkas/hooktail/comments{/number}","issue_comment_url":"https://api.github.com/repos/klipitkas/hooktail/issues/comments{/number}","contents_url":"https://api.github.com/repos/klipitkas/hooktail/contents/{+path}","compare_url":"https://api.github.com/repos/klipitkas/hooktail/compare/{base}...{head}","merges_url":"https://api.github.com/repos/klipitkas/hooktail/merges","archive_url":"https://api.github.com/repos/klipitkas/hooktail/{archive_format}{/ref}","downloads_url":"https://api.github.com/repos/klipitkas/hooktail/downloads","issues_url":"https://api.github.com/repos/klipitkas/hooktail/issues{/number}","pulls_url":"https://api.github.com/repos/klipitkas/hooktail/pulls{/number}","milestones_url":"https://api.github.com/repos/klipitkas/hooktail/milestones{/number}","notifications_url":"https://api.github.com/repos/klipitkas/hooktail/notifications{?since,all,participating}","labels_url":"https://api.github.com/repos/klipitkas/hooktail/labels{/name}","releases_url":"https://api.github.com/repos/klipitkas/hooktail/releases{/id}","deployments_url":"https://api.github.com/repos/klipitkas/hooktail/deployments","created_at":1571326320,"updated_at":"2019-10-19T15:24:55Z","pushed_at":1571498942,"git_url":"git://github.com/klipitkas/hooktail.git","ssh_url":"git@github.com:klipitkas/hooktail.git","clone_url":"https://github.com/klipitkas/hooktail.git","svn_url":"https://github.com/klipitkas/hooktail","homepage":null,"size":22,"stargazers_count":0,"watchers_count":0,"language":"Go","has_issues":true,"has_projects":true,"has_downloads":true,"has_wiki":false,"has_pages":false,"forks_count":0,"mirror_url":null,"archived":false,"disabled":false,"open_issues_count":0,"license":{"key":"unlicense","name":"The Unlicense","spdx_id":"Unlicense","url":"https://api.github.com/licenses/unlicense","node_id":"MDc6TGljZW5zZTE1"},"forks":0,"open_issues":0,"watchers":0,"default_branch":"master","stargazers":0,"master_branch":"master"},"pusher":{"name":"klipitkas","email":"klipitkas@users.noreply.github.com"},"sender":{"login":"klipitkas","id":3259834,"node_id":"MDQ6VXNlcjMyNTk4MzQ=","avatar_url":"https://avatars2.githubusercontent.com/u/3259834?v=4","gravatar_id":"","url":"https://api.github.com/users/klipitkas","html_url":"https://github.com/klipitkas","followers_url":"https://api.github.com/users/klipitkas/followers","following_url":"https://api.github.com/users/klipitkas/following{/other_user}","gists_url":"https://api.github.com/users/klipitkas/gists{/gist_id}","starred_url":"https://api.github.com/users/klipitkas/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/klipitkas/subscriptions","organizations_url":"https://api.github.com/users/klipitkas/orgs","repos_url":"https://api.github.com/users/klipitkas/repos","events_url":"https://api.github.com/users/klipitkas/events{/privacy}","received_events_url":"https://api.github.com/users/klipitkas/received_events","type":"User","site_admin":false},"created":false,"deleted":false,"forced":false,"base_ref":null,"compare":"https://github.com/klipitkas/hooktail/compare/69b3f3855037...5979ddf50f80","commits":[{"id":"5979ddf50f80eece2af7ccaca21fcb776cbade3b","tree_id":"06e6db392175983f16f55a513f531b8d35380ae5","distinct":true,"message":"Remove invalid unit test","timestamp":"2019-10-19T18:28:51+03:00","url":"https://github.com/klipitkas/hooktail/commit/5979ddf50f80eece2af7ccaca21fcb776cbade3b","author":{"name":"Konstantinos Lypitkas","email":"klipitkas@gmail.com","username":"klipitkas"},"committer":{"name":"Konstantinos Lypitkas","email":"klipitkas@gmail.com","username":"klipitkas"},"added":[],"removed":["main_test.go"],"modified":[]}],"head_commit":{"id":"5979ddf50f80eece2af7ccaca21fcb776cbade3b","tree_id":"06e6db392175983f16f55a513f531b8d35380ae5","distinct":true,"message":"Remove invalid unit test","timestamp":"2019-10-19T18:28:51+03:00","url":"https://github.com/klipitkas/hooktail/commit/5979ddf50f80eece2af7ccaca21fcb776cbade3b","author":{"name":"Konstantinos Lypitkas","email":"klipitkas@gmail.com","username":"klipitkas"},"committer":{"name":"Konstantinos Lypitkas","email":"klipitkas@gmail.com","username":"klipitkas"},"added":[],"removed":["main_test.go"],"modified":[]}}`,
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
			err := req.Parse([]byte(tt.args.body))
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
			req.JSONBody = tt.args.body
			req.Headers = make(map[string][]string, 1)
			req.Headers["X-Hub-Signature"] = []string{"sha1=77ca6ab111eac1d56346565bf3cdf6cdb0d2a890"}
			got := req.HasValidSignature(tt.args.secret)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %+v (%T), want = %+v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}

func TestRequestHash(t *testing.T) {

	type args struct {
		headerName  string
		headerValue string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test that the signature is correct",
			args{
				headerName:  "X-Hub-Signature",
				headerValue: "sha1=77ca6ab111eac1d56346565bf3cdf6cdb0d2a890",
			},
			"77ca6ab111eac1d56346565bf3cdf6cdb0d2a890",
		},
		{
			"Test that the signature is empty when header is not present",
			args{
				headerName:  "",
				headerValue: "",
			},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := request.Request{}
			req.Headers = make(map[string][]string, 1)
			req.Headers[tt.args.headerName] = []string{tt.args.headerValue}
			got := req.Hash()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %+v (%T), want = %+v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}
