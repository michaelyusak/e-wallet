package dto

type ValidationErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrResponse struct {
	Message string                   `json:"message"`
	Code    int                      `json:"code"`
	Details []ValidationErrorMessage `json:"details,omitempty"`
}
