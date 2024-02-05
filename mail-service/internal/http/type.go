package http_server

type errResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type Response struct {
	Message string `json:"message"`
	Value   any    `json:"value,omitempty"`
}
