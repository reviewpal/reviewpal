// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package prmate

import (
	"bytes"
	"context"
	"log/slog"

	"github.com/reviewpal/reviewpal/codehost/target"
	"github.com/reviewpal/reviewpal/engine"
)

func Load(ctx context.Context, log *slog.Logger, scmClient target.Target, buf *bytes.Buffer) (*engine.ReviewpadFile, error) {
	file, err := engine.Load(ctx, log, scmClient, buf.Bytes())
	if err != nil {
		return nil, err
	}

	log.With("file", file).Debug("loaded reviewpad file")

	reserved := []string{}

	err = engine.Lint(file, reserved, log)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func Run(
	ctx context.Context,
	logger *slog.Logger,
	scmClient target.Target,
	targetEntity *target.Entity,
	reviewpadFile *engine.ReviewpadFile,
) error {
	interpreter, err := engine.NewInterpreter(ctx, logger, scmClient, targetEntity)
	if err != nil {
		logger.With("err", err).Error("unable to create interpreter")
		return err
	}

	if err = interpreter.ExecConfigurationFile(reviewpadFile); err != nil {
		logger.With("err", err).Error("unable to execute reviewpad file")
	}

	return err
}
