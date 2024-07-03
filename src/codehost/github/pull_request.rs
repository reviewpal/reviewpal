use crate::codehost::github::client::GithubClient;
use crate::codehost::target::pull_request::PullRequest;
use crate::codehost::target::file::CommitFile;
use crate::codehost::github::errors::Error;
use crate::codehost::github::converters::pull_request;
use crate::codehost::github::converters::file;
use futures_util::StreamExt;

impl GithubClient {
    pub async fn get_pull_request(&self) -> Result<PullRequest, Error> {
        let pr = self.client.pulls(&self.entity.owner, &self.entity.repo).get(self.entity.number).await;
        if pr.is_err() {
            return Err(Error::UnableToGetPullRequest);
        }

        Ok(pull_request::pull_request_from_github_pull_request(pr.ok().unwrap()))
    }
}
