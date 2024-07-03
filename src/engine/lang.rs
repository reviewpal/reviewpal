use std::collections::HashMap;

#[derive(PartialEq, Eq, Clone)]
struct PadRule {
    pub name: String,
    pub description: String,
    pub spec: String,
}

#[derive(PartialEq, Eq, Clone)]
struct PadWorkflowRule {
    pub name: String,
    pub spec: String,
}

#[derive(PartialEq, Eq, Clone)]
struct PadLabel {
    name: String,
    color: String,
    description: String,
}

#[derive(PartialEq, Eq, Clone)]
pub struct PadWorkflow {
    pub name: String,
    pub description: String,
    pub rules: Vec<PadWorkflowRule>,
    pub runs: Vec<PadWorkflowRunBlock>,
    pub non_normalized_rules: Option<serde_yaml::Value>,
    pub non_normalized_actions: Option<serde_yaml::Value>,
    pub non_normalized_else: Option<serde_yaml::Value>,
    pub non_normalized_run: Option<serde_yaml::Value>,
}

#[derive(PartialEq, Eq, Clone)]
pub struct PadWorkflowRunBlock {
    pub if_rules: Vec<PadWorkflowRule>,
    pub then_blocks: Vec<PadWorkflowRunBlock>,
    pub else_blocks: Vec<PadWorkflowRunBlock>,
    pub actions: Vec<String>,
    pub for_each: Option<PadWorkflowRunForEachBlock>,
}

#[derive(PartialEq, Eq, Clone)]
struct PadWorkflowRunForEachBlock {
    key: String,
    value: String,
    in_value: String,
    do_blocks: Vec<PadWorkflowRunBlock>,
}

#[derive(PartialEq, Eq, Clone)]
struct PadGroup {
    pub name: String,
    pub spec: String,
}

#[derive(PartialEq, Eq, Clone)]
pub struct ReviewpadFile {
    pub groups: Vec<PadGroup>,
    pub rules: Vec<PadRule>,
    pub labels: HashMap<String, PadLabel>,
    pub workflows: Vec<PadWorkflow>,
}
