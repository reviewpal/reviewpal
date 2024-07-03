pub mod branch;
pub mod comment;
pub mod commit;
pub mod label;
pub mod pull_request;
pub mod user;
pub mod entity;
pub mod file;
mod errors;

pub trait Target {
    fn get_entity(&self) -> entity::Entity;
    fn get_pull_request(&self) -> Result<pull_request::PullRequest, errors::Error>;
    fn comment(&self, comment: String);
}
