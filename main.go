package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	aTest      = "testbdjfbdgb"
	REFERRER   = "https://www.google.com"
	USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"
)

type spamRequest struct {
	MethodName     string      `json:"method_name"`
	Message        string      `json:"message"`
	AuthKey        string      `json:"auth_key"`
	SenderEmail    string      `json:"sender_email"`
	SenderNickname string      `json:"sender_nickname"`
	SenderIp       string      `json:"sender_ip"`
	JsOn           int         `json:"js_on"`
	SubmitTime     int         `json:"submit_time"`
	SenderInfo     *senderInfo `json:"sender_info"`
}

type senderInfo struct {
	Referrer  string `json:"referrer"`
	UserAgent string `json:"user_agent"`
}

type spamResponse struct {
	Allow         int    `json:"allow"`
	Comment       string `json:"comment"`
	StopQueue     int    `json:"stop_queue"`
	Spam          int    `json:"spam"`
	Blacklisted   int    `json:"blacklisted"`
	AccountStatus int    `json:"account_status"`
}

type userRequest struct {
	UserName string `query:"userName"`
	Email    string `query:"email"`
	Message  string `query:"message"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(logger.New())

	app.Post("/api/spam-check", func(c *fiber.Ctx) error {
		var req userRequest
		fmt.Println("Inside api call")
		if err := c.BodyParser(&req); err != nil {
			log.Fatal("error while parsing a request body")
			log.Fatal(err)
			return err
		}
		res, err := checkMessageSpam(req)
		if err != nil {
			return err
		}

		var dynamicErr string
		if strings.Contains(res.Comment, "Trial period expired.") || strings.Contains(res.Comment, "Antispam disabled") {
			dynamicErr = "Something went wrong. Please contact to support."
		} else if res.Allow == 0 || res.Spam == 1 || res.AccountStatus == 0 {
			dynamicErr = ""
		}
		return c.Status(200).JSON(fiber.Map{
			"spamResponse": res,
			"apiErr":       dynamicErr,
		})
	})
	app.Listen(":8080")
}

func checkMessageSpam(requestBody userRequest) (*spamResponse, error) {
	userIP, errip := getIP()
	if errip != nil {
		log.Fatal(errip)
		log.Println(errip)
		fmt.Println(errip)
		return nil, errip
	}

	request := &spamRequest{
		MethodName:     "check_message",
		Message:        requestBody.Message,
		AuthKey:        aTest,
		SenderEmail:    requestBody.Email,
		SenderNickname: requestBody.UserName,
		SenderIp:       userIP,
		JsOn:           1,
		SubmitTime:     15,
		SenderInfo: &senderInfo{
			Referrer:  REFERRER,
			UserAgent: USER_AGENT,
		},
	}

	url := "https://moderate.cleantalk.org/api2.0"
	jsonReq, err := json.Marshal(request)
	if err != nil {
		log.Fatal("Issue while marshalling the json call")
		log.Fatal(err)
		fmt.Println(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatal("Issue while calling the api")
		log.Fatal(err)
		fmt.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var response spamResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return nil, err
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
