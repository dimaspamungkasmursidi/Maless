package helpers

type CustomError struct {
	message string
	code    int
}

func NewCustomError(message string, code int) *CustomError {
	return &CustomError{
		message: message,
		code:    code,
	}
}
