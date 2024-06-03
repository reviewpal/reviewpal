package converters

import (
	"github.com/reviewpal/reviewpal/codehost/target"
	"github.com/xanzy/go-gitlab"
)

func ReviewFromGitlabApproval(r *gitlab.MergeRequestApprovals) *target.Review {
	return &target.Review{
		ID:    int64(r.ID),
		State: target.Approved,
	}
}
