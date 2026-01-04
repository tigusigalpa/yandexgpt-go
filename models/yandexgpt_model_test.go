package models

import "testing"

func TestIsValidModel(t *testing.T) {
	tests := []struct {
		name     string
		model    string
		expected bool
	}{
		{"Valid YandexGPTLite", YandexGPTLite, true},
		{"Valid YandexGPT", YandexGPT, true},
		{"Valid AliceAI", AliceAI, true},
		{"Invalid model", "invalid-model", false},
		{"Empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidModel(tt.model)
			if result != tt.expected {
				t.Errorf("IsValidModel(%s) = %v, expected %v", tt.model, result, tt.expected)
			}
		})
	}
}

func TestGetAllModels(t *testing.T) {
	models := GetAllModels()
	if len(models) != 3 {
		t.Errorf("Expected 3 models, got %d", len(models))
	}

	expectedModels := []string{YandexGPTLite, YandexGPT, AliceAI}
	for _, expected := range expectedModels {
		found := false
		for _, model := range models {
			if model == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected model %s not found in GetAllModels()", expected)
		}
	}
}

func TestGetModelDescriptions(t *testing.T) {
	descriptions := GetModelDescriptions()
	if len(descriptions) != 3 {
		t.Errorf("Expected 3 model descriptions, got %d", len(descriptions))
	}

	if _, ok := descriptions[YandexGPTLite]; !ok {
		t.Error("YandexGPTLite description not found")
	}
	if _, ok := descriptions[YandexGPT]; !ok {
		t.Error("YandexGPT description not found")
	}
	if _, ok := descriptions[AliceAI]; !ok {
		t.Error("AliceAI description not found")
	}
}

func TestGetModelURI(t *testing.T) {
	tests := []struct {
		name     string
		model    string
		folderID string
		expected string
	}{
		{
			name:     "YandexGPTLite URI",
			model:    YandexGPTLite,
			folderID: "test-folder",
			expected: "gpt://test-folder/yandexgpt-lite",
		},
		{
			name:     "YandexGPT URI",
			model:    YandexGPT,
			folderID: "test-folder",
			expected: "gpt://test-folder/yandexgpt",
		},
		{
			name:     "AliceAI URI",
			model:    AliceAI,
			folderID: "test-folder",
			expected: "gpt://test-folder/aliceai-llm",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetModelURI(tt.model, tt.folderID)
			if result != tt.expected {
				t.Errorf("GetModelURI(%s, %s) = %s, expected %s", tt.model, tt.folderID, result, tt.expected)
			}
		})
	}
}
