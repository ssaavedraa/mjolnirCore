package interfaces

type Email struct {
	ReceiverAddress string
	SenderAddress   string
	Subject         string
	TemplateName    string
	Locale          string
	TemplateData    map[string]string
}

type EmailSender interface {
	Send(email Email) error
}
