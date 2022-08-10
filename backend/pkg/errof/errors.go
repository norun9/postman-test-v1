package errof

type UserErr string
type InternalErr string

func (e UserErr) Error() (msg string) {
	var ok bool
	if msg, ok = ErrCodeNames[e]; !ok {
		return string(e)
	}
	return msg
}

func (e InternalErr) Error() (msg string) {
	var ok bool
	if msg, ok = InternalErrCodeNames[e]; !ok {
		return string(e)
	}
	return msg
}

// InternalErrCodeNames :
var InternalErrCodeNames = map[InternalErr]string{
	ErrInternal: "internal error has occurred.",
	ErrDatabase: "inconsistency has occurred in the database.",
}

// CmdErrCodeNames :
var ErrCodeNames = map[UserErr]string{
	ErrNoPost: "post model was not found.",
}

var (
	// Internal Err
	ErrInternal InternalErr = "ErrInternal"
	ErrDatabase InternalErr = "ErrDatabase"
	// User Err
	ErrNoPost UserErr = "ErrNoPost"
)
