# YandexGPT Go SDK

![YandexGPT Golang](https://github.com/user-attachments/assets/35073f18-14ce-486f-937b-4c70e9af9e6c)

[![Go Version](https://img.shields.io/github/go-mod/go-version/tigusigalpa/yandexgpt-go)](https://github.com/tigusigalpa/yandexgpt-go)
[![License](https://img.shields.io/github/license/tigusigalpa/yandexgpt-go)](https://github.com/tigusigalpa/yandexgpt-go/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/tigusigalpa/yandexgpt-go)](https://goreportcard.com/report/github.com/tigusigalpa/yandexgpt-go)

**🌐 Language:** English | [Русский](README-ru.md)

A full-featured Go/Golang SDK for working with YandexGPT API. The package provides a convenient interface for integrating with Yandex Cloud AI models, including YandexART support.

> **Note:** This package uses [yandex-cloud-client-go](https://github.com/tigusigalpa/yandex-cloud-client-go) for managing Yandex Cloud infrastructure (organizations, clouds, folders, authorization).

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

Follow the link to [get an OAuth token](https://oauth.yandex.com/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb):

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

---

## 🤖 Available models

| Model            | Description                                  | Constant                | Context |
|------------------|----------------------------------------------|-------------------------|---------|
| `yandexgpt-lite` | Fast and economical model                    | `models.YandexGPTLite`  | 32K     |
| `yandexgpt`      | Standard model                               | `models.YandexGPT`      | 32K     |
| `aliceai-llm`    | Alice AI LLM - advanced conversational model | `models.AliceAI`        | 32K     |

📚 **Complete list of available models:** [Generation models in Yandex AI Studio](https://yandex.cloud/en/docs/ai-studio/concepts/generation/models)

---

## 🔧 Generation parameters

```go
type CompletionOptions struct {
    Stream      bool    // Streaming (not yet supported)
    Temperature float64 // Creativity (0.0 - 1.0)
    MaxTokens   int     // Maximum number of tokens
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

## 🤝 Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details.

---

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🔗 Links

- 📖 [YandexGPT Quickstart](https://yandex.cloud/en/docs/foundation-models/quickstart/yandexgpt)
- 🔑 [API Authentication](https://yandex.cloud/en/docs/iam/concepts/authorization/iam-token)
- 🏗️ [Resource Management](https://yandex.cloud/en/docs/resource-manager/)
- 🤖 [API Foundation Models](https://yandex.cloud/en/docs/foundation-models/concepts/api)
- 💰 [YandexGPT Pricing](https://yandex.cloud/en/docs/foundation-models/pricing)

---

## 👤 Author

**Igor Sazonov**
- Email: sovletig@gmail.com
- GitHub: [@tigusigalpa](https://github.com/tigusigalpa)
