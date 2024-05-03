package target

import (
	"context"
)

type Target interface {
	GetPullRequest(ctx context.Context) (*PullRequest, error)
	GetPullRequestPatch(ctx context.Context) (Patch, error)
	AddLabels(ctx context.Context, labels []string) ([]string, error)
	Comment(ctx context.Context, comment string) error
}
