package request

type SendEmail struct {
	From    *string `json:"from" validate:"required" id:"Pengirim" en:"From"`
	To      *string `json:"to" validate:"required" id:"Penerima" en:"To"`
	Subject *string `json:"subject" validate:"required" id:"Subjek" en:"Subject"`
	Message *string `json:"message" validate:"required" id:"Pesan" en:"Message"`
}
