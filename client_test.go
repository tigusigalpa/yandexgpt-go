package yandexgpt

import (
	"testing"

	"github.com/tigusigalpa/yandexgpt-go/models"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name        string
		oauthToken  string
		folderID    string
		expectError bool
	}{
		{
			name:        "Valid credentials",
			oauthToken:  "test_token",
			folderID:    "test_folder",
			expectError: false,
		},
		{
			name:        "Empty OAuth token",
			oauthToken:  "",
			folderID:    "test_folder",
			expectError: true,
		},
		{
			name:        "Empty folder ID",
			oauthToken:  "test_token",
			folderID:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.oauthToken, tt.folderID)
			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if client == nil {
					t.Error("Expected client but got nil")
				}
			}
		})
	}
}

func TestGetAvailableModels(t *testing.T) {
	client, err := NewClient("test_token", "test_folder")
	if err != nil {
		t.Fatal(err)
	}

	availableModels := client.GetAvailableModels()
	if len(availableModels) == 0 {
		t.Error("Expected at least one model")
	}

	expectedModels := []string{
		models.YandexGPTLite,
		models.YandexGPT,
		models.AliceAI,
	}

	for _, expected := range expectedModels {
		found := false
		for _, model := range availableModels {
			if model == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected model %s not found", expected)
		}
	}
}

func TestGetModelDescriptions(t *testing.T) {
	client, err := NewClient("test_token", "test_folder")
	if err != nil {
		t.Fatal(err)
	}

	descriptions := client.GetModelDescriptions()
	if len(descriptions) == 0 {
		t.Error("Expected at least one model description")
	}

	if _, ok := descriptions[models.YandexGPTLite]; !ok {
		t.Error("Expected YandexGPTLite description")
	}
}

func TestGetSetFolderID(t *testing.T) {
	client, err := NewClient("test_token", "test_folder")
	if err != nil {
		t.Fatal(err)
	}

	if client.GetFolderID() != "test_folder" {
		t.Errorf("Expected folder ID 'test_folder', got '%s'", client.GetFolderID())
	}

	client.SetFolderID("new_folder")
	if client.GetFolderID() != "new_folder" {
		t.Errorf("Expected folder ID 'new_folder', got '%s'", client.GetFolderID())
	}
}

func TestNormalizeARTMessages(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int
	}{
		{
			name:     "String input",
			input:    "test message",
			expected: 1,
		},
		{
			name:     "String slice input",
			input:    []string{"msg1", "msg2"},
			expected: 2,
		},
		{
			name: "ARTMessage slice input",
			input: []ARTMessage{
				{Text: "msg1", Weight: 1},
				{Text: "msg2", Weight: 2},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeARTMessages(tt.input)
			if len(result) != tt.expected {
				t.Errorf("Expected %d messages, got %d", tt.expected, len(result))
			}
		})
	}
}
