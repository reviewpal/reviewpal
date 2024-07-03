use super::{errors::{Error, ForEachError}, lang::{PadWorkflow, PadWorkflowRunBlock, ReviewpadFile}};

pub fn normalize(file: ReviewpadFile) -> Result<ReviewpadFile, Error> {
    let mut normalized_file = ReviewpadFile{
        groups: file.groups,
        rules: file.rules,
        labels: file.labels,
        workflows: file.workflows,
    };

    Ok(normalized_file)
}

fn process_workflow(workflow: PadWorkflow) -> Result<PadWorkflow, Error> {
    let normalized_workflow = PadWorkflow{
        name: workflow.name,
        description: workflow.description,
        rules: workflow.rules,
    };

    Ok(normalized_workflow)
}

fn process_run(run: serde_yaml::Value) -> Result<Vec<PadWorkflowRunBlock>, Error> {
    let mut block = PadWorkflowRunBlock{
        if_rules: vec![],
        then_blocks: vec![],
        else_blocks: vec![],
        for_each: None,
        actions: vec![]
    };
    match run {
        serde_yaml::Value::String(val) => {
            block.actions = vec![val];
            Ok(vec![block])
        }
        serde_yaml::Value::Mapping(map) => {
            if let Some(for_each_block) = map.get(String::from("forEach")) {
                if !for_each_block.is_mapping() {
                    Err(Error::ForEachError(ForEachError::MustBeAMap));
                }

                let for_each_block = for_each_block.as_mapping().unwrap();
                let value = for_each_block.get("value");
                let iterables = for_each_block.get("in");
                let actions = for_each_block.get("do");

                if value.is_none() || !value.unwrap().is_string() {
                    Err(Error::ForEachError(ForEachError::ValueMustBeAString));
                }

                if iterables.is_none() || !iterables.unwrap().is_string() {
                    Err(Error::ForEachError(ForEachError::InMustBeAString));
                }

                if actions.is_none() {
                    Err(Error::ForEachError(ForEachError::DoBlockIsRequired));
                }

                let processed_do = process_run(actions.unwrap().to_owned());
                if processed_do.is_err() {
                    Err(Error::UnableToProcessGroup);
                }


            }
        }
    }
}
