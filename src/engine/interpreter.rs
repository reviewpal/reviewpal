use std::fmt;

use crate::codehost::target;
use evalexpr::{ContextWithMutableVariables, EvalexprError, HashMapContext, Value, ValueType};
use crate::engine::errors::Error as EngineError;

pub struct Interpreter {
    pr: target::pull_request::PullRequest,
    context: HashMapContext,
    tgt: Box<dyn target::Target>,
}

impl Interpreter {
    pub fn new(tgt: Box<dyn target::Target>) -> Result<Interpreter, EngineError> {
        let context = HashMapContext::new();
        let pr = tgt.get_pull_request().ok().unwrap();

        Ok(Interpreter {
            context,
            pr,
            tgt,
        })
    }
}

impl Interpreter {
    pub fn process_group(&mut self, group_name: String, ex: String) -> Result<(), EvalexprError> {
        let res = evalexpr::eval_with_context(&ex, &self.context)?;
        if !res.is_tuple() {
            return Err(EvalexprError::type_error(res, vec![ValueType::Tuple]));
        }

        self.context.set_value(group_name, res)?;

        Ok(())
    }

    pub fn eval_expr(&self, ex: &str) -> Result<bool, EvalexprError> {
        let res = evalexpr::eval_with_context(&ex, &self.context)?;
        if !res.is_boolean() {
            return Err(EvalexprError::type_error(res, vec![ValueType::Boolean]));
        }

        Ok(res.as_boolean().unwrap())
    }

    pub fn exec_statement(&self, stmt: &str) -> Result<(), EvalexprError> {
        evalexpr::eval_with_context(stmt, &self.context)?;

        Ok(())
    }
}
