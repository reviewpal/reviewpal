package gitlab

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/reviewpal/reviewpal/codehost/target"
	"github.com/xanzy/go-gitlab"
)

const (
	maxPerPage = 20
)

type GitlabClient struct {
	client *gitlab.Client
	token  string
	logger *slog.Logger
	tgt    *target.Entity
}

func New(ctx context.Context, logger *slog.Logger, tgt *target.Entity, token string) (*GitlabClient, error) {
	cl, err := gitlab.NewClient(token)
	if err != nil {
		return nil, err
	}

	client := &GitlabClient{
		client: cl,
		token:  token,
		logger: logger,
		tgt:    tgt,
	}

	return client, nil
}

func (c *GitlabClient) ProjectID() string {
	return fmt.Sprintf("%s/%s", c.tgt.Owner, c.tgt.Repo)
}
