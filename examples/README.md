# YandexGPT Go SDK Examples

This directory contains examples demonstrating how to use the YandexGPT Go SDK.

## 📋 Prerequisites

Before running the examples, you need to:

1. **Get OAuth token**: Follow the [configuration guide](../README.md#configuration)
2. **Get Folder ID**: Set up your Yandex Cloud folder
3. **Set environment variables**:
   ```bash
   export YANDEX_GPT_OAUTH_TOKEN=your_oauth_token_here
   export YANDEX_GPT_FOLDER_ID=your_folder_id_here
   ```

## 🚀 Running Examples

### Basic Usage

Simple text generation with YandexGPT:

```bash
cd examples/basic
go run main.go
```

This example demonstrates:
- Creating a client
- Sending a simple text generation request
- Handling the response

### Dialogue

Multi-turn conversation with context:

```bash
cd examples/dialogue
go run main.go
```

This example demonstrates:
- Creating a conversation with multiple messages
- Maintaining context across turns
- Using system messages to set behavior

### Image Generation

Generate images with YandexART:

```bash
cd examples/image
go run main.go
```

This example demonstrates:
- Generating images from text prompts
- Handling async operations
- Saving images to disk

**Note**: Image generation may take several minutes to complete.

### Custom Options

Using custom generation parameters:

```bash
cd examples/options
go run main.go
```

This example demonstrates:
- Setting custom temperature
- Limiting token count
- Configuring generation options

## 📚 Example Code Snippets

### Simple Text Generation

```go
client, err := yandexgpt.NewClient(oauthToken, folderID)
if err != nil {
    log.Fatal(err)
}

response, err := client.GenerateText(
    "Tell me about Go programming",
    models.YandexGPTLite,
    nil,
)
if err != nil {
    log.Fatal(err)
}

fmt.Println(response.Result.Alternatives[0].Message.Text)
```

### With Custom Options

```go
options := &yandexgpt.CompletionOptions{
    Temperature: 0.8,
    MaxTokens:   1000,
}

response, err := client.GenerateText(
    "Write a creative story",
    models.YandexGPT,
    options,
)
```

### Dialogue with Context

```go
messages := []yandexgpt.Message{
    {Role: "system", Text: "You are a helpful assistant"},
    {Role: "user", Text: "Hello!"},
    {Role: "assistant", Text: "Hi! How can I help you?"},
    {Role: "user", Text: "Tell me about Go"},
}

response, err := client.GenerateFromMessages(
    messages,
    models.YandexGPTLite,
    nil,
)
```

### Image Generation

```go
result, err := client.GenerateImage(
    "A beautiful landscape",
    nil,
    nil,
)
if err != nil {
    log.Fatal(err)
}

imageData, _ := base64.StdEncoding.DecodeString(result.ImageBase64)
os.WriteFile("output.jpg", imageData, 0644)
```

## 🔧 Troubleshooting

### Authentication Errors

If you get authentication errors:
- Check that your OAuth token is valid
- Ensure your Folder ID is correct
- Verify you have the necessary permissions

### API Errors

If you get API errors:
- Check your internet connection
- Verify your Yandex Cloud account is active
- Ensure you haven't exceeded rate limits

### Timeout Errors

If operations timeout:
- Image generation can take several minutes
- Increase timeout if needed
- Check Yandex Cloud status

## 📖 Additional Resources

- [Main README](../README.md)
- [API Documentation](../docs/api.md)
- [Configuration Guide](../docs/configuration.md)
- [Yandex Cloud Documentation](https://yandex.cloud/en/docs/foundation-models/)

## 💡 Tips

1. **Start with basic example**: Get familiar with the SDK
2. **Use appropriate models**: Choose based on your needs
3. **Handle errors properly**: Always check for errors
4. **Set reasonable limits**: Use appropriate token limits
5. **Monitor usage**: Keep track of API usage and costs

## 🤝 Contributing

Found an issue or want to add an example? See [CONTRIBUTING.md](../CONTRIBUTING.md)
