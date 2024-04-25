// Copyright (C) 2022 Explore.dev, Unipessoal Lda - All Rights Reserved
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package github

import (
	"context"
	"fmt"
	"io"

	"github.com/google/go-github/v52/github"
	"github.com/reviewpal/reviewpal/codehost/target"
)

type DownloadMethod string

const (
	DownloadMethodSHA        DownloadMethod = "withSHA"
	DownloadMethodBranchName DownloadMethod = "withBranchName"
)

type DownloadContentsOptions struct {
	Method DownloadMethod
}

func (c *GithubClient) DownloadContents(ctx context.Context, tgt *target.Entity, filePath string, branch *target.Branch, options *DownloadContentsOptions) ([]byte, error) {
	var branchRef string
	switch options.Method {
	case DownloadMethodSHA:
		branchRef = branch.SHA
	case DownloadMethodBranchName:
		branchRef = branch.Name
	default:
		return nil, fmt.Errorf("invalid download method specified")
	}

	ioReader, _, err := c.client.Repositories.DownloadContents(ctx, tgt.Owner, tgt.Repo, filePath, &github.RepositoryContentGetOptions{
		Ref: branchRef,
	})

	if err != nil {
		return nil, err
	}

	return io.ReadAll(ioReader)
}
