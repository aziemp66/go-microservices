package request

type CreateLogEntry struct {
	Name string `json:"name" validate:"required,gt=1"`
	Data any    `json:"data" validate:"required"`
}
