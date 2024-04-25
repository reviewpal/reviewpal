package gitlab

import (
	"context"

	"github.com/reviewpal/reviewpal/codehost/gitlab/converters"
	"github.com/reviewpal/reviewpal/codehost/target"
	"github.com/xanzy/go-gitlab"
)

func (c *GitlabClient) GetPullRequestPatch(ctx context.Context) (target.Patch, error) {
	patchMap := make(map[string]*target.File)

	files, err := c.GetPullRequestFiles(ctx)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		patchFile, err := target.NewFile(&file)
		if err != nil {
			return nil, err
		}

		patchMap[file.Filename] = patchFile
	}

	return patchMap, nil
}

func (c *GitlabClient) GetPullRequest(ctx context.Context) (*target.PullRequest, error) {
	pr, _, err := c.client.MergeRequests.GetMergeRequest(c.ProjectID(), c.tgt.Number, nil)
	if err != nil {
		return nil, err
	}

	return converters.PullRequestFromGitlabPullRequest(pr), nil
}

func (c *GitlabClient) GetPullRequestFiles(ctx context.Context) ([]target.CommitFile, error) {
	fs, err := paginatedRequest(
		func() interface{} {
			return []*gitlab.MergeRequestDiff{}
		},
		func(i interface{}, page int) (interface{}, *gitlab.Response, error) {
			fls := i.([]*gitlab.MergeRequestDiff)
			fs, resp, err := c.client.MergeRequests.ListMergeRequestDiffs(c.ProjectID(), c.tgt.Number, &gitlab.ListMergeRequestDiffsOptions{
				ListOptions: gitlab.ListOptions{
					Page:    page,
					PerPage: maxPerPage,
				},
			})
			if err != nil {
				return nil, nil, err
			}
			fls = append(fls, fs...)
			return fls, resp, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return converters.CommitFilesFromGitlabMergeRequestDiffs(fs.([]*gitlab.MergeRequestDiff)), nil
}
