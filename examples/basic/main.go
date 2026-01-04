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

	fmt.Println("Sending request to YandexGPT...")

	response, err := client.GenerateText(
		"Tell me about the benefits of Go programming language",
		models.YandexGPTLite,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nResponse:")
	fmt.Println(response.Result.Alternatives[0].Message.Text)
	fmt.Printf("\nTokens used: %d\n", response.Result.Usage.TotalTokens)
}
