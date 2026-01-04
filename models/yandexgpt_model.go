package models

import "fmt"

const (
	YandexGPTLite = "yandexgpt-lite"
	YandexGPT     = "yandexgpt"
	AliceAI       = "aliceai-llm"
)

var modelDescriptions = map[string]string{
	YandexGPTLite: "Lightweight version of YandexGPT for simple tasks",
	YandexGPT:     "Standard YandexGPT model",
	AliceAI:       "Alice AI LLM - advanced conversational model with 32K context",
}

func IsValidModel(model string) bool {
	_, exists := modelDescriptions[model]
	return exists
}

func GetAllModels() []string {
	return []string{
		YandexGPTLite,
		YandexGPT,
		AliceAI,
	}
}

func GetModelDescriptions() map[string]string {
	result := make(map[string]string, len(modelDescriptions))
	for k, v := range modelDescriptions {
		result[k] = v
	}
	return result
}

func GetModelURI(model, folderID string) string {
	return fmt.Sprintf("gpt://%s/%s", folderID, model)
}
