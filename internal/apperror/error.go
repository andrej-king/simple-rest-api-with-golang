package apperror

import "encoding/json"

// ErrNotFound custom error not found.
var (
	ErrNotFound = NewAppError(nil, "not found", "", "ERR-000001")
)

// AppError custom err with same interface as 'error'
type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

// Error interface error
func (e *AppError) Error() string {
	return e.Message
}

// Unwrap returns the next error in the error chain.
// If there is no next error, Unwrap returns nil.
func (e *AppError) Unwrap() error {
	return e.Err
}

// Marshal error in json format
func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

// NewAppError custom error
func NewAppError(err error, message, developerMessage, code string) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

// systemError cast system error to own
func systemError(err error) *AppError {
	return NewAppError(err, "internal system error", err.Error(), "ERR-000000")
}
