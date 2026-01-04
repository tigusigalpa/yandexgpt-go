// Package yandexgpt provides a Go SDK for working with YandexGPT API.
//
// This package offers a convenient interface for integrating with Yandex Cloud AI models,
// including YandexGPT for text generation and YandexART for image generation.
//
// # Features
//
//   - Easy integration with the YandexGPT API
//   - YandexART integration for image generation
//   - Automatic management of OAuth and IAM tokens
//   - Support for all available YandexGPT models
//   - Support for dialogues and single requests
//   - Automatic token renewal
//
// # Basic Usage
//
// Create a client and generate text:
//
//	client, err := yandexgpt.NewClient("oauth_token", "folder_id")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	response, err := client.GenerateText(
//	    "Tell me about Go programming",
//	    models.YandexGPTLite,
//	    nil,
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	fmt.Println(response.Result.Alternatives[0].Message.Text)
//
// # Working with Dialogues
//
// Create multi-turn conversations:
//
//	messages := []yandexgpt.Message{
//	    {Role: "system", Text: "You are a helpful assistant"},
//	    {Role: "user", Text: "Hello!"},
//	    {Role: "assistant", Text: "Hi! How can I help?"},
//	    {Role: "user", Text: "Tell me about Go"},
//	}
//
//	response, err := client.GenerateFromMessages(
//	    messages,
//	    models.YandexGPTLite,
//	    nil,
//	)
//
// # Image Generation
//
// Generate images with YandexART:
//
//	result, err := client.GenerateImage(
//	    "A beautiful landscape",
//	    nil,
//	    nil,
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	imageData, _ := base64.StdEncoding.DecodeString(result.ImageBase64)
//	os.WriteFile("output.jpg", imageData, 0644)
//
// # Custom Options
//
// Configure generation parameters:
//
//	options := &yandexgpt.CompletionOptions{
//	    Temperature: 0.8,
//	    MaxTokens:   1000,
//	}
//
//	response, err := client.GenerateText(
//	    "Write a creative story",
//	    models.YandexGPT,
//	    options,
//	)
//
// # Error Handling
//
// The package provides specific error types:
//
//	response, err := client.GenerateText("Hello", models.YandexGPTLite, nil)
//	if err != nil {
//	    var authErr *yandexgpt.AuthenticationError
//	    var apiErr *yandexgpt.APIError
//
//	    switch {
//	    case errors.As(err, &authErr):
//	        // Handle authentication error
//	    case errors.As(err, &apiErr):
//	        // Handle API error
//	    default:
//	        // Handle other errors
//	    }
//	}
//
// For more information, see the documentation at:
// https://github.com/tigusigalpa/yandexgpt-go
package yandexgpt
