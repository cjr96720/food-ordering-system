package domains

type HealthCheckResponse struct {
	Messaage string `json:"message"`
}

type DefaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
