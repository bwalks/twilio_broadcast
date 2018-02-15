package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"io"
	"net/http"
	"os"
)

const baseUrl string = "https://api.twilio.com/2010-04-01"

var (
	accountSID string
	authToken string
	toNumber string
	fromNumber string
	messageChannel chan string
)

func init(){
	accountSID = os.Getenv("ACCOUNT_SID")
	authToken = os.Getenv("AUTH_TOKEN")
	toNumber = os.Getenv("TO_PHONE_NUMBER")
	fromNumber = os.Getenv("FROM_PHONE_NUMBER")
	messageChannel = make(chan string, 5)
}

func main(){
	go worker(messageChannel)
	http.HandleFunc("/twilio/broadcast", broadcast)
	http.ListenAndServe(":8000", nil)
}

func worker(c chan string){
	for {
		val := <- c
		fmt.Println(val)
		url := fmt.Sprintf("%s/Accounts/%s/Messages.json", baseUrl, accountSID)
		httpclient.
			WithHeader("Authorization", "Basic " + basicAuth(accountSID, authToken)).
			WithHeader("Content-Type", "application/json").
			Post(url, map[string]string {
				"To": toNumber,
				"From": fromNumber,
				"Body": val,
			})
	}
}

func toString(r io.Reader) string{
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return buf.String()
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
  }

func broadcast(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("Body")	
	if body == ""{
		return
	}
	messageChannel <- body
}