// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package actions

import (
	"fmt"

	"github.com/reviewpal/reviewpal/codehost/target"
)

func (a *builtinActions) Review(reviewEvent target.ReviewState, reviewBody string) (*target.Review, error) {
	if a.pr.IsDraft {
		a.logger.Info("skipping review because the pull request is in draft")
		return nil, nil
	}

	if a.pr.Status == target.Closed {
		a.logger.Info("skipping review because the pull request is closed")
		return nil, nil
	}

	if reviewEvent != target.Approved && reviewBody == "" {
		return nil, fmt.Errorf("review comment required in %s state", reviewEvent)
	}

	return a.scmClient.Review(a.ctx, reviewEvent, reviewBody)
}
