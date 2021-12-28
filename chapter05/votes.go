package chapter05

// Voter defines someone who can be identified as a voter.
// Be that via name, or SSN, or other ID.
type Voter interface {
	~string
}

// Voters holds all the people who already voted.
type Voters[T Voter] map[T]struct{}

// New creates a new voter registrar.
func New[T Voter]() Voters[T] {
	return make(Voters[T])
}

// CheckVoted is a misleading name in the book. It's rather, RegisterVote because
// it also saves the vote and not just returns whether someone voted or not.
func (v Voters[T]) CheckVoted(id T) string {
	_, ok := v[id]
	if ok {
		return "already voted"
	}
	v[id] = struct{}{}
	return "thank you for your vote"
}
