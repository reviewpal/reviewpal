package github

import (
	"context"
)

func (c *GithubClient) AddLabels(ctx context.Context, labels []string) ([]string, error) {
	ls, _, err := c.client.Issues.AddLabelsToIssue(ctx, c.tgt.Owner, c.tgt.Repo, c.tgt.Number, labels)
	if err != nil {
		return nil, err
	}

	labels = make([]string, len(ls))
	for i, l := range ls {
		labels[i] = l.GetName()
	}

	return labels, nil
}
