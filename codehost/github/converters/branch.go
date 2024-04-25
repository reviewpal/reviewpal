package converters

import (
	"github.com/google/go-github/v52/github"
	"github.com/reviewpal/reviewpal/codehost/target"
)

func BranchFromGithubBranch(b *github.PullRequestBranch) target.Branch {
	return target.Branch{
		Name: b.GetRef(),
		SHA:  b.GetSHA(),
	}
}
