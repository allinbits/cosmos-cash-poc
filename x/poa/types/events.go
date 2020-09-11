package types

// poa module event types
const (
	EventTypeVote = "vote"

	AttributeKeyCandidate = "name"

	// TODO: Some events may not have values for that reason you want to emit that something happened.
	// AttributeValueDoubleSign = "double_sign"

	AttributeValueCategory = ModuleName
)
