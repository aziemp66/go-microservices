package mail_usecase

func NewEmailUsecase(
	domain string,
	host string,
	port int,
	username string,
	password string,
	encryption string,
	fromAddress string,
	fromName string,
) MailUsecase {
	return &Mail{
		Domain:      domain,
		Host:        host,
		Port:        port,
		Username:    username,
		Password:    password,
		Encryption:  encryption,
		FromAddress: fromAddress,
		FromName:    fromName,
	}
}
