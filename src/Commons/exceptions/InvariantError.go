package exceptions

type InvariantError struct {
	ClientError
}

func NewInvariantError(message string) *InvariantError {
	clientError := *NewClientError(message, nil)
	return &InvariantError{
		clientError,
	}
}
