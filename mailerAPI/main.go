package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"github.com/gorilla/mux"
	"mailer"
)

type Message struct {
    From      string `json:"from"`
    Message   string `json:"message"`
}

var senderEmailAccount string;
var receiverEmailAccount string;
var smtpServerUrl string;
var smtpSecretPass string;


func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func mailHandler(w http.ResponseWriter, r *http.Request) {

    reqBody, _ := ioutil.ReadAll(r.Body) //read request body

	var message Message 
    json.Unmarshal(reqBody, &message); //convert to a message type
	sendEmail(receiverEmailAccount, "Contact from " + message.From , message.Message)
}

func handleRequest(port string) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homeHandler)
	myRouter.HandleFunc("/", mailHandler).Methods("POST")
	
	if port == "" { port = ":8080" }
	log.Fatal(http.ListenAndServe(port, myRouter))
}

func main() {
	port := os.Getenv("PORT")
	senderEmailAccount = os.Getenv("SENDER_EMAIL_ACCOUNT")
	receiverEmailAccount = os.Getenv("RECEIVER_EMAIL_ACCOUNT")
	smtpServerUrl = os.Getenv("SMTP_SERVER_URL")
	smtpSecretPass = os.Getenv("SMTP_SECRET_PASS")

	handleRequest(port)
}