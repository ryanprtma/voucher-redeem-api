package exceptions

type NotFoundError struct {
	ClientError
}

func NewNotFoundError(message string) *NotFoundError {
	defaultStatusCode := 404
	clientError := *NewClientError(message, &defaultStatusCode)
	return &NotFoundError{
		clientError,
	}
}
