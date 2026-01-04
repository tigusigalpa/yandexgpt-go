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

	options := &yandexgpt.CompletionOptions{
		Temperature: 0.8,
		MaxTokens:   1000,
		Stream:      false,
	}

	fmt.Println("Sending request with custom options...")
	fmt.Printf("Temperature: %.1f, MaxTokens: %d\n\n", options.Temperature, options.MaxTokens)

	response, err := client.GenerateText(
		"Write a short poem about programming in Go",
		models.YandexGPT,
		options,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response:")
	fmt.Println(response.Result.Alternatives[0].Message.Text)
	fmt.Printf("\nTokens used: %d\n", response.Result.Usage.TotalTokens)
	fmt.Printf("Model version: %s\n", response.Result.ModelVersion)
}
