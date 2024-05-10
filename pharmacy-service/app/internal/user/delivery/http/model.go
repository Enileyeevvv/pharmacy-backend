package http

type SignUpRequest struct {
	Login    string `json:"login" validate:"required,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

type SignUpResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message,omitempty"`
	Msg     map[string]string `json:"msg,omitempty"`
}
