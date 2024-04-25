package actions

import (
	"context"

	"github.com/reviewpal/reviewpal/codehost/target"
)

func New(ctx context.Context, scmClient target.Target, targetEntity *target.Entity, pr *target.PullRequest) map[string]any {
	acts := &builtinActions{
		ctx,
		scmClient,
		targetEntity,
		pr,
	}

	return map[string]any{
		"$addLabels": acts.AddLabels,
	}
}

type builtinActions struct {
	ctx          context.Context
	scmClient    target.Target
	targetEntity *target.Entity
	pr           *target.PullRequest
}
