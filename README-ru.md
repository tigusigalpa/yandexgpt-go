# YandexGPT Go SDK

![YandexGPT Golang](https://github.com/user-attachments/assets/35073f18-14ce-486f-937b-4c70e9af9e6c)

[![Go Version](https://img.shields.io/github/go-mod/go-version/tigusigalpa/yandexgpt-go)](https://github.com/tigusigalpa/yandexgpt-go)
[![License](https://img.shields.io/github/license/tigusigalpa/yandexgpt-go)](https://github.com/tigusigalpa/yandexgpt-go/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/tigusigalpa/yandexgpt-go)](https://goreportcard.com/report/github.com/tigusigalpa/yandexgpt-go)

**🌐 Язык:** Русский | [English](README.md)

**📦 Package:** [pkg.go.dev/github.com/tigusigalpa/yandexgpt-go](https://pkg.go.dev/github.com/tigusigalpa/yandexgpt-go)

Полнофункциональный Go/Golang SDK для работы с YandexGPT API. Пакет предоставляет удобный интерфейс для интеграции с AI
моделями Yandex Cloud, включая поддержку YandexART.

> **Примечание:** Пакет использует [yandex-cloud-client-go](https://github.com/tigusigalpa/yandex-cloud-client-go) для
> управления облачной инфраструктурой Yandex Cloud (организации, облака, каталоги, авторизация).

## 🚀 Возможности

- 🔌 Простая интеграция с YandexGPT API
- 🔨 **Поддержка YandexART**
- 🔐 Автоматическое управление OAuth и IAM токенами
- 🎯 Поддержка всех доступных моделей YandexGPT
- 📝 Поддержка диалогов и одиночных запросов
- ⚡ Автоматическое обновление токенов
- 🧪 Покрытие тестами
- 📚 Подробная документация
- 🚀 Высокая производительность и поддержка конкурентности

---

## 📦 Установка

Установите пакет с помощью `go get`:

```bash
go get github.com/tigusigalpa/yandexgpt-go
```

---

## ⚙️ Настройка

### 1. Получение OAuth токена

📚 **Документация:** [OAuth-токен](https://yandex.cloud/ru/docs/iam/concepts/authorization/oauth-token)

Перейдите по ссылке
для [получения OAuth токена](https://oauth.yandex.ru/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb):

```
https://oauth.yandex.ru/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb
```

### 2. Настройка окружения

Установите переменные окружения или передайте их напрямую в клиент:

```bash
export YANDEX_GPT_OAUTH_TOKEN=your_oauth_token_here
export YANDEX_GPT_FOLDER_ID=your_folder_id_here
```

### 3. Подготовка Yandex Cloud

Для подробных инструкций по настройке см. [Руководство по настройке](docs/configuration-ru.md).

---

## 💡 Использование

### Базовое использование

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/tigusigalpa/yandexgpt-go"
    "github.com/tigusigalpa/yandexgpt-go/models"
)

func main() {
    // Создание клиента
    client, err := yandexgpt.NewClient("your_oauth_token", "your_folder_id")
    if err != nil {
        log.Fatal(err)
    }
    
    // Простой запрос
    response, err := client.GenerateText(
        "Расскажи о преимуществах языка программирования Go",
        models.YandexGPTLite,
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(response.Result.Alternatives[0].Message.Text)
}
```

### Работа с диалогами

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
            Text: "Ты полезный помощник-программист",
        },
        {
            Role: "user",
            Text: "Как создать REST API на Go?",
        },
        {
            Role: "assistant",
            Text: "Для создания REST API на Go можно использовать стандартный пакет net/http...",
        },
        {
            Role: "user",
            Text: "А как добавить валидацию?",
        },
    }
    
    response, err := client.GenerateFromMessages(messages, models.YandexGPTLite, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(response.Result.Alternatives[0].Message.Text)
}
```

### Генерация изображений с YandexART

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
    
    // Генерация изображения
    result, err := client.GenerateImage(
        "Красивый закат над горами",
        nil,
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    
    // Декодирование base64 изображения
    imageData, err := base64.StdEncoding.DecodeString(result.ImageBase64)
    if err != nil {
        log.Fatal(err)
    }
    
    // Сохранение в файл
    err = os.WriteFile("output.jpg", imageData, 0644)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Изображение сохранено в output.jpg")
}
```

### Пользовательские параметры

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
        "Напиши стихотворение о программировании",
        models.YandexGPT,
        options,
    )
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(response.Result.Alternatives[0].Message.Text)
}
```

### Режим рассуждений

Режим рассуждений позволяет моделям выполнять цепочку рассуждений для решения сложных задач:

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
        "Решите логическую задачу: Если все розы - цветы, и некоторые цветы - красные, можно ли утверждать, что некоторые розы - красные?",
        models.YandexGPT,
        options,
    )
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(response.Result.Alternatives[0].Message.Text)
    
    // Проверка использования токенов рассуждения
    if response.Result.Usage.ReasoningTokens > 0 {
        fmt.Printf("Использовано токенов рассуждения: %d\n", response.Result.Usage.ReasoningTokens)
    }
}
```

**Параметры режима рассуждений:**

- `Mode`: `"DISABLED"` (по умолчанию), `"ENABLED_HIDDEN"` (включает рассуждения без показа цепочки)
- `Effort`: `"low"`, `"medium"`, `"high"` (управляет глубиной рассуждений, опционально)

📚 **Документация:
** [Режим рассуждений в YandexGPT](https://yandex.cloud/ru/docs/ai-studio/concepts/generation/chain-of-thought)

---

## 🤖 Доступные модели

| Модель           | Описание                                      | Константа              | Контекст |
|------------------|-----------------------------------------------|------------------------|----------|
| `yandexgpt-lite` | Быстрая и экономичная модель                  | `models.YandexGPTLite` | 32K      |
| `yandexgpt`      | Стандартная модель                            | `models.YandexGPT`     | 32K      |
| `aliceai-llm`    | Alice AI LLM - продвинутая разговорная модель | `models.AliceAI`       | 32K      |

📚 **Полный список доступных моделей:
** [Модели генерации в Yandex AI Studio](https://yandex.cloud/ru/docs/ai-studio/concepts/generation/models)

---

## 🔧 Параметры генерации

```go
type CompletionOptions struct {
    Stream           bool              // Потоковая передача (пока не поддерживается)
    Temperature      float64           // Креативность (0.0 - 1.0)
    MaxTokens        int               // Максимальное количество токенов
    ReasoningOptions *ReasoningOptions // Настройки режима рассуждений (опционально)
}

type ReasoningOptions struct {
    Mode   string  // "DISABLED", "ENABLED_HIDDEN"
    Effort *string // "low", "medium", "high" (опционально)
}
```

---

## ⚠️ Обработка ошибок

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
    
    response, err := client.GenerateText("Привет!", models.YandexGPTLite, nil)
    if err != nil {
        var authErr *yandexgpt.AuthenticationError
        var apiErr *yandexgpt.APIError
        
        switch {
        case errors.As(err, &authErr):
            fmt.Printf("Ошибка аутентификации: %v\n", err)
        case errors.As(err, &apiErr):
            fmt.Printf("Ошибка API: %v\n", err)
        default:
            fmt.Printf("Неизвестная ошибка: %v\n", err)
        }
        return
    }
    
    fmt.Println(response.Result.Alternatives[0].Message.Text)
}
```

---

## 📚 Примеры

См. директорию [examples](examples/) для дополнительных примеров использования:

- [Базовое использование](examples/basic/main.go)
- [Диалог](examples/dialogue/main.go)
- [Генерация изображений](examples/image/main.go)
- [Пользовательские параметры](examples/options/main.go)
- [Режим рассуждений](examples/reasoning/main.go)

---

## 🧪 Тестирование

Запуск тестов:

```bash
go test ./...
```

Запуск тестов с покрытием:

```bash
go test -cover ./...
```

---

## 📖 Документация

Для подробной документации см.:

- [Руководство по настройке](docs/configuration-ru.md)
- [Справочник API](docs/api-ru.md)
- [Примеры](examples/)

---

## 🤝 Вклад в проект

Вклад приветствуется! Пожалуйста, прочитайте [CONTRIBUTING.md](CONTRIBUTING.md) для деталей.

---

## 📝 Лицензия

Этот проект лицензирован под лицензией MIT - см. файл [LICENSE](LICENSE) для деталей.

---

## 🔗 Ссылки

- 📖 [Быстрый старт YandexGPT](https://yandex.cloud/ru/docs/foundation-models/quickstart/yandexgpt)
- 🔑 [Аутентификация в API](https://yandex.cloud/ru/docs/iam/concepts/authorization/iam-token)
- 🏗️ [Управление ресурсами](https://yandex.cloud/ru/docs/resource-manager/)
- 🤖 [API Foundation Models](https://yandex.cloud/ru/docs/foundation-models/concepts/api)
- 💰 [Тарифы YandexGPT](https://yandex.cloud/ru/docs/foundation-models/pricing)

---

## 👤 Автор

**Igor Sazonov**

- Email: sovletig@gmail.com
- GitHub: [@tigusigalpa](https://github.com/tigusigalpa)
