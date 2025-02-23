package exceptions

type ClientError struct {
	Message    string
	StatusCode int
}

func NewClientError(message string, statusCode *int) *ClientError {
	defaultStatusCode := 400

	if statusCode != nil {
		defaultStatusCode = *statusCode
	}

	return &ClientError{
		Message:    message,
		StatusCode: defaultStatusCode,
	}
}

func (e *ClientError) Error() string {
	return e.Message
}
