// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package engine

import (
	"context"
	"log/slog"

	"github.com/reviewpal/reviewpal/codehost/target"
	"gopkg.in/yaml.v3"
)

func Load(ctx context.Context, logger *slog.Logger, scmClient target.Target, data []byte) (*ReviewpadFile, error) {
	file, err := parse(data)
	if err != nil {
		return nil, err
	}

	return normalizeInlineRules(file)
}

func parse(data []byte) (*ReviewpadFile, error) {
	file := ReviewpadFile{}
	err := yaml.Unmarshal([]byte(data), &file)
	if err != nil {
		return nil, err
	}

	return &file, nil
}
