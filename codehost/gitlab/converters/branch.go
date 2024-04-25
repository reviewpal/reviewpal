package converters

import (
	"github.com/reviewpal/reviewpal/codehost/target"
	"github.com/xanzy/go-gitlab"
)

func BranchFromGitlabBranch(b *gitlab.Branch) target.Branch {
	return target.Branch{
		Name: b.Name,
		SHA:  b.Commit.ID,
	}
}
