package converters

import (
	"github.com/google/go-github/v52/github"
	"github.com/reviewpal/reviewpal/codehost/target"
)

func CommitFromGithubCommit(commit *github.Commit) target.Commit {
	return target.Commit{
		SHA: commit.GetSHA(),
	}
}
