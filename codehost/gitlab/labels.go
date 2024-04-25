package gitlab

import (
	"context"

	"github.com/xanzy/go-gitlab"
)

func (c *GitlabClient) AddLabels(ctx context.Context, labels []string) ([]string, error) {
	ls, _, err := c.client.MergeRequests.UpdateMergeRequest(c.ProjectID(), c.tgt.Number, &gitlab.UpdateMergeRequestOptions{
		AddLabels: (*gitlab.LabelOptions)(&labels),
	})
	if err != nil {
		return nil, err
	}

	return ls.Labels, nil
}
