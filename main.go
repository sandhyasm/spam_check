package main

import "fmt"

type spamRequest struct {
	MethodName  string `json:"method_name"`
	AuthKey     string `json:"auth_key"`
	SenderEmail string `json:"sender_email"`
	SenderIp    string `json:"sender_ip"`
	JsOn        int    `json:"js_on"`
	SubmitTime  int    `json:"submit_time"`
}

type spamResponse struct {
	Data []
}

func main() {
	fmt.Println("Init checks")
}
