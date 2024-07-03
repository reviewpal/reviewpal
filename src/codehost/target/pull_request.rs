use crate::codehost::target::branch::Branch;
use crate::codehost::target::user::User;

pub enum PullRequestState {
    Open,
    Closed,
    Merged,
    Locked,
}

pub struct PullRequest {
    pub id: u64,
    pub number: u64,
    pub is_draft: bool,
    pub status: PullRequestState,
    pub base: Branch,
    pub head: Branch,
    pub user: User,
}
