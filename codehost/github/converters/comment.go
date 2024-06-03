package converters

import (
	"github.com/google/go-github/v52/github"
	"github.com/reviewpal/reviewpal/codehost/target"
)

func CommentFromGithubComment(c *github.IssueComment) *target.Comment {
	return &target.Comment{
		ID:   c.GetID(),
		Body: c.GetBody(),
	}
}
