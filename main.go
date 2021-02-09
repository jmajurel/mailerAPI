package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
)

type Message struct {
	From    string `json:"from" validate:"required,email"`
	Message string `json:"message" validate:"required"`
}

var senderEmailAccount string
var receiverEmailAccount string
var smtpServerURL string
var smtpSecretPass string

var validate *validator.Validate

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func mailHandler(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body) //read request body

	sanitizer := bluemonday.UGCPolicy()

	var message Message
	json.Unmarshal(reqBody, &message) //convert to a message type
	err := validate.Struct(message)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	message.From, message.Message = sanitizer.Sanitize(message.From), sanitizer.Sanitize(message.Message)
	err = sendEmail(receiverEmailAccount, "Contact from "+message.From, message.Message)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func handleRequest(port string) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(mux.CORSMethodMiddleware(myRouter))
	myRouter.HandleFunc("/", homeHandler).Methods("GET")
	myRouter.HandleFunc("/", mailHandler).Methods("POST")

	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func sendEmail(sendTo string, subject string, msg string) error {

	var emailBody []byte
	auth := smtp.PlainAuth("", senderEmailAccount, smtpSecretPass, smtpServerURL)

	to := []string{sendTo}
	emailBody = []byte("To: " + sendTo + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		msg + "\r\n")

	err := smtp.SendMail(smtpServerURL+":587", auth, senderEmailAccount, to, emailBody)

	return err
}

func main() {
	port := os.Getenv("PORT")
	senderEmailAccount = os.Getenv("SENDER_EMAIL_ACCOUNT")
	receiverEmailAccount = os.Getenv("RECEIVER_EMAIL_ACCOUNT")
	smtpServerURL = os.Getenv("SMTP_SERVER_URL")
	smtpSecretPass = os.Getenv("SMTP_SECRET_PASS")

	validate = validator.New()

	handleRequest(port)
}
