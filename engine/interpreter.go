// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package engine

import (
	"context"
	"fmt"
	"log/slog"
	"reflect"

	"github.com/expr-lang/expr"
	"github.com/reviewpal/reviewpal/codehost/target"
	"github.com/reviewpal/reviewpal/plugins/actions"
	"github.com/reviewpal/reviewpal/plugins/functions"
	"github.com/samber/lo"
)

type Interpreter struct {
	ctx         context.Context
	scmClient   target.Target
	target      *target.Entity
	logger      *slog.Logger
	registerMap map[string]any
	patch       target.Patch
	pr          *target.PullRequest
	actions     map[string]any
	functions   map[string]any
}

func (i *Interpreter) ProcessGroup(groupName string, ex string) error {
	program, err := expr.Compile(ex, expr.AsKind(reflect.Slice))
	if err != nil {
		return err
	}

	res, err := expr.Run(program, i.functions)
	if err != nil {
		return err
	}

	i.registerMap[groupName] = res
	return nil
}

func BuildInternalRuleName(name string) string {
	return fmt.Sprintf("@rule:%v", name)
}

func (i *Interpreter) ProcessRule(name, spec string) error {
	internalRuleName := BuildInternalRuleName(name)

	i.registerMap[internalRuleName] = spec
	return nil
}

func (i *Interpreter) EvalExpr(ex string) (bool, error) {
	program, err := expr.Compile(ex, expr.AsBool())
	if err != nil {
		return false, err
	}

	res, err := expr.Run(program, i.functions)
	if err != nil {
		return false, err
	}

	return res.(bool), nil
}

func (i *Interpreter) ExecStatement(statement string) error {
	program, err := expr.Compile(statement)
	if err != nil {
		return err
	}

	_, err = expr.Run(program, lo.Assign(i.functions, i.actions))

	return err
}

func (i *Interpreter) GetTarget() *target.Entity {
	return i.target
}

func (i *Interpreter) GetLogger() *slog.Logger {
	return i.logger
}

func NewInterpreter(
	ctx context.Context,
	logger *slog.Logger,
	scmClient target.Target,
	targetEntity *target.Entity,
) (*Interpreter, error) {
	pr, err := scmClient.GetPullRequest(ctx)
	if err != nil {
		return nil, err
	}

	patch, err := scmClient.GetPullRequestPatch(ctx)
	if err != nil {
		return nil, err
	}

	funcs := functions.New(ctx, scmClient, targetEntity, pr)
	actions := actions.New(ctx, scmClient, targetEntity, pr)

	return &Interpreter{
		ctx:         ctx,
		scmClient:   scmClient,
		target:      targetEntity,
		patch:       patch,
		pr:          pr,
		logger:      logger,
		registerMap: make(map[string]any),
		functions:   funcs,
		actions:     actions,
	}, nil
}
