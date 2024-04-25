package converters

import (
	"github.com/reviewpal/reviewpal/codehost/target"
	"github.com/xanzy/go-gitlab"
)

func PullRequestFromGitlabPullRequest(pr *gitlab.MergeRequest) *target.PullRequest {
	return &target.PullRequest{
		ID:     int64(pr.ID),
		Number: pr.IID,
		Base: BranchFromGitlabBranch(&gitlab.Branch{
			Name: pr.TargetBranch,
			Commit: &gitlab.Commit{
				ID: pr.DiffRefs.BaseSha,
			},
		}),
		Head: BranchFromGitlabBranch(&gitlab.Branch{
			Name: pr.SourceBranch,
			Commit: &gitlab.Commit{
				ID: pr.DiffRefs.HeadSha,
			},
		}),
		User: UserFromGitlabUser(pr.Author),
	}
}
