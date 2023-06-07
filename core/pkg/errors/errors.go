package errors

type BaseError struct {
	Message string
	Err     error
}

func (e *BaseError) Error() string {
	return e.Message
}

func (e *BaseError) Unwrap() error {
	return e.Err
}

type HashingError struct {
	BaseError
}

func NewHashingError(message string, err error) *HashingError {
	return &HashingError{
		BaseError: BaseError{
			Message: message,
			Err:     err,
		},
	}
}

type DatabaseError struct {
	BaseError
}

func NewDatabaseError(message string, err error) *DatabaseError {
	return &DatabaseError{
		BaseError: BaseError{
			Message: message,
			Err:     err,
		},
	}
}

type AuthenticationError struct {
	BaseError
}

func NewAuthenticationError(message string, err error) *AuthenticationError {
	return &AuthenticationError{
		BaseError: BaseError{
			Message: message,
			Err:     err,
		},
	}
}
