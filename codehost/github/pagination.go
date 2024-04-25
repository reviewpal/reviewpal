// Copyright (C) 2022 Explore.dev, Unipessoal Lda - All Rights Reserved
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package github

import (
	"github.com/google/go-github/v52/github"
)

func paginatedRequest(
	initFn func() interface{},
	reqFn func(interface{}, int) (interface{}, *github.Response, error),
) (interface{}, error) {
	page := 1
	results, resp, err := reqFn(initFn(), page)
	if err != nil {
		return nil, err
	}

	numPages := resp.LastPage
	for page <= numPages && resp.NextPage > page {
		page++
		results, resp, err = reqFn(results, page)
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}
