// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package engine

// ExecConfigurationFile processes
// Pre-condition Lint(file) == nil
func (i *Interpreter) ExecConfigurationFile(file *ReviewpadFile) error {
	log := i.logger
	rules := make(map[string]PadRule)

	// process groups
	for _, group := range file.Groups {
		err := i.ProcessGroup(
			group.Name,
			group.Spec,
		)
		if err != nil {
			return err
		}
	}

	// process rules
	for _, rule := range file.Rules {
		err := i.ProcessRule(rule.Name, rule.Spec)
		if err != nil {
			return err
		}
		rules[rule.Name] = rule
	}

	// process workflows
	for _, workflow := range file.Workflows {
		log.With("name", workflow.Name).Info("executing workflow")

		for _, run := range workflow.Runs {
			err := i.execStatement(run, rules)
			if err != nil {
				return err
			}
		}
	}

	for _, pipeline := range file.Pipelines {
		pipelineLog := log.With("pipeline", pipeline.Name)

		pipelineLog.With("name", pipeline.Name).Info("executing pipeline")

		var err error
		activated := pipeline.Trigger == ""
		if !activated {
			activated, err = i.EvalExpr(pipeline.Trigger)
			if err != nil {
				return err
			}
		}

		if !activated {
			pipelineLog.Info("skipping pipeline because the trigger condition was not met")
			continue
		}

		for num, stage := range pipeline.Stages {
			pipelineLog.With("stage", num).Info("processing pipeline stage")

			if stage.Until == "" {
				pipelineLog.Info("pipeline stage 'until' condition not specified, executing actions")

				err = i.execActions(stage.Actions)
				if err != nil {
					return err
				}

				break
			}

			isStageCompleted, err := i.EvalExpr(stage.Until)
			if err != nil {
				return err
			}

			if isStageCompleted {
				pipelineLog.Info("pipeline stage `until` condition was met, skipping stage")
				continue
			}

			pipelineLog.Info("pipeline stage `until` condition was not met, executing actions")

			err = i.execActions(stage.Actions)
			if err != nil {
				return err
			}

			// If the stage was been executed, the pipeline should stop
			break
		}
	}

	return nil
}

func (i *Interpreter) execActions(actions []string) error {
	for _, action := range actions {
		if err := i.ExecStatement(action); err != nil {
			return err
		}
	}

	return nil
}

func (i *Interpreter) execStatement(run PadWorkflowRunBlock, rules map[string]PadRule) error {
	// if the run block was just a simple string
	// there is no rule to evaluate, so just return the actions
	if run.ForEach == nil && run.If == nil {
		// execute the actions
		return i.execActions(run.Actions)
	}

	for _, rule := range run.If {
		ruleName := rule.Rule
		ruleDefinition := rules[ruleName]

		thenClause, err := i.EvalExpr(ruleDefinition.Spec)
		if err != nil {
			return err
		}

		if thenClause {
			if len(run.Then) > 0 {
				err := i.execStatementBlock(run.Then, rules)
				if err != nil {
					return err
				}
			}

			return nil
		}

		if run.Else != nil {
			return i.execStatementBlock(run.Else, rules)
		}
	}

	return nil
}

func (i *Interpreter) execStatementBlock(runs []PadWorkflowRunBlock, rules map[string]PadRule) error {
	for _, run := range runs {
		err := i.execStatement(run, rules)
		if err != nil {
			return err
		}
	}

	return nil
}
