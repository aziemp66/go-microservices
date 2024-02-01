package http_server

type errResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
	Value   any    `json:"value,omitempty"`
}
