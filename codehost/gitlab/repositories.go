package gitlab

import (
	"context"
	"fmt"

	"github.com/reviewpal/reviewpal/codehost/target"
	"github.com/xanzy/go-gitlab"
)

type DownloadMethod string

const (
	DownloadMethodSHA        DownloadMethod = "withSHA"
	DownloadMethodBranchName DownloadMethod = "withBranchName"
)

type DownloadContentsOptions struct {
	Method DownloadMethod
}

func (c *GitlabClient) DownloadContents(ctx context.Context, filePath string, branch *target.Branch, options *DownloadContentsOptions) ([]byte, error) {
	var branchRef string
	switch options.Method {
	case DownloadMethodSHA:
		branchRef = branch.SHA
	case DownloadMethodBranchName:
		branchRef = branch.Name
	default:
		return nil, fmt.Errorf("invalid download method specified")
	}

	file, _, err := c.client.RepositoryFiles.GetRawFile(c.ProjectID(), filePath, &gitlab.GetRawFileOptions{
		Ref: &branchRef,
	})

	if err != nil {
		return nil, err
	}

	return file, nil
}
