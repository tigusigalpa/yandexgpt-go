# YandexGPT Go SDK

![YandexGPT Golang](https://i.postimg.cc/tCxgMvN4/531694893-35073f18-14ce-486f-937b-4c70e9af9e6c.jpg)

[![Go Version](https://img.shields.io/github/go-mod/go-version/tigusigalpa/yandexgpt-go)](https://github.com/tigusigalpa/yandexgpt-go)
[![License](https://img.shields.io/github/license/tigusigalpa/yandexgpt-go)](https://github.com/tigusigalpa/yandexgpt-go/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/tigusigalpa/yandexgpt-go/v2)](https://goreportcard.com/report/github.com/tigusigalpa/yandexgpt-go/v2)

English | [Русский](README.md)

[pkg.go.dev/github.com/tigusigalpa/yandexgpt-go/v2](https://pkg.go.dev/github.com/tigusigalpa/yandexgpt-go/v2)

Go SDK for YandexGPT API with text and image generation (YandexART) support.

Handles token management, provides methods for dialogues, single requests, and the Conversations API.

> **Note:** This package uses [yandex-cloud-client-go](https://github.com/tigusigalpa/yandex-cloud-client-go) for
> managing Yandex Cloud infrastructure (organizations, clouds, folders, authorization).

## Features

- Text generation (YandexGPT, YandexGPT Lite, AliceAI LLM)
- Image generation (YandexART)
- Conversations API — server-side conversation storage
- Automatic OAuth/IAM token management
- Reasoning mode (Chain of Thought)
- Test coverage

---

## Installation

Install the package using `go get`:

```bash
go get github.com/tigusigalpa/yandexgpt-go/v2
```

---

## Configuration

### 1. Getting an OAuth token

**Documentation:** [OAuth-token](https://yandex.cloud/en/docs/iam/concepts/authorization/oauth-token)

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

## Usage

### Basic usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/tigusigalpa/yandexgpt-go/v2"
    "github.com/tigusigalpa/yandexgpt-go/v2/models"
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
    
    "github.com/tigusigalpa/yandexgpt-go/v2"
    "github.com/tigusigalpa/yandexgpt-go/v2/models"
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

### Conversations API

The SDK supports the [Conversations API](https://yandex.cloud/ru/docs/ai-studio/conversations/) for managing conversations and their items on the Yandex Cloud server side.

**Available methods:**

| Method | Description |
|--------|-------------|
| `Create()` | Create a new conversation |
| `Get()` | Retrieve a conversation by ID |
| `Update()` | Update conversation metadata |
| `Delete()` | Delete a conversation |
| `CreateItems()` | Add items to a conversation |
| `ListItems()` | List conversation items |
| `GetItem()` | Retrieve a single conversation item |
| `DeleteItem()` | Delete an item from a conversation |

**Managing conversations:**

```go
package main

import (
    "fmt"
    "log"

    "github.com/tigusigalpa/yandexgpt-go/v2"
)

func main() {
    client, err := yandexgpt.NewClient("your_oauth_token", "your_folder_id")
    if err != nil {
        log.Fatal(err)
    }

    // Create a conversation with metadata
    conv, err := client.Conversations().Create(
        map[string]string{"title": "Technical Support", "user_id": "12345"},
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Conversation ID: %s\n", conv.ID)

    // Retrieve a conversation
    conv, err = client.Conversations().Get(conv.ID)
    if err != nil {
        log.Fatal(err)
    }

    // Update metadata
    conv, err = client.Conversations().Update(conv.ID, map[string]string{
        "title":  "Updated Title",
        "status": "active",
    })
    if err != nil {
        log.Fatal(err)
    }

    // Delete a conversation
    deleted, err := client.Conversations().Delete(conv.ID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Deleted: %v\n", deleted.Deleted)
}
```

**Managing conversation items:**

```go
package main

import (
    "fmt"
    "log"

    "github.com/tigusigalpa/yandexgpt-go/v2"
)

func main() {
    client, err := yandexgpt.NewClient("your_oauth_token", "your_folder_id")
    if err != nil {
        log.Fatal(err)
    }

    conversationID := "conv_123"

    // Add items to a conversation
    items, err := client.Conversations().CreateItems(conversationID, []yandexgpt.ConversationItem{
        {
            Type: "message",
            Role: "user",
            Content: []yandexgpt.ConversationContentPart{
                {Type: "input_text", Text: "Hello! How are you?"},
            },
        },
        {
            Type: "message",
            Role: "assistant",
            Content: []yandexgpt.ConversationContentPart{
                {Type: "output_text", Text: "Hi! I am doing great, how can I help?"},
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    // List items (with pagination)
    limit := 20
    order := "asc"
    list, err := client.Conversations().ListItems(conversationID, &yandexgpt.ListItemsOptions{
        Limit: &limit,
        Order: &order,
    })
    if err != nil {
        log.Fatal(err)
    }

    for _, item := range list.Data {
        if len(item.Content) > 0 {
            fmt.Printf("%s: %s\n", item.Role, item.Content[0].Text)
        }
    }

    // Pagination: get the next page
    if list.HasMore {
        nextList, _ := client.Conversations().ListItems(conversationID, &yandexgpt.ListItemsOptions{
            Limit: &limit,
            Order: &order,
            After: &list.LastID,
        })
        _ = nextList
    }

    // Retrieve a single item
    item, err := client.Conversations().GetItem(conversationID, items.FirstID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Item: %s (%s)\n", item.ID, item.Role)

    // Delete an item
    _, err = client.Conversations().DeleteItem(conversationID, items.FirstID)
    if err != nil {
        log.Fatal(err)
    }
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
    
    "github.com/tigusigalpa/yandexgpt-go/v2"
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
    
    "github.com/tigusigalpa/yandexgpt-go/v2"
    "github.com/tigusigalpa/yandexgpt-go/v2/models"
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
    
    "github.com/tigusigalpa/yandexgpt-go/v2"
    "github.com/tigusigalpa/yandexgpt-go/v2/models"
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

**Documentation:** [Reasoning Mode in YandexGPT](https://yandex.cloud/ru/docs/ai-studio/concepts/generation/chain-of-thought)

---

## Available models

| Model            | Description                                  | Constant               | Context |
|------------------|----------------------------------------------|------------------------|---------|
| `yandexgpt-lite` | Fast and economical model                    | `models.YandexGPTLite` | 32K     |
| `yandexgpt`      | Standard model                               | `models.YandexGPT`     | 32K     |
| `aliceai-llm`    | Alice AI LLM - advanced conversational model | `models.AliceAI`       | 32K     |

**Full list of models:** [Generation models in Yandex AI Studio](https://yandex.cloud/en/docs/ai-studio/concepts/generation/models)

---

## Generation parameters

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

## Error handling

```go
package main

import (
    "errors"
    "fmt"
    "log"
    
    "github.com/tigusigalpa/yandexgpt-go/v2"
    "github.com/tigusigalpa/yandexgpt-go/v2/models"
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

## Examples

See the [examples](examples/) directory for more usage examples:

- [Basic usage](examples/basic/main.go)
- [Dialogue](examples/dialogue/main.go)
- [Image generation](examples/image/main.go)
- [Custom options](examples/options/main.go)
- [Reasoning mode](examples/reasoning/main.go)

---

## Models comparison

| Model | Speed | Quality | Use case |
|-------|-------|---------|----------|
| YandexGPT Lite | Fast | Good | Chatbots, simple tasks, prototypes |
| YandexGPT | Medium | Excellent | Analysis, article generation, logic tasks |
| AliceAI LLM | Medium | Very good | Dialogues, assistants |

All models support 32K token context.

---

## Advanced features

### Dialogue context management

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

## Best practices

### Prompt optimization

- Be specific about the task and expected format
- Add examples (few-shot)
- Separate instructions from data

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

## Troubleshooting

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

## FAQ

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

## Performance

- YandexGPT Lite is 2-3x faster
- Limit dialogue history to 10-15 messages
- Process independent requests concurrently
- Cache frequent requests

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

## Security

- Store tokens in environment variables or secret stores
- Rotate OAuth tokens regularly
- Validate user input before sending to API
- Don't log tokens

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

## Testing

Run tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

---

## Documentation

- [Configuration Guide](docs/configuration.md)
- [API Reference](docs/api.md)
- [Examples](examples/)

---

## Framework integration

### Gin Web Framework

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/tigusigalpa/yandexgpt-go/v2"
    "github.com/tigusigalpa/yandexgpt-go/v2/models"
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
    "github.com/tigusigalpa/yandexgpt-go/v2"
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

## Monitoring and logging

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

## Roadmap

Done:
- YandexGPT, YandexGPT Lite, AliceAI LLM
- YandexART
- Reasoning mode (Chain of Thought)
- Automatic token management

In progress:
- Response streaming
- Function calling

Planned:
- Multimodal support (images in prompts)
- Async operations
- Embeddings
- Vector database integration

---

## Contributing

Bugs, ideas, PRs — all welcome. See [CONTRIBUTING.md](CONTRIBUTING.md).

---

## License

MIT — see [LICENSE](LICENSE).

---

## Links

**Yandex Cloud docs:**
- [YandexGPT Quickstart](https://yandex.cloud/en/docs/foundation-models/quickstart/yandexgpt)
- [API Authentication](https://yandex.cloud/en/docs/iam/concepts/authorization/iam-token)
- [Resource Management](https://yandex.cloud/en/docs/resource-manager/)
- [API Foundation Models](https://yandex.cloud/en/docs/foundation-models/concepts/api)
- [YandexGPT Pricing](https://yandex.cloud/en/docs/foundation-models/pricing)
- [YandexART](https://yandex.cloud/en/docs/foundation-models/concepts/yandexart)
- [Reasoning Mode](https://yandex.cloud/en/docs/ai-studio/concepts/generation/chain-of-thought)
- [Limits and Quotas](https://yandex.cloud/en/docs/foundation-models/concepts/limits)

**Community:**
- [Yandex Cloud Telegram](https://t.me/yandexcloud)
- [Training Materials](https://yandex.cloud/en/training)

**Related projects:**
- [yandex-cloud-client-go](https://github.com/tigusigalpa/yandex-cloud-client-go) — Yandex Cloud management client
- [Official Yandex Cloud Go SDK](https://github.com/yandex-cloud/go-sdk)

---

## Changelog

All changes in [CHANGELOG.md](CHANGELOG.md).

---

## Author

**Igor Sazonov**

- Email: sovletig@gmail.com
- GitHub: [@tigusigalpa](https://github.com/tigusigalpa)
