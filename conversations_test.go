package yandexgpt

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupConversationsTestServer(t *testing.T, handler http.HandlerFunc) (*Client, *httptest.Server) {
	t.Helper()

	iamCalled := false
	mux := http.NewServeMux()

	mux.HandleFunc("/iam/v1/tokens", func(w http.ResponseWriter, r *http.Request) {
		iamCalled = true
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"iamToken": "test_iam_token"})
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !iamCalled && r.URL.Path != "/iam/v1/tokens" {
			// First call should be IAM token
			iamCalled = true
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"iamToken": "test_iam_token"})
			return
		}
		handler(w, r)
	})

	server := httptest.NewServer(mux)

	client, err := NewClientWithHTTPClient("test_oauth_token", "test_folder_id", server.Client())
	if err != nil {
		t.Fatal(err)
	}

	// Pre-set IAM token to avoid IAM call complexity
	client.iamToken = "test_iam_token"
	client.tokenExpiry = client.tokenExpiry.Add(1<<63 - 1) // far future

	return client, server
}

func TestConversationsCreate(t *testing.T) {
	client, server := setupConversationsTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if r.Header.Get("Authorization") != "Bearer test_iam_token" {
			t.Errorf("Expected Bearer test_iam_token, got %s", r.Header.Get("Authorization"))
		}

		var body map[string]interface{}
		json.NewDecoder(r.Body).Decode(&body)

		if meta, ok := body["metadata"].(map[string]interface{}); !ok || meta["title"] != "Test" {
			t.Errorf("Expected metadata with title=Test, got %v", body["metadata"])
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Conversation{
			ID:        "conv_123",
			Object:    "conversation",
			Metadata:  map[string]string{"title": "Test"},
			CreatedAt: 1700000000,
		})
	})
	defer server.Close()

	// Override the base URL
	origURL := conversationsBaseURL
	defer func() { setConversationsBaseURL(origURL) }()
	setConversationsBaseURL(server.URL)

	conv, err := client.Conversations().Create(map[string]string{"title": "Test"}, nil)
	if err != nil {
		t.Fatal(err)
	}

	if conv.ID != "conv_123" {
		t.Errorf("Expected conv_123, got %s", conv.ID)
	}
	if conv.Object != "conversation" {
		t.Errorf("Expected conversation, got %s", conv.Object)
	}
}

func TestConversationsGet(t *testing.T) {
	client, server := setupConversationsTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Conversation{
			ID:        "conv_123",
			Object:    "conversation",
			CreatedAt: 1700000000,
		})
	})
	defer server.Close()

	origURL := conversationsBaseURL
	defer func() { setConversationsBaseURL(origURL) }()
	setConversationsBaseURL(server.URL)

	conv, err := client.Conversations().Get("conv_123")
	if err != nil {
		t.Fatal(err)
	}

	if conv.ID != "conv_123" {
		t.Errorf("Expected conv_123, got %s", conv.ID)
	}
}

func TestConversationsUpdate(t *testing.T) {
	client, server := setupConversationsTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Conversation{
			ID:        "conv_123",
			Object:    "conversation",
			Metadata:  map[string]string{"title": "Updated"},
			CreatedAt: 1700000000,
		})
	})
	defer server.Close()

	origURL := conversationsBaseURL
	defer func() { setConversationsBaseURL(origURL) }()
	setConversationsBaseURL(server.URL)

	conv, err := client.Conversations().Update("conv_123", map[string]string{"title": "Updated"})
	if err != nil {
		t.Fatal(err)
	}

	if conv.Metadata["title"] != "Updated" {
		t.Errorf("Expected Updated, got %s", conv.Metadata["title"])
	}
}

func TestConversationsDelete(t *testing.T) {
	client, server := setupConversationsTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ConversationDeleted{
			Object:  "conversation.deleted",
			Deleted: true,
			ID:      "conv_123",
		})
	})
	defer server.Close()

	origURL := conversationsBaseURL
	defer func() { setConversationsBaseURL(origURL) }()
	setConversationsBaseURL(server.URL)

	result, err := client.Conversations().Delete("conv_123")
	if err != nil {
		t.Fatal(err)
	}

	if !result.Deleted {
		t.Error("Expected deleted to be true")
	}
	if result.Object != "conversation.deleted" {
		t.Errorf("Expected conversation.deleted, got %s", result.Object)
	}
}

func TestConversationsCreateItems(t *testing.T) {
	client, server := setupConversationsTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ConversationItemsList{
			Object: "list",
			Data: []ConversationItem{
				{Type: "message", ID: "item_1", Role: "user", Status: "completed"},
			},
			HasMore: false,
			FirstID: "item_1",
			LastID:  "item_1",
		})
	})
	defer server.Close()

	origURL := conversationsBaseURL
	defer func() { setConversationsBaseURL(origURL) }()
	setConversationsBaseURL(server.URL)

	items := []ConversationItem{
		{
			Type: "message",
			Role: "user",
			Content: []ConversationContentPart{
				{Type: "input_text", Text: "Hello"},
			},
		},
	}

	result, err := client.Conversations().CreateItems("conv_123", items)
	if err != nil {
		t.Fatal(err)
	}

	if result.Object != "list" {
		t.Errorf("Expected list, got %s", result.Object)
	}
	if len(result.Data) != 1 {
		t.Errorf("Expected 1 item, got %d", len(result.Data))
	}
}

func TestConversationsListItems(t *testing.T) {
	client, server := setupConversationsTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}

		query := r.URL.Query()
		if query.Get("limit") != "10" {
			t.Errorf("Expected limit=10, got %s", query.Get("limit"))
		}
		if query.Get("order") != "asc" {
			t.Errorf("Expected order=asc, got %s", query.Get("order"))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ConversationItemsList{
			Object: "list",
			Data: []ConversationItem{
				{Type: "message", ID: "item_1", Role: "user"},
				{Type: "message", ID: "item_2", Role: "assistant"},
			},
			HasMore: false,
			FirstID: "item_1",
			LastID:  "item_2",
		})
	})
	defer server.Close()

	origURL := conversationsBaseURL
	defer func() { setConversationsBaseURL(origURL) }()
	setConversationsBaseURL(server.URL)

	limit := 10
	order := "asc"
	result, err := client.Conversations().ListItems("conv_123", &ListItemsOptions{
		Limit: &limit,
		Order: &order,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Data) != 2 {
		t.Errorf("Expected 2 items, got %d", len(result.Data))
	}
}

func TestConversationsGetItem(t *testing.T) {
	client, server := setupConversationsTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ConversationItem{
			Type:   "message",
			ID:     "item_1",
			Role:   "user",
			Status: "completed",
			Content: []ConversationContentPart{
				{Type: "input_text", Text: "Hello"},
			},
		})
	})
	defer server.Close()

	origURL := conversationsBaseURL
	defer func() { setConversationsBaseURL(origURL) }()
	setConversationsBaseURL(server.URL)

	item, err := client.Conversations().GetItem("conv_123", "item_1")
	if err != nil {
		t.Fatal(err)
	}

	if item.ID != "item_1" {
		t.Errorf("Expected item_1, got %s", item.ID)
	}
	if item.Type != "message" {
		t.Errorf("Expected message, got %s", item.Type)
	}
}

func TestConversationsDeleteItem(t *testing.T) {
	client, server := setupConversationsTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Conversation{
			ID:        "conv_123",
			Object:    "conversation",
			CreatedAt: 1700000000,
		})
	})
	defer server.Close()

	origURL := conversationsBaseURL
	defer func() { setConversationsBaseURL(origURL) }()
	setConversationsBaseURL(server.URL)

	conv, err := client.Conversations().DeleteItem("conv_123", "item_1")
	if err != nil {
		t.Fatal(err)
	}

	if conv.ID != "conv_123" {
		t.Errorf("Expected conv_123, got %s", conv.ID)
	}
}

func TestConversationsClientIsSingleton(t *testing.T) {
	client, _ := NewClient("test_token", "test_folder")

	c1 := client.Conversations()
	c2 := client.Conversations()

	if c1 != c2 {
		t.Error("Expected same ConversationsClient instance")
	}
}

func TestConversationsAPIError(t *testing.T) {
	client, server := setupConversationsTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Bad Request: invalid conversation_id",
		})
	})
	defer server.Close()

	origURL := conversationsBaseURL
	defer func() { setConversationsBaseURL(origURL) }()
	setConversationsBaseURL(server.URL)

	_, err := client.Conversations().Get("invalid_id")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("Expected APIError, got %T", err)
	}
	if apiErr.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", apiErr.StatusCode)
	}
}
