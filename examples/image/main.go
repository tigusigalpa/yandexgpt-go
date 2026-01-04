package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/tigusigalpa/yandexgpt-go"
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

	fmt.Println("Generating image with YandexART...")
	fmt.Println("This may take a few minutes...")

	result, err := client.GenerateImage(
		"A beautiful sunset over mountains, digital art",
		nil,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	imageData, err := base64.StdEncoding.DecodeString(result.ImageBase64)
	if err != nil {
		log.Fatal(err)
	}

	outputFile := "output.jpg"
	err = os.WriteFile(outputFile, imageData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nImage saved to %s\n", outputFile)
	fmt.Printf("Operation ID: %s\n", result.OperationID)
}
