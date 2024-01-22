package http_server

type (
	Response struct {
		Message string      `json:"message" binding:"required"`
		Value   interface{} `json:"value"`
	}

	Error struct {
		Message string `json:"message" binding:"required"`
	}
)
