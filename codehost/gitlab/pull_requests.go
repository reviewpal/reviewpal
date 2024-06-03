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

func (c *GitlabClient) Comment(ctx context.Context, comment string) (*target.Comment, error) {
	note, _, err := c.client.Notes.CreateMergeRequestNote(c.ProjectID(), c.tgt.Number, &gitlab.CreateMergeRequestNoteOptions{
		Body: gitlab.Ptr(comment),
	})
	if err != nil {
		return nil, err
	}

	return converters.CommentFromGitlabNote(note), err
}

func (c *GitlabClient) Review(ctx context.Context, event target.ReviewState, body string) (*target.Review, error) {
	if event == target.Approved {
		approval, _, err := c.client.MergeRequestApprovals.ApproveMergeRequest(c.ProjectID(), c.tgt.Number, nil)
		if err != nil {
			return nil, err
		}
		return converters.ReviewFromGitlabApproval(approval), nil
	}

	if event == target.ChangesRequested {
		_, err := c.client.MergeRequestApprovals.UnapproveMergeRequest(c.ProjectID(), c.tgt.Number)
		if err != nil {
			return nil, err
		}

		return &target.Review{
			State: target.ChangesRequested,
		}, nil
	}

	return nil, nil
}
