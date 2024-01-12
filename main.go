package main

import "fmt"

// MessageSender = sebuah interface yang mendefinisikan kontrak untuk mengirim pesan.
type MessageSender interface {
	SendMessage(message string)
}

// EmailSender = implementasi konkret dari interface MessageSender.
type EmailSender struct{}

// SendMessage = method dari EmailSender yang mengirim pesan email.
func (e *EmailSender) SendMessage(message string) {
	fmt.Println("Mengirim email:", message)
}

// SMS = sebuah struct yang merepresentasikan pesan SMS.
type SMS struct {
	Sender MessageSender // Dependency Inversion: SMS bergantung pada interface MessageSender, bukan implementasi konkret.
}

// SendSMS = method dari struct SMS yang mengirim pesan SMS menggunakan MessageSender yang ada.
func (s *SMS) SendSMS(message string) {
	s.Sender.SendMessage(message)
}

func main() {
	emailSender := &EmailSender{} // Dependency Injection: EmailSender diinjeksikan ke dalam struct SMS.
	sms := &SMS{Sender: emailSender}
	sms.SendSMS("Hello, world!")
}
