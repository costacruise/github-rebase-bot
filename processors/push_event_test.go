package processors

import (
	"testing"

	"github.com/google/go-github/github"
)

func TestPushEvent(t *testing.T) {
	ch := make(chan *github.PushEvent, 1)

	t.Run("adds open PRs on mainline push", func(t *testing.T) {
		out := PushEvent(Repository{
			Owner:    "test",
			Name:     "test",
			Mainline: "master",
		}, fakePullRequestResponse(2), ch)
		ch <- &github.PushEvent{
			Ref: stringVal("refs/heads/master"),
			Repo: &github.PushEventRepository{
				Name: stringVal("test"),
				Owner: &github.PushEventRepoOwner{
					Name: stringVal("test"),
				},
			},
		}

		close(ch)

		if _, ok := <-out; !ok {
			t.Fatal("Expected output on PR, but didn't receive")
		}
		if _, ok := <-out; !ok {
			t.Fatal("Expected output on PR, but didn't receive")
		}
	})
}
