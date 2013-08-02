package main

import (
	"log"
	"net/http"
)

func TwilioHandler(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println("ParseForm: ", err)
		return
	}
	src := req.Form["From"]
	msg := req.Form["Body"]
	log.Println(src, msg)
}

func main() {
	http.HandleFunc("/twilio", TwilioHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
