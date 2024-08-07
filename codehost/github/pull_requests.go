package github

import (
	"context"

	"github.com/google/go-github/v52/github"
	"github.com/reviewpal/reviewpal/codehost/github/converters"
	"github.com/reviewpal/reviewpal/codehost/target"
)

func (c *GithubClient) GetPullRequestPatch(ctx context.Context) (target.Patch, error) {
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

func (c *GithubClient) GetPullRequest(ctx context.Context) (*target.PullRequest, error) {
	pr, _, err := c.client.PullRequests.Get(ctx, c.tgt.Owner, c.tgt.Repo, c.tgt.Number)
	if err != nil {
		return nil, err
	}

	return converters.PullRequestFromGithubPullRequest(pr), nil
}

func (c *GithubClient) GetPullRequestFiles(ctx context.Context) ([]target.CommitFile, error) {
	fs, err := paginatedRequest(
		func() interface{} {
			return []*github.CommitFile{}
		},
		func(i interface{}, page int) (interface{}, *github.Response, error) {
			fls := i.([]*github.CommitFile)
			fs, resp, err := c.client.PullRequests.ListFiles(ctx, c.tgt.Owner, c.tgt.Repo, c.tgt.Number, &github.ListOptions{
				Page:    page,
				PerPage: maxPerPage,
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

	return converters.CommitFilesFromGithubCommitFiles(fs.([]*github.CommitFile)), nil
}

func (c *GithubClient) Comment(ctx context.Context, comment string) (*target.Comment, error) {
	cmt, _, err := c.client.Issues.CreateComment(ctx, c.tgt.Owner, c.tgt.Repo, c.tgt.Number, &github.IssueComment{
		Body: github.String(comment),
	})
	if err != nil {
		return nil, err
	}

	return converters.CommentFromGithubComment(cmt), nil
}

func (c *GithubClient) Review(ctx context.Context, reviewEvent target.ReviewState, reviewBody string) (*target.Review, error) {
	rvw, _, err := c.client.PullRequests.CreateReview(ctx, c.tgt.Owner, c.tgt.Repo, c.tgt.Number, &github.PullRequestReviewRequest{
		Body:  github.String(reviewBody),
		Event: github.String(reviewEvent.String()),
	})
	if err != nil {
		return nil, err
	}

	return converters.ReviewFromGithubReview(rvw), nil
}
