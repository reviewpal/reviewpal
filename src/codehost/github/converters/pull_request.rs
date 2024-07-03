use octocrab::models::pulls::PullRequest as GithubPullRequest;
use octocrab::models::IssueState;
use crate::codehost::target::pull_request::{PullRequest, PullRequestState};
use crate::codehost::target::branch::Branch;
use crate::codehost::target::user::User;

fn pull_request_state_from_issue_state(s: IssueState) -> PullRequestState {
    match s {
        IssueState::Open => PullRequestState::Open,
        IssueState::Closed => PullRequestState::Closed,
        _ => PullRequestState::Open,
    }
}

pub fn pull_request_from_github_pull_request(p: GithubPullRequest) -> PullRequest {
    let user = p.user.unwrap();

    PullRequest {
        id: p.id.into_inner(),
        number: p.number,
        is_draft: p.draft.unwrap_or(false),
        status: pull_request_state_from_issue_state(p.state.unwrap()),
        head: Branch{
            name: p.head.label.unwrap(),
            sha: p.head.sha,
        },
        base: Branch {
            name: p.base.label.unwrap(),
            sha: p.base.sha,
        },
        user: User {
            id: user.id.into_inner(),
            login: user.login,
        },
    }
}
