package request

type CreateLogEntry struct {
	Name string `json:"name" validate:"required,gt=1" en:"Name" id:"Nama"`
	Data any    `json:"data" validate:"required" en:"Data" id:"Data"`
}

type UpdateLogEntry struct {
	Name string `json:"name" validate:"required,gt=1" en:"Name" id:"Nama"`
	Data any    `json:"data" validate:"required" en:"Data" id:"Data"`
}
