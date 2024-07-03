use octocrab::models::repos::DiffEntry as GithubCommitFile;
use crate::codehost::target::file::CommitFile;

pub fn commit_file_from_github_commit_file(file: GithubCommitFile) -> CommitFile {
        CommitFile {
            filename: file.filename.clone(),
            patch: file.patch.clone().unwrap_or("".to_string()),
        }
}
