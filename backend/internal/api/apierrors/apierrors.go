package apierrors

type APIError struct {
	RequestID string `json:"request_id"`
	ErrType   string `json:"type"`
	Message   string `json:"message"`
}
