package converters

import (
	"github.com/reviewpal/reviewpal/codehost/target"
	"github.com/xanzy/go-gitlab"
)

func CommentFromGitlabNote(c *gitlab.Note) *target.Comment {
	return &target.Comment{
		ID:   int64(c.ID),
		Body: c.Body,
	}
}
