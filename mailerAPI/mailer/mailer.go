package mailer 

import (
	"log"
	"net/smtp"
	"main"
)

func sendEmail(sendTo string, subject string, msg Message) {
	auth := smtp.PlainAuth("", "email@gmail.com", "secret_pass", "smtp.mailtrap.io")

	to := []sendTo
	msg := []byte("To: "+ sendTo + "\r\n" + 
		"Subject: " + subject + "\r\n" + 
		"\r\n" + 
		msg.Message + "\r\n"
	)
	err := stmp.SendMail("smtp.mailtrap.io:25", auth, "piotr@mailtrap.io", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}