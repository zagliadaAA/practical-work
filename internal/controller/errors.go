package controller

type ValidationError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func NewValidationError(param string, message string) *ValidationError {
	return &ValidationError{
		Param:   param,
		Message: message,
	}

}
