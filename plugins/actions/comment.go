// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package actions

import "github.com/reviewpal/reviewpal/codehost/target"

func (a *builtinActions) Comment(comment string) (*target.Comment, error) {
	return a.scmClient.Comment(a.ctx, comment)
}
