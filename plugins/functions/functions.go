package functions

import (
	"context"

	"github.com/reviewpal/reviewpal/codehost/target"
)

func New(ctx context.Context, scmClient target.Target, targetEntity *target.Entity, pr *target.PullRequest) map[string]any {
	funcs := &builtinFunctions{
		ctx,
		scmClient,
		targetEntity,
		pr,
	}

	return map[string]any{
		"$author": funcs.Author,
	}
}

type builtinFunctions struct {
	ctx          context.Context
	githubClient target.Target
	targetEntity *target.Entity
	pr           *target.PullRequest
}
