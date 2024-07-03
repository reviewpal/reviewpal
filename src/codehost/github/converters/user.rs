use octocrab::models::UserProfile as GithubUser;
use crate::codehost::target::user::User;

pub fn user_from_github_user(u: GithubUser) -> User {
    User {
        id: u.id.into_inner(),
        login: u.login,
    }
}
