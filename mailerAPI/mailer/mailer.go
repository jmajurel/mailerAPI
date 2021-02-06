package mailer 

import (
	"log"
	"net/smtp"
	"main"
)

func sendEmail(sendTo string, subject string, msg string) {
	auth := smtp.PlainAuth("", senderEmailAccount, smtpSecretPass, smtpServerUrl)

	to := []sendTo
	msg := []byte("To: "+ sendTo + "\r\n" + 
		"Subject: " + subject + "\r\n" + 
		"\r\n" + 
		msg + "\r\n"
	)
	err := stmp.SendMail(smtpServerUrl + ":465", auth, senderEmailAccount, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}