package target

type PullRequest struct {
	ID     int64
	Number int
	Base   Branch
	Head   Branch
	User   User
}
