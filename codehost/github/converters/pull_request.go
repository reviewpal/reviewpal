package converters

import (
	"github.com/google/go-github/v52/github"
	"github.com/reviewpal/reviewpal/codehost/target"
)

func PullRequestFromGithubPullRequest(pr *github.PullRequest) *target.PullRequest {
	return &target.PullRequest{
		ID:      pr.GetID(),
		Number:  pr.GetNumber(),
		IsDraft: pr.GetDraft(),
		Status:  target.PullRequestStatus(pr.GetState()),
		Base:    BranchFromGithubBranch(pr.GetBase()),
		Head:    BranchFromGithubBranch(pr.GetHead()),
		User:    UserFromGithubUser(pr.GetUser()),
	}
}
