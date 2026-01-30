# YandexGPT Go SDK

![YandexGPT Golang](https://github.com/user-attachments/assets/35073f18-14ce-486f-937b-4c70e9af9e6c)

[![Go Version](https://img.shields.io/github/go-mod/go-version/tigusigalpa/yandexgpt-go)](https://github.com/tigusigalpa/yandexgpt-go)
[![License](https://img.shields.io/github/license/tigusigalpa/yandexgpt-go)](https://github.com/tigusigalpa/yandexgpt-go/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/tigusigalpa/yandexgpt-go)](https://goreportcard.com/report/github.com/tigusigalpa/yandexgpt-go)

**🌐 Language:** English | [Русский](README.md)

**📦 Package:** [pkg.go.dev/github.com/tigusigalpa/yandexgpt-go](https://pkg.go.dev/github.com/tigusigalpa/yandexgpt-go)

A full-featured Go/Golang SDK for working with YandexGPT API. The package provides a convenient interface for
integrating with Yandex Cloud AI models, including YandexART support for image generation.

**YandexGPT Go SDK** is a powerful library for Go developers that simplifies integration with artificial intelligence from Yandex Cloud. The library enables you to create intelligent applications, chatbots, automation systems, and AI assistants using advanced language models like YandexGPT, YandexGPT Lite, and AliceAI LLM.

### Why Choose YandexGPT Go SDK?

- **🚀 Rapid Development**: Ready-to-use methods for text generation, dialogues, and image creation
- **🔒 Security**: Automatic OAuth and IAM token management with renewal support
- **⚡ Performance**: Optimized API interaction with concurrent request support
- **🎨 Versatility**: Support for YandexGPT text models and image generation via YandexART
- **📖 Documentation**: Detailed code examples and quick start guides
- **🧪 Reliability**: Full test coverage and active support

### Main Use Cases

- **Chatbots and Virtual Assistants**: Create intelligent dialogue systems for customer support
- **Content Generation**: Automatically create texts, articles, product descriptions
- **Text Analysis and Processing**: Summarization, classification, information extraction
- **Image Generation**: Create visual content based on text descriptions
- **Business Process Automation**: Intelligent document and data processing
- **Educational Platforms**: AI assistants for learning and testing

> **Note:** This package uses [yandex-cloud-client-go](https://github.com/tigusigalpa/yandex-cloud-client-go) for
> managing Yandex Cloud infrastructure (organizations, clouds, folders, authorization).

## 🚀 Features

- 🔌 Easy integration with the YandexGPT API
- 🔨 **YandexART integration**
- 🔐 Automatic management of OAuth and IAM tokens
- 🎯 Support for all available YandexGPT models
- 📝 Support for dialogues and single requests
- ⚡ Automatic token renewal
- 🧪 Test coverage
- 📚 Detailed documentation
- 🚀 High performance and concurrency support

---

## 📦 Installation

Install the package using `go get`:

```bash
go get github.com/tigusigalpa/yandexgpt-go
```

---

## ⚙️ Configuration

### 1. Getting an OAuth token

📚 **Documentation:** [OAuth-token](https://yandex.cloud/en/docs/iam/concepts/authorization/oauth-token)

Follow the link
to [get an OAuth token](https://oauth.yandex.com/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb):

```
https://oauth.yandex.com/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb
```

### 2. Environment configuration

Set environment variables or pass them directly to the client:

```bash
export YANDEX_GPT_OAUTH_TOKEN=your_oauth_token_here
export YANDEX_GPT_FOLDER_ID=your_folder_id_here
```

### 3. Yandex Cloud preparation

For detailed setup instructions, see the [Configuration Guide](docs/configuration.md).

---

## 💡 Usage

### Basic usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/tigusigalpa/yandexgpt-go"
    "github.com/tigusigalpa/yandexgpt-go/models"
)

func main() {
    // Create a client
    client, err := yandexgpt.NewClient("your_oauth_token", "your_folder_id")
    if err != nil {
        log.Fatal(err)
    }
    
    // Simple request
    response, err := client.GenerateText(
        "Tell me about the benefits of Go programming language",
        models.YandexGPTLite,
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(response.Result.Alternatives[0].Message.Text)
}
```

### Working with dialogues

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/tigusigalpa/yandexgpt-go"
    "github.com/tigusigalpa/yandexgpt-go/models"
)

func main() {
    client, err := yandexgpt.NewClient("your_oauth_token", "your_folder_id")
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
            Text: "To create a REST API in Go, you can use the standard net/http package...",
        },
        {
            Role: "user",
            Text: "And how to add validation?",
        },
    }
    
    response, err := client.GenerateFromMessages(messages, models.YandexGPTLite, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(response.Result.Alternatives[0].Message.Text)
}
```

### Image generation with YandexART

```go
package main

import (
    "encoding/base64"
    "fmt"
    "log"
    "os"
    
    "github.com/tigusigalpa/yandexgpt-go"
)

func main() {
    client, err := yandexgpt.NewClient("your_oauth_token", "your_folder_id")
    if err != nil {
        log.Fatal(err)
    }
    
    // Generate image
    result, err := client.GenerateImage(
        "A beautiful sunset over mountains",
        nil,
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    
    // Decode base64 image
    imageData, err := base64.StdEncoding.DecodeString(result.ImageBase64)
    if err != nil {
        log.Fatal(err)
    }
    
    // Save to file
    err = os.WriteFile("output.jpg", imageData, 0644)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Image saved to output.jpg")
}
```

### Custom options

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/tigusigalpa/yandexgpt-go"
    "github.com/tigusigalpa/yandexgpt-go/models"
)

func main() {
    client, err := yandexgpt.NewClient("your_oauth_token", "your_folder_id")
    if err != nil {
        log.Fatal(err)
    }
    
    options := &yandexgpt.CompletionOptions{
        Temperature: 0.8,
        MaxTokens:   1000,
    }
    
    response, err := client.GenerateText(
        "Write a poem about programming",
        models.YandexGPT,
        options,
    )
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(response.Result.Alternatives[0].Message.Text)
}
```

### Reasoning Mode

The reasoning mode enables models to perform chain-of-thought reasoning for complex tasks:

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/tigusigalpa/yandexgpt-go"
    "github.com/tigusigalpa/yandexgpt-go/models"
)

func main() {
    client, err := yandexgpt.NewClient("your_oauth_token", "your_folder_id")
    if err != nil {
        log.Fatal(err)
    }
    
    effortMedium := "medium"
    options := &yandexgpt.CompletionOptions{
        Temperature: 0.1,
        MaxTokens:   2000,
        ReasoningOptions: &yandexgpt.ReasoningOptions{
            Mode:   "ENABLED_HIDDEN",
            Effort: &effortMedium,
        },
    }
    
    response, err := client.GenerateText(
        "Solve this logic problem: If all roses are flowers, and some flowers are red, can we conclude that some roses are red?",
        models.YandexGPT,
        options,
    )
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(response.Result.Alternatives[0].Message.Text)
    
    // Check reasoning tokens usage
    if response.Result.Usage.ReasoningTokens > 0 {
        fmt.Printf("Reasoning tokens used: %d\n", response.Result.Usage.ReasoningTokens)
    }
}
```

**Reasoning Mode Options:**

- `Mode`: `"DISABLED"` (default), `"ENABLED_HIDDEN"` (enables reasoning without showing the chain)
- `Effort`: `"low"`, `"medium"`, `"high"` (controls reasoning depth, optional)

📚 **Documentation:
** [Reasoning Mode in YandexGPT](https://yandex.cloud/ru/docs/ai-studio/concepts/generation/chain-of-thought)

---

## 🤖 Available models

| Model            | Description                                  | Constant               | Context |
|------------------|----------------------------------------------|------------------------|---------|
| `yandexgpt-lite` | Fast and economical model                    | `models.YandexGPTLite` | 32K     |
| `yandexgpt`      | Standard model                               | `models.YandexGPT`     | 32K     |
| `aliceai-llm`    | Alice AI LLM - advanced conversational model | `models.AliceAI`       | 32K     |

📚 **Complete list of available models:
** [Generation models in Yandex AI Studio](https://yandex.cloud/en/docs/ai-studio/concepts/generation/models)

---

## 🔧 Generation parameters

```go
type CompletionOptions struct {
    Stream           bool              // Streaming (not yet supported)
    Temperature      float64           // Creativity (0.0 - 1.0)
    MaxTokens        int               // Maximum number of tokens
    ReasoningOptions *ReasoningOptions // Reasoning mode settings (optional)
}

type ReasoningOptions struct {
    Mode   string  // "DISABLED", "ENABLED_HIDDEN"
    Effort *string // "low", "medium", "high" (optional)
}
```

---

## ⚠️ Error handling

```go
package main

import (
    "errors"
    "fmt"
    "log"
    
    "github.com/tigusigalpa/yandexgpt-go"
    "github.com/tigusigalpa/yandexgpt-go/models"
)

func main() {
    client, err := yandexgpt.NewClient("your_oauth_token", "your_folder_id")
    if err != nil {
        log.Fatal(err)
    }
    
    response, err := client.GenerateText("Hello!", models.YandexGPTLite, nil)
    if err != nil {
        var authErr *yandexgpt.AuthenticationError
        var apiErr *yandexgpt.APIError
        
        switch {
        case errors.As(err, &authErr):
            fmt.Printf("Authentication error: %v\n", err)
        case errors.As(err, &apiErr):
            fmt.Printf("API error: %v\n", err)
        default:
            fmt.Printf("Unknown error: %v\n", err)
        }
        return
    }
    
    fmt.Println(response.Result.Alternatives[0].Message.Text)
}
```

---

## 📚 Examples

See the [examples](examples/) directory for more usage examples:

- [Basic usage](examples/basic/main.go)
- [Dialogue](examples/dialogue/main.go)
- [Image generation](examples/image/main.go)
- [Custom options](examples/options/main.go)
- [Reasoning mode](examples/reasoning/main.go)

---

## 🔍 YandexGPT Models Comparison

### Choosing the Right Model for Your Task

| Feature | YandexGPT Lite | YandexGPT | AliceAI LLM |
|---------|----------------|-----------|-------------|
| **Speed** | ⚡⚡⚡ Very Fast | ⚡⚡ Fast | ⚡⚡ Fast |
| **Quality** | ⭐⭐⭐ Good | ⭐⭐⭐⭐⭐ Excellent | ⭐⭐⭐⭐ Very Good |
| **Cost** | 💰 Economical | 💰💰 Standard | 💰💰 Standard |
| **Context** | 32K tokens | 32K tokens | 32K tokens |
| **Use Case** | Simple tasks, chatbots | Complex tasks, analysis | Dialogues, assistants |
| **Creativity** | Medium | High | High |

### Model Selection Recommendations

**YandexGPT Lite** - ideal for:
- Quick chatbot responses
- Simple text classification
- Short text generation
- High-load applications
- Prototyping and testing

**YandexGPT** - recommended for:
- Complex text analysis
- Long article generation
- Creative tasks (poems, stories)
- Technical documentation
- Logical problem solving

**AliceAI LLM** - optimal for:
- Natural dialogues
- Virtual assistants
- Personalized recommendations
- Context-dependent responses
- Emotionally colored communication

---

## 🎯 Advanced Features

### Dialogue Context Management

For creating quality dialogue systems, proper context management is essential:

```go
type DialogManager struct {
    client   *yandexgpt.Client
    messages []yandexgpt.Message
    maxHistory int
}

func NewDialogManager(client *yandexgpt.Client, systemPrompt string) *DialogManager {
    return &DialogManager{
        client: client,
        messages: []yandexgpt.Message{
            {Role: "system", Text: systemPrompt},
        },
        maxHistory: 10, // Keep last 10 messages
    }
}

func (dm *DialogManager) AddUserMessage(text string) error {
    dm.messages = append(dm.messages, yandexgpt.Message{
        Role: "user",
        Text: text,
    })
    
    // Limit history to save tokens
    if len(dm.messages) > dm.maxHistory {
        // Keep system message and last N messages
        dm.messages = append(
            dm.messages[:1],
            dm.messages[len(dm.messages)-dm.maxHistory+1:]...,
        )
    }
    
    response, err := dm.client.GenerateFromMessages(
        dm.messages,
        models.YandexGPT,
        nil,
    )
    if err != nil {
        return err
    }
    
    assistantMessage := response.Result.Alternatives[0].Message.Text
    dm.messages = append(dm.messages, yandexgpt.Message{
        Role: "assistant",
        Text: assistantMessage,
    })
    
    return nil
}
```

### Batch Request Processing

For processing multiple requests, use concurrency:

```go
func ProcessBatch(client *yandexgpt.Client, prompts []string) ([]string, error) {
    results := make([]string, len(prompts))
    errors := make([]error, len(prompts))
    
    var wg sync.WaitGroup
    semaphore := make(chan struct{}, 5) // Limit: 5 concurrent requests
    
    for i, prompt := range prompts {
        wg.Add(1)
        go func(index int, text string) {
            defer wg.Done()
            semaphore <- struct{}{}
            defer func() { <-semaphore }()
            
            response, err := client.GenerateText(text, models.YandexGPTLite, nil)
            if err != nil {
                errors[index] = err
                return
            }
            results[index] = response.Result.Alternatives[0].Message.Text
        }(i, prompt)
    }
    
    wg.Wait()
    
    // Check for errors
    for _, err := range errors {
        if err != nil {
            return results, err
        }
    }
    
    return results, nil
}
```

### Working with Large Texts

For processing texts exceeding context limits:

```go
func SummarizeLongText(client *yandexgpt.Client, longText string) (string, error) {
    // Split text into chunks of ~3000 characters
    chunkSize := 3000
    var chunks []string
    
    for i := 0; i < len(longText); i += chunkSize {
        end := i + chunkSize
        if end > len(longText) {
            end = len(longText)
        }
        chunks = append(chunks, longText[i:end])
    }
    
    // Summarize each chunk
    var summaries []string
    for _, chunk := range chunks {
        response, err := client.GenerateText(
            fmt.Sprintf("Briefly summarize the following text:\n\n%s", chunk),
            models.YandexGPT,
            &yandexgpt.CompletionOptions{Temperature: 0.3},
        )
        if err != nil {
            return "", err
        }
        summaries = append(summaries, response.Result.Alternatives[0].Message.Text)
    }
    
    // Final summarization
    finalPrompt := fmt.Sprintf(
        "Combine the following summaries into one coherent text:\n\n%s",
        strings.Join(summaries, "\n\n"),
    )
    
    response, err := client.GenerateText(finalPrompt, models.YandexGPT, nil)
    if err != nil {
        return "", err
    }
    
    return response.Result.Alternatives[0].Message.Text, nil
}
```

---

## 💡 Best Practices

### Prompt Optimization

1. **Be Specific**: Clearly formulate the task and expected response format
2. **Use Examples**: Few-shot learning improves result quality
3. **Structure Requests**: Separate instructions from data
4. **Control Length**: Optimize token usage

```go
// ❌ Bad prompt
prompt := "Tell me about Go"

// ✅ Good prompt
prompt := `Write a brief description of the Go programming language (up to 100 words).
Include the following aspects:
1. Main purpose
2. Key features
3. Popular application areas

Format: structured text with subheadings.`
```

### Temperature Management

- **0.0-0.3**: Deterministic, precise answers (analysis, classification)
- **0.4-0.7**: Balanced answers (general tasks)
- **0.8-1.0**: Creative, diverse answers (content generation)

### Error Handling and Retries

```go
func GenerateWithRetry(client *yandexgpt.Client, prompt string, maxRetries int) (string, error) {
    var lastErr error
    
    for attempt := 0; attempt < maxRetries; attempt++ {
        response, err := client.GenerateText(prompt, models.YandexGPT, nil)
        if err == nil {
            return response.Result.Alternatives[0].Message.Text, nil
        }
        
        lastErr = err
        
        // Exponential backoff
        backoff := time.Duration(math.Pow(2, float64(attempt))) * time.Second
        time.Sleep(backoff)
    }
    
    return "", fmt.Errorf("failed after %d attempts: %w", maxRetries, lastErr)
}
```

### Result Caching

```go
type CachedClient struct {
    client *yandexgpt.Client
    cache  map[string]string
    mu     sync.RWMutex
}

func (cc *CachedClient) GenerateText(prompt string) (string, error) {
    // Check cache
    cc.mu.RLock()
    if cached, ok := cc.cache[prompt]; ok {
        cc.mu.RUnlock()
        return cached, nil
    }
    cc.mu.RUnlock()
    
    // Generate new response
    response, err := cc.client.GenerateText(prompt, models.YandexGPT, nil)
    if err != nil {
        return "", err
    }
    
    result := response.Result.Alternatives[0].Message.Text
    
    // Save to cache
    cc.mu.Lock()
    cc.cache[prompt] = result
    cc.mu.Unlock()
    
    return result, nil
}
```

---

## 🔧 Troubleshooting

### Common Issues and Solutions

#### Authentication Error

**Problem**: `Authentication error: invalid OAuth token`

**Solution**:
1. Check OAuth token expiration (valid for 1 year)
2. Get a new token via [link](https://oauth.yandex.com/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb)
3. Ensure token is correctly passed to client

```go
// Token validation
if token == "" {
    log.Fatal("OAuth token is required")
}
```

#### Folder Access Error

**Problem**: `Permission denied for folder`

**Solution**:
1. Ensure you have folder permissions in Yandex Cloud
2. Verify folder_id correctness
3. Assign `ai.languageModels.user` role to your account

```bash
yc resource-manager folder add-access-binding <folder-id> \
  --role ai.languageModels.user \
  --subject userAccount:<user-id>
```

#### Token Limit Exceeded

**Problem**: `Maximum token limit exceeded`

**Solution**:
1. Reduce prompt or dialogue history length
2. Use `MaxTokens` parameter to limit response
3. Split large text into parts

```go
options := &yandexgpt.CompletionOptions{
    MaxTokens: 1000, // Limit response length
}
```

#### Slow Responses

**Problem**: Long API response time

**Solution**:
1. Use YandexGPT Lite for simple tasks
2. Optimize prompt length
3. Consider caching frequent requests
4. Use concurrency for batch processing

#### Low Quality Responses

**Problem**: Responses don't meet expectations

**Solution**:
1. Improve prompt formulation (add examples, context)
2. Use YandexGPT instead of Lite for complex tasks
3. Adjust temperature (lower for precision, higher for creativity)
4. Add system message with instructions

---

## ❓ Frequently Asked Questions (FAQ)

### General Questions

**Q: Do I need a paid Yandex Cloud account?**

A: Yes, using YandexGPT API requires a Yandex Cloud account with payment information. A trial period with free grants is available.

**Q: What are the API usage limits?**

A: Limits depend on your plan. By default:
- 200 requests per minute
- 32K token context
- Monthly token limits according to plan

**Q: Is streaming supported?**

A: In the current SDK version, streaming is under development. Check [CHANGELOG.md](CHANGELOG.md) for updates.

**Q: Can I use the SDK in production?**

A: Yes, the SDK is production-ready. Recommended:
- Implement error handling and retry logic
- Set up monitoring and logging
- Use caching for optimization
- Monitor token usage

### Technical Questions

**Q: How to store OAuth token securely?**

A: Use environment variables or secret stores:

```go
// From environment variables
token := os.Getenv("YANDEX_GPT_OAUTH_TOKEN")

// From config file (don't commit to git!)
// Or use vault systems (HashiCorp Vault, AWS Secrets Manager)
```

**Q: How to track token usage?**

A: Check the `Usage` field in response:

```go
response, err := client.GenerateText(prompt, models.YandexGPT, nil)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Input tokens: %d\n", response.Result.Usage.InputTextTokens)
fmt.Printf("Output tokens: %d\n", response.Result.Usage.CompletionTokens)
fmt.Printf("Total tokens: %d\n", response.Result.Usage.TotalTokens)
```

**Q: Can I use multiple models simultaneously?**

A: Yes, just specify the needed model in each request:

```go
// Quick response
quickResponse, _ := client.GenerateText(prompt, models.YandexGPTLite, nil)

// Detailed analysis
detailedResponse, _ := client.GenerateText(prompt, models.YandexGPT, nil)
```

**Q: How to handle rate limiting?**

A: Implement exponential backoff and check response headers:

```go
func handleRateLimit(err error) bool {
    var apiErr *yandexgpt.APIError
    if errors.As(err, &apiErr) {
        if apiErr.StatusCode == 429 { // Too Many Requests
            return true
        }
    }
    return false
}
```

---

## ⚡ Performance and Optimization

### Optimization Tips

1. **Model Selection**: Use YandexGPT Lite for simple tasks - it's 2-3x faster
2. **Context Management**: Limit dialogue history to last 10-15 messages
3. **Parallelism**: Process independent requests concurrently
4. **Caching**: Cache frequent or identical requests
5. **Connection Pools**: Reuse HTTP clients

### Performance Monitoring

```go
type MetricsCollector struct {
    totalRequests   int64
    totalTokens     int64
    totalDuration   time.Duration
    mu              sync.Mutex
}

func (mc *MetricsCollector) TrackRequest(duration time.Duration, tokens int) {
    mc.mu.Lock()
    defer mc.mu.Unlock()
    
    mc.totalRequests++
    mc.totalTokens += int64(tokens)
    mc.totalDuration += duration
}

func (mc *MetricsCollector) GetStats() (avgDuration time.Duration, avgTokens float64) {
    mc.mu.Lock()
    defer mc.mu.Unlock()
    
    if mc.totalRequests == 0 {
        return 0, 0
    }
    
    avgDuration = mc.totalDuration / time.Duration(mc.totalRequests)
    avgTokens = float64(mc.totalTokens) / float64(mc.totalRequests)
    return
}
```

### Cost Estimation

Calculate usage costs:

```go
func EstimateCost(inputTokens, outputTokens int, model string) float64 {
    // Example prices (check current prices on Yandex Cloud website)
    var pricePerMillion float64
    
    switch model {
    case "yandexgpt-lite":
        pricePerMillion = 0.4 // $0.4 per 1M tokens
    case "yandexgpt":
        pricePerMillion = 1.2 // $1.2 per 1M tokens
    default:
        pricePerMillion = 1.2
    }
    
    totalTokens := float64(inputTokens + outputTokens)
    return (totalTokens / 1_000_000) * pricePerMillion
}
```

---

## 🔐 Security

### Security Recommendations

1. **Never store tokens in code**: Use environment variables or secret stores
2. **Token Rotation**: Regularly update OAuth tokens
3. **Input Validation**: Verify user input before sending to API
4. **Access Limitation**: Use minimal necessary permissions in Yandex Cloud
5. **Logging**: Don't log sensitive data and tokens

### Secure Configuration Example

```go
type SecureConfig struct {
    OAuthToken string
    FolderID   string
}

func LoadSecureConfig() (*SecureConfig, error) {
    token := os.Getenv("YANDEX_GPT_OAUTH_TOKEN")
    folderID := os.Getenv("YANDEX_GPT_FOLDER_ID")
    
    if token == "" || folderID == "" {
        return nil, errors.New("missing required environment variables")
    }
    
    // Format validation
    if len(token) < 20 {
        return nil, errors.New("invalid token format")
    }
    
    return &SecureConfig{
        OAuthToken: token,
        FolderID:   folderID,
    }, nil
}
```

### Content Filtering

```go
func SanitizeInput(input string) string {
    // Remove potentially dangerous characters
    input = strings.TrimSpace(input)
    
    // Length limitation
    maxLength := 10000
    if len(input) > maxLength {
        input = input[:maxLength]
    }
    
    return input
}
```

---

## 🧪 Testing

Run tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

---

## 📖 Documentation

For detailed documentation, see:

- [Configuration Guide](docs/configuration.md)
- [API Reference](docs/api.md)
- [Examples](examples/)

---

## 🌟 Integration with Popular Frameworks

### Gin Web Framework

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/tigusigalpa/yandexgpt-go"
    "github.com/tigusigalpa/yandexgpt-go/models"
)

func main() {
    client, _ := yandexgpt.NewClient(
        os.Getenv("YANDEX_GPT_OAUTH_TOKEN"),
        os.Getenv("YANDEX_GPT_FOLDER_ID"),
    )
    
    r := gin.Default()
    
    r.POST("/api/chat", func(c *gin.Context) {
        var req struct {
            Message string `json:"message"`
        }
        
        if err := c.BindJSON(&req); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        
        response, err := client.GenerateText(req.Message, models.YandexGPT, nil)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        
        c.JSON(200, gin.H{
            "response": response.Result.Alternatives[0].Message.Text,
        })
    })
    
    r.Run(":8080")
}
```

### Echo Framework

```go
package main

import (
    "github.com/labstack/echo/v4"
    "github.com/tigusigalpa/yandexgpt-go"
)

func main() {
    client, _ := yandexgpt.NewClient(
        os.Getenv("YANDEX_GPT_OAUTH_TOKEN"),
        os.Getenv("YANDEX_GPT_FOLDER_ID"),
    )
    
    e := echo.New()
    
    e.POST("/generate", func(c echo.Context) error {
        prompt := c.FormValue("prompt")
        
        response, err := client.GenerateText(prompt, models.YandexGPTLite, nil)
        if err != nil {
            return c.JSON(500, map[string]string{"error": err.Error()})
        }
        
        return c.JSON(200, map[string]string{
            "text": response.Result.Alternatives[0].Message.Text,
        })
    })
    
    e.Start(":8080")
}
```

### gRPC Service

```go
type AIService struct {
    client *yandexgpt.Client
}

func (s *AIService) Generate(ctx context.Context, req *pb.GenerateRequest) (*pb.GenerateResponse, error) {
    response, err := s.client.GenerateText(
        req.Prompt,
        models.YandexGPT,
        &yandexgpt.CompletionOptions{
            Temperature: float64(req.Temperature),
            MaxTokens:   int(req.MaxTokens),
        },
    )
    if err != nil {
        return nil, err
    }
    
    return &pb.GenerateResponse{
        Text: response.Result.Alternatives[0].Message.Text,
        TokensUsed: int32(response.Result.Usage.TotalTokens),
    }, nil
}
```

---

## 📊 Monitoring and Logging

### Prometheus Integration

```go
import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    requestsTotal = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "yandexgpt_requests_total",
            Help: "Total number of YandexGPT requests",
        },
        []string{"model", "status"},
    )
    
    requestDuration = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "yandexgpt_request_duration_seconds",
            Help: "YandexGPT request duration",
        },
        []string{"model"},
    )
    
    tokensUsed = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "yandexgpt_tokens_used_total",
            Help: "Total tokens used",
        },
        []string{"model", "type"},
    )
)

func GenerateWithMetrics(client *yandexgpt.Client, prompt string, model string) (string, error) {
    start := time.Now()
    
    response, err := client.GenerateText(prompt, model, nil)
    
    duration := time.Since(start).Seconds()
    requestDuration.WithLabelValues(model).Observe(duration)
    
    if err != nil {
        requestsTotal.WithLabelValues(model, "error").Inc()
        return "", err
    }
    
    requestsTotal.WithLabelValues(model, "success").Inc()
    tokensUsed.WithLabelValues(model, "input").Add(float64(response.Result.Usage.InputTextTokens))
    tokensUsed.WithLabelValues(model, "output").Add(float64(response.Result.Usage.CompletionTokens))
    
    return response.Result.Alternatives[0].Message.Text, nil
}
```

### Structured Logging

```go
import "go.uber.org/zap"

type LoggingClient struct {
    client *yandexgpt.Client
    logger *zap.Logger
}

func (lc *LoggingClient) GenerateText(prompt string, model string) (string, error) {
    lc.logger.Info("generating text",
        zap.String("model", model),
        zap.Int("prompt_length", len(prompt)),
    )
    
    start := time.Now()
    response, err := lc.client.GenerateText(prompt, model, nil)
    duration := time.Since(start)
    
    if err != nil {
        lc.logger.Error("generation failed",
            zap.Error(err),
            zap.Duration("duration", duration),
        )
        return "", err
    }
    
    lc.logger.Info("generation completed",
        zap.Duration("duration", duration),
        zap.Int("tokens_used", response.Result.Usage.TotalTokens),
        zap.Int("response_length", len(response.Result.Alternatives[0].Message.Text)),
    )
    
    return response.Result.Alternatives[0].Message.Text, nil
}
```

---

## 🚀 Roadmap and Future Features

### Planned Features

- ✅ YandexGPT and YandexGPT Lite support
- ✅ YandexART integration
- ✅ Reasoning mode (Chain of Thought)
- ✅ Automatic token management
- 🔄 Response streaming
- 🔄 Function calling support
- 📋 Multimodal support (images in prompts)
- 📋 Asynchronous operations
- 📋 Extended embeddings support
- 📋 Vector database integration

### How to Contribute

We welcome community contributions! You can help by:

- 🐛 Reporting bugs via [Issues](https://github.com/tigusigalpa/yandexgpt-go/issues)
- 💡 Suggesting new features
- 📝 Improving documentation
- 🔧 Submitting Pull Requests
- ⭐ Starring the project on GitHub

---

## 🤝 Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details.

---

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🔗 Useful Links

### Official Yandex Cloud Documentation

- 📖 [YandexGPT Quickstart](https://yandex.cloud/en/docs/foundation-models/quickstart/yandexgpt)
- 🔑 [API Authentication](https://yandex.cloud/en/docs/iam/concepts/authorization/iam-token)
- 🏗️ [Resource Management](https://yandex.cloud/en/docs/resource-manager/)
- 🤖 [API Foundation Models](https://yandex.cloud/en/docs/foundation-models/concepts/api)
- 💰 [YandexGPT Pricing](https://yandex.cloud/en/docs/foundation-models/pricing)
- 🎨 [YandexART Documentation](https://yandex.cloud/en/docs/foundation-models/concepts/yandexart)
- 🧠 [Reasoning Mode](https://yandex.cloud/en/docs/ai-studio/concepts/generation/chain-of-thought)
- 📊 [Limits and Quotas](https://yandex.cloud/en/docs/foundation-models/concepts/limits)

### Community and Support

- 💬 [Yandex Cloud Telegram Community](https://t.me/yandexcloud)
- 🎓 [Training Materials](https://yandex.cloud/en/training)
- 📺 [Yandex Cloud YouTube Channel](https://www.youtube.com/@YandexCloudRu)
- 📰 [Yandex Cloud Blog](https://yandex.cloud/en/blog)

### Related Projects

- 🔧 [yandex-cloud-client-go](https://github.com/tigusigalpa/yandex-cloud-client-go) - Yandex Cloud management client
- 📦 [Official Yandex Cloud Go SDK](https://github.com/yandex-cloud/go-sdk)
- 🐍 [Python SDK for YandexGPT](https://github.com/yandex-cloud/python-sdk)

### Examples and Tutorials

- 🎯 [Building a Chatbot with YandexGPT](examples/chatbot/)
- 🖼️ [Image Generator with YandexART](examples/image-generator/)
- 📝 [Content Generation System](examples/content-generator/)
- 🔍 [Text Analyzer](examples/text-analyzer/)

---

## 📈 Project Statistics

![GitHub stars](https://img.shields.io/github/stars/tigusigalpa/yandexgpt-go?style=social)
![GitHub forks](https://img.shields.io/github/forks/tigusigalpa/yandexgpt-go?style=social)
![GitHub watchers](https://img.shields.io/github/watchers/tigusigalpa/yandexgpt-go?style=social)

![GitHub issues](https://img.shields.io/github/issues/tigusigalpa/yandexgpt-go)
![GitHub pull requests](https://img.shields.io/github/issues-pr/tigusigalpa/yandexgpt-go)
![GitHub last commit](https://img.shields.io/github/last-commit/tigusigalpa/yandexgpt-go)

---

## 🏆 Acknowledgments

Thank you to all contributors and users who help improve this project!

Special thanks to:
- Yandex Cloud team for the excellent API
- Go developer community
- Everyone who reports bugs and suggests improvements

---

## 📄 Changelog

All significant changes to the project are documented in [CHANGELOG.md](CHANGELOG.md).

---

## 🔖 Keywords

YandexGPT, YandexGPT API, Go SDK, Golang, Yandex Cloud, AI, artificial intelligence, machine learning, neural networks, language models, LLM, GPT, chatbot, text generation, YandexART, image generation, AI SDK, Yandex AI, Foundation Models, natural language processing, NLP, dialogue systems, virtual assistants, automation, cloud technologies, cloud AI, Russian AI

---

## 👤 Author

**Igor Sazonov**

- Email: sovletig@gmail.com
- GitHub: [@tigusigalpa](https://github.com/tigusigalpa)
