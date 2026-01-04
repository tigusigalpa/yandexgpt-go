package models

import "testing"

func TestGetARTModelURI(t *testing.T) {
	tests := []struct {
		name      string
		catalogID string
		expected  string
	}{
		{
			name:      "Standard catalog ID",
			catalogID: "test-catalog",
			expected:  "art://test-catalog/yandex-art/latest",
		},
		{
			name:      "Different catalog ID",
			catalogID: "another-catalog",
			expected:  "art://another-catalog/yandex-art/latest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetARTModelURI(tt.catalogID)
			if result != tt.expected {
				t.Errorf("GetARTModelURI(%s) = %s, expected %s", tt.catalogID, result, tt.expected)
			}
		})
	}
}
