// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package actions

func (a *builtinActions) AddLabels(labels ...string) error {
	_, err := a.scmClient.AddLabels(a.ctx, labels)
	return err
}
