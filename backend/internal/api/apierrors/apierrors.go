package apierrors

type UserError struct {
	ErrType string `json:"type"`
	Message string `json:"message"`
}

type TaskError struct {
	ErrType string `json:"type"`
	Message string `json:"message"`
}
