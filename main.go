package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type spamRequest struct {
	MethodName     string `json:"method_name"`
	Message        string `json:"message"`
	AuthKey        string `json:"auth_key"`
	SenderEmail    string `json:"sender_email"`
	SenderNickname string `json:"sender_nickname"`
	SenderIp       string `json:"sender_ip"`
	JsOn           int    `json:"js_on"`
	SubmitTime     int    `json:"submit_time"`
}

type spamResponse struct {
	Data map[string]interface{} `json:"data"`
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/api/spam-check", func(c *fiber.Ctx) error {
		res, err := checkMessageSpam()
		if err != nil {
			return err
		}
		return c.Status(200).JSON(fiber.Map{
			"spamResponse": res,
		})
	})
	response, err := checkMessageSpam()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

func checkMessageSpam() (*spamResponse, error) {
	userIP, errip := getIP()
	if errip != nil {
		log.Fatal(errip)
		log.Println(errip)
		fmt.Println(errip)
		return nil, errip
	}

	request := &spamRequest{
		MethodName:     "check_message",
		Message:        "hydg agdhs sgfn sgnf",
		AuthKey:        "hdfhdhgdhvgdhgb",
		SenderEmail:    "abc@test.com",
		SenderNickname: "Abc Test",
		SenderIp:       userIP,
		JsOn:           1,
		SubmitTime:     15,
	}

	url := "https://moderate.cleantalk.org/api2.0"
	jsonReq, err := json.Marshal(request)
	if err != nil {
		log.Fatal("Issue while marshalling the json call")
		log.Fatal(err)
		fmt.Println(err)
		return nil, err
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatal("Issue while calling the api")
		log.Fatal(err)
		fmt.Println(err)
		return nil, err
	}

	var response spamResponse
	resErr := json.NewDecoder(res.Body).Decode(&response)
	if resErr != nil {
		log.Fatal("Error while decoding the response")
		log.Fatal(resErr)
		fmt.Println(resErr)
		return nil, resErr
	}

	return &response, nil
}

func getIP() (string, error) {
	res, err := http.Get("https://ipify.org/?format=text")
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
