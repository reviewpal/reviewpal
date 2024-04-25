// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file

package github

import (
	"context"
	"log/slog"

	"github.com/google/go-github/v52/github"
	"github.com/reviewpal/reviewpal/codehost/target"
	"golang.org/x/oauth2"
)

const (
	maxPerPage = 100
)

type GithubClient struct {
	client *github.Client
	token  string
	logger *slog.Logger
	tgt    *target.Entity
}

func NewGithubClientFromToken(ctx context.Context, logger *slog.Logger, token string, tgt *target.Entity) *GithubClient {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := &GithubClient{
		client: github.NewClient(tc),
		token:  token,
		logger: logger,
		tgt:    tgt,
	}

	return client
}
