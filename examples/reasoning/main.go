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

	fmt.Println("=== Example 1: Reasoning Mode with ENABLED_HIDDEN ===")
	fmt.Println()

	effortLow := "low"
	options1 := &yandexgpt.CompletionOptions{
		Temperature: 0.1,
		MaxTokens:   2000,
		Stream:      false,
		ReasoningOptions: &yandexgpt.ReasoningOptions{
			Mode:   "ENABLED_HIDDEN",
			Effort: &effortLow,
		},
	}

	response1, err := client.GenerateText(
		"Решите задачу: У Маши было 5 яблок. Она отдала 2 яблока Пете, а потом купила еще 3 яблока. Сколько яблок у нее теперь?",
		models.YandexGPT,
		options1,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response:")
	fmt.Println(response1.Result.Alternatives[0].Message.Text)
	fmt.Printf("\nTokens used: %d\n", response1.Result.Usage.TotalTokens)
	if response1.Result.Usage.ReasoningTokens > 0 {
		fmt.Printf("Reasoning tokens: %d\n", response1.Result.Usage.ReasoningTokens)
	}
	fmt.Printf("Model version: %s\n", response1.Result.ModelVersion)

	fmt.Println()
	fmt.Println("=== Example 2: Reasoning Mode with medium effort ===")
	fmt.Println()

	effortMedium := "medium"
	options2 := &yandexgpt.CompletionOptions{
		Temperature: 0.1,
		MaxTokens:   2000,
		Stream:      false,
		ReasoningOptions: &yandexgpt.ReasoningOptions{
			Mode:   "ENABLED_HIDDEN",
			Effort: &effortMedium,
		},
	}

	response2, err := client.GenerateText(
		"Если все розы - цветы, и некоторые цветы - красные, можно ли утверждать, что некоторые розы - красные?",
		models.YandexGPT,
		options2,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response:")
	fmt.Println(response2.Result.Alternatives[0].Message.Text)
	fmt.Printf("\nTokens used: %d\n", response2.Result.Usage.TotalTokens)
	if response2.Result.Usage.ReasoningTokens > 0 {
		fmt.Printf("Reasoning tokens: %d\n", response2.Result.Usage.ReasoningTokens)
	}
	fmt.Printf("Model version: %s\n", response2.Result.ModelVersion)

	fmt.Println()
	fmt.Println("=== Example 3: Disabled Reasoning Mode (default) ===")
	fmt.Println()

	options3 := &yandexgpt.CompletionOptions{
		Temperature: 0.1,
		MaxTokens:   1000,
		Stream:      false,
		ReasoningOptions: &yandexgpt.ReasoningOptions{
			Mode: "DISABLED",
		},
	}

	response3, err := client.GenerateText(
		"Что такое Go?",
		models.YandexGPT,
		options3,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response:")
	fmt.Println(response3.Result.Alternatives[0].Message.Text)
	fmt.Printf("\nTokens used: %d\n", response3.Result.Usage.TotalTokens)
	if response3.Result.Usage.ReasoningTokens > 0 {
		fmt.Printf("Reasoning tokens: %d\n", response3.Result.Usage.ReasoningTokens)
	}
	fmt.Printf("Model version: %s\n", response3.Result.ModelVersion)
}
