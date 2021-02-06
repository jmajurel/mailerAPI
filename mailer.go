package main 

import (
	"log"
	"net/smtp"
)

func sendEmail(sendTo string, subject string, msg string) {

	var emailBody []byte
	auth := smtp.PlainAuth("", senderEmailAccount, smtpSecretPass, smtpServerUrl)

	to := []string{sendTo}
	emailBody = []byte("To: " + sendTo + "\r\n" + 
		"Subject: " + subject + "\r\n" + 
		"\r\n" + 
		msg + "\r\n")

	err := smtp.SendMail(smtpServerUrl + ":465", auth, senderEmailAccount, to, emailBody)
	if err != nil {
		log.Fatal(err)
	}
}