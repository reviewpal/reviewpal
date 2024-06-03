package converters

import (
	"github.com/google/go-github/v52/github"
	"github.com/reviewpal/reviewpal/codehost/target"
)

var githubStateToReviewState = map[string]target.ReviewState{
	"APPROVED":          target.Approved,
	"CHANGES_REQUESTED": target.ChangesRequested,
	"COMMENTED":         target.Commented,
}

func githubReviewStateToState(s string) target.ReviewState {
	return githubStateToReviewState[s]
}

func ReviewFromGithubReview(r *github.PullRequestReview) *target.Review {
	return &target.Review{
		ID:    r.GetID(),
		State: githubReviewStateToState(r.GetState()),
	}
}
