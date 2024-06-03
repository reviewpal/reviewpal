package main

import (
	"bytes"
	"context"
	"log/slog"
	"os"

	prmate "github.com/reviewpal/reviewpal"
	gh "github.com/reviewpal/reviewpal/codehost/github"
	"github.com/reviewpal/reviewpal/codehost/target"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	ctx := context.Background()
	tgt := &target.Entity{
		Owner:  "reviewpal",
		Repo:   "playground",
		Number: 1,
	}
	githubClient := gh.NewGithubClientFromToken(ctx, logger, "", tgt)
	// gitlabClient, err := gl.New(ctx, logger, tgt, "")

	file, err := os.ReadFile("./reviewpal.yml")
	if err != nil {
		slog.With("err", err).Error("unable to load reviewpal file")
		return
	}

	reviewpadFile, err := prmate.Load(ctx, logger, githubClient, bytes.NewBuffer(file))
	if err != nil {
		slog.With("err", err).Error("unable to parse reviewpal file")
		return
	}

	prmate.Run(ctx, logger, githubClient, tgt, reviewpadFile)
}
