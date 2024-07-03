use octocrab::models::commits::Commit as GithubCommit;
use crate::codehost::target::commit::Commit;

pub fn commit_from_github_commit(c: GithubCommit) -> Commit {
    Commit {
        sha: c.sha,
    }
}
