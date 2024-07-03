use octocrab::{OctocrabBuilder, Error, Octocrab};
use crate::codehost::target::entity::Entity;

pub struct GithubClient {
    pub client: Octocrab,
    pub entity: Entity,
}

impl GithubClient {
    pub fn new(token: String, entity: Entity) -> Result<GithubClient, Error> {
        let client = OctocrabBuilder::new().personal_token(token).build()?;

        Ok(GithubClient {
            client,
            entity,
        })
    }

    pub fn get_entity(&self) -> Entity {
        self.entity.clone()
    }
}
