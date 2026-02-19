# YandexGPT Go SDK

![YandexGPT Golang](https://github.com/user-attachments/assets/35073f18-14ce-486f-937b-4c70e9af9e6c)

[![Go Version](https://img.shields.io/github/go-mod/go-version/tigusigalpa/yandexgpt-go)](https://github.com/tigusigalpa/yandexgpt-go)
[![License](https://img.shields.io/github/license/tigusigalpa/yandexgpt-go)](https://github.com/tigusigalpa/yandexgpt-go/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/tigusigalpa/yandexgpt-go)](https://goreportcard.com/report/github.com/tigusigalpa/yandexgpt-go)

**🌐 Язык:** Русский | [English](README-en.md)

**📦 Package:** [pkg.go.dev/github.com/tigusigalpa/yandexgpt-go](https://pkg.go.dev/github.com/tigusigalpa/yandexgpt-go)

Полнофункциональный Go/Golang SDK для работы с YandexGPT API. Пакет предоставляет удобный интерфейс для интеграции с AI
моделями Yandex Cloud, включая поддержку YandexART для генерации изображений.

**YandexGPT Go SDK** - это мощная библиотека для разработчиков на языке Go, которая упрощает интеграцию с искусственным интеллектом от Yandex Cloud. Библиотека позволяет создавать интеллектуальные приложения, чат-боты, системы автоматизации и AI-ассистенты с использованием передовых языковых моделей YandexGPT, YandexGPT Lite и AliceAI LLM.

### Почему выбирают YandexGPT Go SDK?

- **🚀 Быстрая разработка**: Готовые к использованию методы для генерации текста, ведения диалогов и создания изображений
- **🔒 Безопасность**: Автоматическое управление токенами OAuth и IAM с поддержкой их обновления
- **⚡ Производительность**: Оптимизированная работа с API, поддержка конкурентных запросов
- **🎨 Универсальность**: Поддержка текстовых моделей YandexGPT и генерации изображений через YandexART
- **📖 Документация**: Подробные примеры кода и руководства для быстрого старта
- **🧪 Надежность**: Полное покрытие тестами и активная поддержка

### Основные сценарии использования

- **Чат-боты и виртуальные ассистенты**: Создание интеллектуальных диалоговых систем для поддержки клиентов
- **Генерация контента**: Автоматическое создание текстов, статей, описаний товаров
- **Анализ и обработка текста**: Суммаризация, классификация, извлечение информации
- **Генерация изображений**: Создание визуального контента на основе текстовых описаний
- **Автоматизация бизнес-процессов**: Интеллектуальная обработка документов и данных
- **Образовательные платформы**: AI-помощники для обучения и тестирования

> **Примечание:** Пакет использует [yandex-cloud-client-go](https://github.com/tigusigalpa/yandex-cloud-client-go) для
> управления облачной инфраструктурой Yandex Cloud (организации, облака, каталоги, авторизация).

## 🚀 Возможности

- 🔌 Простая интеграция с YandexGPT API
- 🔨 **Поддержка YandexART**
- 🔐 Автоматическое управление OAuth и IAM токенами
- 🎯 Поддержка всех доступных моделей YandexGPT
- 📝 Поддержка диалогов и одиночных запросов
- 🗂️ **Conversations API** — управление диалогами на стороне сервера
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

### Conversations API

SDK поддерживает работу с [Conversations API](https://yandex.cloud/ru/docs/ai-studio/conversations/) для управления диалогами и их элементами на стороне сервера Yandex Cloud.

**Доступные методы:**

| Метод | Описание |
|-------|----------|
| `Create()` | Создание нового диалога |
| `Get()` | Получение диалога по ID |
| `Update()` | Обновление метаданных диалога |
| `Delete()` | Удаление диалога |
| `CreateItems()` | Добавление элементов в диалог |
| `ListItems()` | Получение списка элементов диалога |
| `GetItem()` | Получение одного элемента диалога |
| `DeleteItem()` | Удаление элемента из диалога |

**Управление диалогами:**

```go
package main

import (
    "fmt"
    "log"

    "github.com/tigusigalpa/yandexgpt-go"
)

func main() {
    client, err := yandexgpt.NewClient("your_oauth_token", "your_folder_id")
    if err != nil {
        log.Fatal(err)
    }

    // Создание диалога с метаданными
    conv, err := client.Conversations().Create(
        map[string]string{"title": "Техническая поддержка", "user_id": "12345"},
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Conversation ID: %s\n", conv.ID)

    // Получение диалога
    conv, err = client.Conversations().Get(conv.ID)
    if err != nil {
        log.Fatal(err)
    }

    // Обновление метаданных
    conv, err = client.Conversations().Update(conv.ID, map[string]string{
        "title":  "Обновлённый заголовок",
        "status": "active",
    })
    if err != nil {
        log.Fatal(err)
    }

    // Удаление диалога
    deleted, err := client.Conversations().Delete(conv.ID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Deleted: %v\n", deleted.Deleted)
}
```

**Управление элементами диалога:**

```go
package main

import (
    "fmt"
    "log"

    "github.com/tigusigalpa/yandexgpt-go"
)

func main() {
    client, err := yandexgpt.NewClient("your_oauth_token", "your_folder_id")
    if err != nil {
        log.Fatal(err)
    }

    conversationID := "conv_123"

    // Добавление элементов в диалог
    items, err := client.Conversations().CreateItems(conversationID, []yandexgpt.ConversationItem{
        {
            Type: "message",
            Role: "user",
            Content: []yandexgpt.ConversationContentPart{
                {Type: "input_text", Text: "Привет! Как дела?"},
            },
        },
        {
            Type: "message",
            Role: "assistant",
            Content: []yandexgpt.ConversationContentPart{
                {Type: "output_text", Text: "Здравствуйте! Всё отлично, чем могу помочь?"},
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    // Получение списка элементов (с пагинацией)
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

    // Пагинация: получение следующей страницы
    if list.HasMore {
        nextList, _ := client.Conversations().ListItems(conversationID, &yandexgpt.ListItemsOptions{
            Limit: &limit,
            Order: &order,
            After: &list.LastID,
        })
        _ = nextList
    }

    // Получение одного элемента
    item, err := client.Conversations().GetItem(conversationID, items.FirstID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Item: %s (%s)\n", item.ID, item.Role)

    // Удаление элемента
    _, err = client.Conversations().DeleteItem(conversationID, items.FirstID)
    if err != nil {
        log.Fatal(err)
    }
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

## 🔍 Сравнение моделей YandexGPT

### Выбор подходящей модели для вашей задачи

| Характеристика | YandexGPT Lite | YandexGPT | AliceAI LLM |
|----------------|----------------|-----------|-------------|
| **Скорость** | ⚡⚡⚡ Очень быстрая | ⚡⚡ Быстрая | ⚡⚡ Быстрая |
| **Качество** | ⭐⭐⭐ Хорошее | ⭐⭐⭐⭐⭐ Отличное | ⭐⭐⭐⭐ Очень хорошее |
| **Стоимость** | 💰 Экономичная | 💰💰 Стандартная | 💰💰 Стандартная |
| **Контекст** | 32K токенов | 32K токенов | 32K токенов |
| **Применение** | Простые задачи, чат-боты | Сложные задачи, анализ | Диалоги, ассистенты |
| **Креативность** | Средняя | Высокая | Высокая |

### Рекомендации по выбору модели

**YandexGPT Lite** - идеальна для:
- Быстрых ответов в чат-ботах
- Простой классификации текста
- Генерации коротких текстов
- Приложений с высокой нагрузкой
- Прототипирования и тестирования

**YandexGPT** - рекомендуется для:
- Сложного анализа текста
- Генерации длинных статей
- Творческих задач (стихи, истории)
- Технической документации
- Решения логических задач

**AliceAI LLM** - оптимальна для:
- Естественных диалогов
- Виртуальных ассистентов
- Персонализированных рекомендаций
- Контекстно-зависимых ответов
- Эмоционально окрашенного общения

---

## 🎯 Продвинутые возможности

### Управление контекстом диалога

Для создания качественных диалоговых систем важно правильно управлять контекстом:

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
        maxHistory: 10, // Хранить последние 10 сообщений
    }
}

func (dm *DialogManager) AddUserMessage(text string) error {
    dm.messages = append(dm.messages, yandexgpt.Message{
        Role: "user",
        Text: text,
    })
    
    // Ограничение истории для экономии токенов
    if len(dm.messages) > dm.maxHistory {
        // Сохраняем системное сообщение и последние N сообщений
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

### Пакетная обработка запросов

Для обработки множества запросов используйте конкурентность:

```go
func ProcessBatch(client *yandexgpt.Client, prompts []string) ([]string, error) {
    results := make([]string, len(prompts))
    errors := make([]error, len(prompts))
    
    var wg sync.WaitGroup
    semaphore := make(chan struct{}, 5) // Ограничение: 5 одновременных запросов
    
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
    
    // Проверка ошибок
    for _, err := range errors {
        if err != nil {
            return results, err
        }
    }
    
    return results, nil
}
```

### Работа с большими текстами

Для обработки текстов, превышающих лимит контекста:

```go
func SummarizeLongText(client *yandexgpt.Client, longText string) (string, error) {
    // Разбиваем текст на части по ~3000 символов
    chunkSize := 3000
    var chunks []string
    
    for i := 0; i < len(longText); i += chunkSize {
        end := i + chunkSize
        if end > len(longText) {
            end = len(longText)
        }
        chunks = append(chunks, longText[i:end])
    }
    
    // Суммаризируем каждую часть
    var summaries []string
    for _, chunk := range chunks {
        response, err := client.GenerateText(
            fmt.Sprintf("Кратко перескажи следующий текст:\n\n%s", chunk),
            models.YandexGPT,
            &yandexgpt.CompletionOptions{Temperature: 0.3},
        )
        if err != nil {
            return "", err
        }
        summaries = append(summaries, response.Result.Alternatives[0].Message.Text)
    }
    
    // Финальная суммаризация
    finalPrompt := fmt.Sprintf(
        "Объедини следующие краткие пересказы в один связный текст:\n\n%s",
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

## 💡 Лучшие практики

### Оптимизация промптов

1. **Будьте конкретны**: Четко формулируйте задачу и ожидаемый формат ответа
2. **Используйте примеры**: Few-shot learning повышает качество результатов
3. **Структурируйте запросы**: Разделяйте инструкции и данные
4. **Контролируйте длину**: Оптимизируйте использование токенов

```go
// ❌ Плохой промпт
prompt := "Расскажи про Go"

// ✅ Хороший промпт
prompt := `Напиши краткое описание языка программирования Go (до 100 слов).
Включи следующие аспекты:
1. Основное назначение
2. Ключевые особенности
3. Популярные области применения

Формат: структурированный текст с подзаголовками.`
```

### Управление температурой

- **0.0-0.3**: Детерминированные, точные ответы (анализ, классификация)
- **0.4-0.7**: Сбалансированные ответы (общие задачи)
- **0.8-1.0**: Креативные, разнообразные ответы (генерация контента)

### Обработка ошибок и повторные попытки

```go
func GenerateWithRetry(client *yandexgpt.Client, prompt string, maxRetries int) (string, error) {
    var lastErr error
    
    for attempt := 0; attempt < maxRetries; attempt++ {
        response, err := client.GenerateText(prompt, models.YandexGPT, nil)
        if err == nil {
            return response.Result.Alternatives[0].Message.Text, nil
        }
        
        lastErr = err
        
        // Экспоненциальная задержка
        backoff := time.Duration(math.Pow(2, float64(attempt))) * time.Second
        time.Sleep(backoff)
    }
    
    return "", fmt.Errorf("failed after %d attempts: %w", maxRetries, lastErr)
}
```

### Кэширование результатов

```go
type CachedClient struct {
    client *yandexgpt.Client
    cache  map[string]string
    mu     sync.RWMutex
}

func (cc *CachedClient) GenerateText(prompt string) (string, error) {
    // Проверяем кэш
    cc.mu.RLock()
    if cached, ok := cc.cache[prompt]; ok {
        cc.mu.RUnlock()
        return cached, nil
    }
    cc.mu.RUnlock()
    
    // Генерируем новый ответ
    response, err := cc.client.GenerateText(prompt, models.YandexGPT, nil)
    if err != nil {
        return "", err
    }
    
    result := response.Result.Alternatives[0].Message.Text
    
    // Сохраняем в кэш
    cc.mu.Lock()
    cc.cache[prompt] = result
    cc.mu.Unlock()
    
    return result, nil
}
```

---

## 🔧 Устранение неполадок

### Частые проблемы и решения

#### Ошибка аутентификации

**Проблема**: `Authentication error: invalid OAuth token`

**Решение**:
1. Проверьте срок действия OAuth токена (действителен 1 год)
2. Получите новый токен по [ссылке](https://oauth.yandex.ru/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb)
3. Убедитесь, что токен правильно передан в клиент

```go
// Проверка токена
if token == "" {
    log.Fatal("OAuth token is required")
}
```

#### Ошибка доступа к каталогу

**Проблема**: `Permission denied for folder`

**Решение**:
1. Убедитесь, что у вас есть права на каталог в Yandex Cloud
2. Проверьте правильность folder_id
3. Назначьте роль `ai.languageModels.user` вашему аккаунту

```bash
yc resource-manager folder add-access-binding <folder-id> \
  --role ai.languageModels.user \
  --subject userAccount:<user-id>
```

#### Превышение лимита токенов

**Проблема**: `Maximum token limit exceeded`

**Решение**:
1. Сократите длину промпта или истории диалога
2. Используйте параметр `MaxTokens` для ограничения ответа
3. Разбейте большой текст на части

```go
options := &yandexgpt.CompletionOptions{
    MaxTokens: 1000, // Ограничение длины ответа
}
```

#### Медленные ответы

**Проблема**: Долгое время ответа API

**Решение**:
1. Используйте YandexGPT Lite для простых задач
2. Оптимизируйте длину промптов
3. Рассмотрите кэширование частых запросов
4. Используйте конкурентность для пакетной обработки

#### Низкое качество ответов

**Проблема**: Ответы не соответствуют ожиданиям

**Решение**:
1. Улучшите формулировку промпта (добавьте примеры, контекст)
2. Используйте YandexGPT вместо Lite для сложных задач
3. Настройте температуру (понизьте для точности, повысьте для креативности)
4. Добавьте системное сообщение с инструкциями

---

## ❓ Часто задаваемые вопросы (FAQ)

### Общие вопросы

**Q: Нужен ли мне платный аккаунт Yandex Cloud?**

A: Да, для использования YandexGPT API необходим аккаунт Yandex Cloud с привязанной платежной информацией. Доступен пробный период с бесплатными грантами.

**Q: Какие лимиты на использование API?**

A: Лимиты зависят от вашего тарифа. По умолчанию:
- 200 запросов в минуту
- 32K токенов контекста
- Ограничения по количеству токенов в месяц согласно тарифу

**Q: Поддерживается ли потоковая передача (streaming)?**

A: В текущей версии SDK потоковая передача находится в разработке. Следите за обновлениями в [CHANGELOG.md](CHANGELOG.md).

**Q: Можно ли использовать SDK в production?**

A: Да, SDK готов для использования в production. Рекомендуется:
- Реализовать обработку ошибок и retry логику
- Настроить мониторинг и логирование
- Использовать кэширование для оптимизации
- Следить за расходом токенов

### Технические вопросы

**Q: Как хранить OAuth токен безопасно?**

A: Используйте переменные окружения или секретные хранилища:

```go
// Из переменных окружения
token := os.Getenv("YANDEX_GPT_OAUTH_TOKEN")

// Из файла конфигурации (не коммитить в git!)
// Или используйте vault-системы (HashiCorp Vault, AWS Secrets Manager)
```

**Q: Как отслеживать использование токенов?**

A: Проверяйте поле `Usage` в ответе:

```go
response, err := client.GenerateText(prompt, models.YandexGPT, nil)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Input tokens: %d\n", response.Result.Usage.InputTextTokens)
fmt.Printf("Output tokens: %d\n", response.Result.Usage.CompletionTokens)
fmt.Printf("Total tokens: %d\n", response.Result.Usage.TotalTokens)
```

**Q: Можно ли использовать несколько моделей одновременно?**

A: Да, просто указывайте нужную модель в каждом запросе:

```go
// Быстрый ответ
quickResponse, _ := client.GenerateText(prompt, models.YandexGPTLite, nil)

// Детальный анализ
detailedResponse, _ := client.GenerateText(prompt, models.YandexGPT, nil)
```

**Q: Как обрабатывать rate limiting?**

A: Реализуйте exponential backoff и проверяйте заголовки ответа:

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

## ⚡ Производительность и оптимизация

### Советы по оптимизации

1. **Выбор модели**: Используйте YandexGPT Lite для простых задач - она в 2-3 раза быстрее
2. **Управление контекстом**: Ограничивайте историю диалога 10-15 последними сообщениями
3. **Параллелизм**: Обрабатывайте независимые запросы конкурентно
4. **Кэширование**: Кэшируйте частые или идентичные запросы
5. **Пулы соединений**: Переиспользуйте HTTP клиенты

### Мониторинг производительности

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

### Оценка стоимости

Рассчитывайте стоимость использования:

```go
func EstimateCost(inputTokens, outputTokens int, model string) float64 {
    // Примерные цены (проверяйте актуальные на сайте Yandex Cloud)
    var pricePerMillion float64
    
    switch model {
    case "yandexgpt-lite":
        pricePerMillion = 0.4 // $0.4 за 1M токенов
    case "yandexgpt":
        pricePerMillion = 1.2 // $1.2 за 1M токенов
    default:
        pricePerMillion = 1.2
    }
    
    totalTokens := float64(inputTokens + outputTokens)
    return (totalTokens / 1_000_000) * pricePerMillion
}
```

---

## 🔐 Безопасность

### Рекомендации по безопасности

1. **Никогда не храните токены в коде**: Используйте переменные окружения или секретные хранилища
2. **Ротация токенов**: Регулярно обновляйте OAuth токены
3. **Валидация входных данных**: Проверяйте пользовательский ввод перед отправкой в API
4. **Ограничение доступа**: Используйте минимально необходимые права в Yandex Cloud
5. **Логирование**: Не логируйте чувствительные данные и токены

### Пример безопасной конфигурации

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
    
    // Валидация формата
    if len(token) < 20 {
        return nil, errors.New("invalid token format")
    }
    
    return &SecureConfig{
        OAuthToken: token,
        FolderID:   folderID,
    }, nil
}
```

### Фильтрация контента

```go
func SanitizeInput(input string) string {
    // Удаление потенциально опасных символов
    input = strings.TrimSpace(input)
    
    // Ограничение длины
    maxLength := 10000
    if len(input) > maxLength {
        input = input[:maxLength]
    }
    
    return input
}
```

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

## 🌟 Интеграция с популярными фреймворками

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

### gRPC сервис

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

## 📊 Мониторинг и логирование

### Интеграция с Prometheus

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

### Структурированное логирование

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

## 🚀 Roadmap и будущие возможности

### Планируемые функции

- ✅ Поддержка YandexGPT и YandexGPT Lite
- ✅ Интеграция с YandexART
- ✅ Режим рассуждений (Chain of Thought)
- ✅ Автоматическое управление токенами
- 🔄 Потоковая передача ответов (Streaming)
- 🔄 Поддержка функций (Function Calling)
- 📋 Поддержка мультимодальности (изображения в промптах)
- 📋 Асинхронные операции
- 📋 Расширенная работа с эмбеддингами
- 📋 Интеграция с векторными базами данных

### Как внести вклад

Мы приветствуем вклад сообщества! Вы можете помочь:

- 🐛 Сообщать об ошибках через [Issues](https://github.com/tigusigalpa/yandexgpt-go/issues)
- 💡 Предлагать новые функции
- 📝 Улучшать документацию
- 🔧 Отправлять Pull Requests
- ⭐ Ставить звезду проекту на GitHub

---

## 🤝 Вклад в проект

Вклад приветствуется! Пожалуйста, прочитайте [CONTRIBUTING.md](CONTRIBUTING.md) для деталей.

---

## 📝 Лицензия

Этот проект лицензирован под лицензией MIT - см. файл [LICENSE](LICENSE) для деталей.

---

## 🔗 Полезные ссылки

### Официальная документация Yandex Cloud

- 📖 [Быстрый старт YandexGPT](https://yandex.cloud/ru/docs/foundation-models/quickstart/yandexgpt)
- 🔑 [Аутентификация в API](https://yandex.cloud/ru/docs/iam/concepts/authorization/iam-token)
- 🏗️ [Управление ресурсами](https://yandex.cloud/ru/docs/resource-manager/)
- 🤖 [API Foundation Models](https://yandex.cloud/ru/docs/foundation-models/concepts/api)
- 💰 [Тарифы YandexGPT](https://yandex.cloud/ru/docs/foundation-models/pricing)
- 🎨 [YandexART документация](https://yandex.cloud/ru/docs/foundation-models/concepts/yandexart)
- 🧠 [Режим рассуждений](https://yandex.cloud/ru/docs/ai-studio/concepts/generation/chain-of-thought)
- 📊 [Лимиты и квоты](https://yandex.cloud/ru/docs/foundation-models/concepts/limits)

### Сообщество и поддержка

- 💬 [Telegram сообщество Yandex Cloud](https://t.me/yandexcloud)
- 🎓 [Обучающие материалы](https://yandex.cloud/ru/training)
- 📺 [YouTube канал Yandex Cloud](https://www.youtube.com/@YandexCloudRu)
- 📰 [Блог Yandex Cloud](https://yandex.cloud/ru/blog)

### Связанные проекты

- 🔧 [yandex-cloud-client-go](https://github.com/tigusigalpa/yandex-cloud-client-go) - Клиент для управления Yandex Cloud
- 📦 [Официальный Go SDK Yandex Cloud](https://github.com/yandex-cloud/go-sdk)
- 🐍 [Python SDK для YandexGPT](https://github.com/yandex-cloud/python-sdk)

### Примеры и туториалы

- 🎯 [Создание чат-бота с YandexGPT](examples/chatbot/)
- 🖼️ [Генератор изображений с YandexART](examples/image-generator/)
- 📝 [Система генерации контента](examples/content-generator/)
- 🔍 [Анализатор текста](examples/text-analyzer/)

---

## 📈 Статистика проекта

![GitHub stars](https://img.shields.io/github/stars/tigusigalpa/yandexgpt-go?style=social)
![GitHub forks](https://img.shields.io/github/forks/tigusigalpa/yandexgpt-go?style=social)
![GitHub watchers](https://img.shields.io/github/watchers/tigusigalpa/yandexgpt-go?style=social)

![GitHub issues](https://img.shields.io/github/issues/tigusigalpa/yandexgpt-go)
![GitHub pull requests](https://img.shields.io/github/issues-pr/tigusigalpa/yandexgpt-go)
![GitHub last commit](https://img.shields.io/github/last-commit/tigusigalpa/yandexgpt-go)

---

## 🏆 Благодарности

Благодарим всех контрибьюторов и пользователей, которые помогают улучшать этот проект!

Особая благодарность:
- Команде Yandex Cloud за отличное API
- Сообществу Go разработчиков
- Всем, кто сообщает об ошибках и предлагает улучшения

---

## 📄 Changelog

Все значимые изменения в проекте документируются в [CHANGELOG.md](CHANGELOG.md).

---

## 🔖 Ключевые слова

YandexGPT, YandexGPT API, Go SDK, Golang, Yandex Cloud, AI, искусственный интеллект, машинное обучение, нейросети, языковые модели, LLM, GPT, чат-бот, генерация текста, YandexART, генерация изображений, Conversations API, AI SDK, Yandex AI, Foundation Models, natural language processing, NLP, диалоговые системы, виртуальные ассистенты, автоматизация, облачные технологии, cloud AI, Russian AI, российский AI

---

## 👤 Автор

**Igor Sazonov**

- Email: sovletig@gmail.com
- GitHub: [@tigusigalpa](https://github.com/tigusigalpa)
