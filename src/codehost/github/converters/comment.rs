use octocrab::models::issues::Comment as GithubComment;
use crate::codehost::target::comment::Comment;

pub fn comment_from_github_comment(c: GithubComment) -> Comment {
    Comment {
        id: c.id.into_inner(),
        body: c.body.unwrap(),
    }
}
