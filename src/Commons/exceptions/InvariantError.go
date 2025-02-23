package exceptions

type InvariantError struct {
	ClientError
}

func NewInvariantError(message string) *InvariantError {
	return &InvariantError{
		ClientError: *NewClientError(message, nil),
	}
}
