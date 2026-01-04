package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tigusigalpa/yandexgpt-go"
	"github.com/tigusigalpa/yandexgpt-go/models"
)

func main() {
	oauthToken := os.Getenv("YANDEX_GPT_OAUTH_TOKEN")
	folderID := os.Getenv("YANDEX_GPT_FOLDER_ID")

	if oauthToken == "" || folderID == "" {
		log.Fatal("Please set YANDEX_GPT_OAUTH_TOKEN and YANDEX_GPT_FOLDER_ID environment variables")
	}

	client, err := yandexgpt.NewClient(oauthToken, folderID)
	if err != nil {
		log.Fatal(err)
	}

	messages := []yandexgpt.Message{
		{
			Role: "system",
			Text: "You are a helpful programming assistant",
		},
		{
			Role: "user",
			Text: "How to create a REST API in Go?",
		},
		{
			Role: "assistant",
			Text: "To create a REST API in Go, you can use the standard net/http package or frameworks like Gin or Echo. Here's a basic example using net/http...",
		},
		{
			Role: "user",
			Text: "And how to add validation?",
		},
	}

	fmt.Println("Sending dialogue to YandexGPT...")

	response, err := client.GenerateFromMessages(messages, models.YandexGPTLite, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nResponse:")
	fmt.Println(response.Result.Alternatives[0].Message.Text)
	fmt.Printf("\nTokens used: %d\n", response.Result.Usage.TotalTokens)
}
