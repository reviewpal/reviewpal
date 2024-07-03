pub enum ForEachError {
    MustBeAMap,
    ValueMustBeAString,
    InMustBeAString,
    DoBlockIsRequired,
}

pub enum Error {
    UnableToProcessGroup,
    UnableToProcessRule,
    UnableToExecuteAction,
    UnableToEvaluateRule,
    UnableToExecuteStatementBlock,
    ForEachError(ForEachError),
}
