package actions

import (
	"context"
	"log/slog"

	"github.com/reviewpal/reviewpal/codehost/target"
)

func New(ctx context.Context, scmClient target.Target, targetEntity *target.Entity, pr *target.PullRequest, logger *slog.Logger) map[string]any {
	acts := &builtinActions{
		ctx,
		scmClient,
		targetEntity,
		pr,
		logger,
	}

	return map[string]any{
		"$addLabels": acts.AddLabels,
		"$comment":   acts.Comment,
		"$review":    acts.Review,
	}
}

type builtinActions struct {
	ctx          context.Context
	scmClient    target.Target
	targetEntity *target.Entity
	pr           *target.PullRequest
	logger       *slog.Logger
}
