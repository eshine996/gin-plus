package gerror

type iStack interface {
	Error() string
	Stack() string
}

func New(err interface{}) error {
	if e, ok := err.(iStack); ok {
		return e
	}

	if e, ok := err.(string); ok {
		return &Error{
			stack: callers(),
			text:  e,
		}
	}

	if e, ok := err.(error); ok {
		return &Error{
			stack: callers(),
			text:  e.Error(),
		}
	}
	return nil
}

func Stack(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(iStack); ok {
		return e.Stack()
	}
	return err.Error()
}
