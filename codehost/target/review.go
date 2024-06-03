package target

type ReviewState int

const (
	Approved = iota
	ChangesRequested
	Commented
)

func (s ReviewState) String() string {
	if s == Approved {
		return "APPROVE"
	}

	if s == ChangesRequested {
		return "REQUEST_CHANGES"
	}

	return "COMMENT"
}

type Review struct {
	ID    int64
	State ReviewState
}
