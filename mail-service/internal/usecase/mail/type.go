package mail_usecase

type MailUsecase interface {
	SendSMTPMessage(msg Message) error
}

type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        string
	DataMap     map[string]any
}
