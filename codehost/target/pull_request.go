package target

type PullRequestStatus string

const (
	Open   PullRequestStatus = "open"
	Closed PullRequestStatus = "closed"
	Merged PullRequestStatus = "merged"
	Locked PullRequestStatus = "locked"
)

type PullRequest struct {
	ID      int64
	Number  int
	IsDraft bool
	Status  PullRequestStatus
	Base    Branch
	Head    Branch
	User    User
}
