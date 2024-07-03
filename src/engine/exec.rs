use super::interpreter::Interpreter;
use super::errors::Error;
use super::lang::{ReviewpadFile, PadWorkflowRunBlock};

impl Interpreter{
    fn exec_configuration_file(&mut self, file: &ReviewpadFile) -> Result<(), Error> {
        for group in file.groups.clone() {
            let result = self.process_group(group.name, group.spec);
            if result.is_err() {
                return Err(Error::UnableToProcessGroup);
            }
        }

        Ok(())
    }

    fn exec_run(&mut self, run: &PadWorkflowRunBlock) -> Result<(), Error> {
        if run.for_each.is_none() && run.if_rules.is_empty() {
            return self.exec_actions(&run.actions)
        }

        for rule in &run.if_rules {
            let result = self.eval_expr(&rule.spec);
            if result.is_err() {
                return Err(Error::UnableToEvaluateRule)
            }

            if result.ok().unwrap() {
                if run.then_blocks.len() > 0 {
                    self.exec_statement_block(&run.then_blocks)?
                }

                return Ok(())
            }

            if !run.else_blocks.is_empty() {
                return self.exec_statement_block(&run.else_blocks)
            }
        }

        Ok(())
    }

    fn exec_actions(&mut self, actions: &Vec<String>) -> Result<(), Error> {
        for action in actions {
            let result = self.exec_statement(&action);
            if result.is_err() {
                return Err(Error::UnableToExecuteAction)
            }
        }

        Ok(())
    }

    fn exec_statement_block(&mut self, runs: &Vec<PadWorkflowRunBlock>) -> Result<(), Error> {
        for run in runs {
            self.exec_run(run)?;
        }

        Ok(())
    }
}
