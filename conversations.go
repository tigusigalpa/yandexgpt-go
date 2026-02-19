package yandexgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var conversationsBaseURL = "https://ai.api.cloud.yandex.net/v1/conversations"

func setConversationsBaseURL(u string) {
	conversationsBaseURL = u
}

// ConversationsClient provides methods for interacting with the Conversations API.
type ConversationsClient struct {
	client *Client
}

// Create creates a new conversation with optional metadata and initial items.
//
// See https://yandex.cloud/ru/docs/ai-studio/conversations/createConversation
func (cc *ConversationsClient) Create(metadata map[string]string, items []ConversationItem) (*Conversation, error) {
	body := make(map[string]interface{})

	if metadata != nil {
		body["metadata"] = metadata
	}

	if items != nil {
		body["items"] = items
	}

	var conversation Conversation
	if err := cc.sendRequest("POST", conversationsBaseURL, body, &conversation); err != nil {
		return nil, err
	}

	return &conversation, nil
}

// Get retrieves a conversation by its ID.
//
// See https://yandex.cloud/ru/docs/ai-studio/conversations/getConversation
func (cc *ConversationsClient) Get(conversationID string) (*Conversation, error) {
	var conversation Conversation
	if err := cc.sendRequest("GET", fmt.Sprintf("%s/%s", conversationsBaseURL, conversationID), nil, &conversation); err != nil {
		return nil, err
	}

	return &conversation, nil
}

// Update updates the metadata of a conversation.
//
// See https://yandex.cloud/ru/docs/ai-studio/conversations/updateConversation
func (cc *ConversationsClient) Update(conversationID string, metadata map[string]string) (*Conversation, error) {
	body := make(map[string]interface{})

	if metadata != nil {
		body["metadata"] = metadata
	}

	var conversation Conversation
	if err := cc.sendRequest("POST", fmt.Sprintf("%s/%s", conversationsBaseURL, conversationID), body, &conversation); err != nil {
		return nil, err
	}

	return &conversation, nil
}

// Delete deletes a conversation by its ID.
//
// See https://yandex.cloud/ru/docs/ai-studio/conversations/deleteConversation
func (cc *ConversationsClient) Delete(conversationID string) (*ConversationDeleted, error) {
	var result ConversationDeleted
	if err := cc.sendRequest("DELETE", fmt.Sprintf("%s/%s", conversationsBaseURL, conversationID), nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateItems adds items to a conversation.
//
// See https://yandex.cloud/ru/docs/ai-studio/conversations/createConversationItems
func (cc *ConversationsClient) CreateItems(conversationID string, items []ConversationItem) (*ConversationItemsList, error) {
	body := map[string]interface{}{
		"items": items,
	}

	var result ConversationItemsList
	if err := cc.sendRequest("POST", fmt.Sprintf("%s/%s/items", conversationsBaseURL, conversationID), body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ListItems retrieves items from a conversation with optional pagination parameters.
//
// See https://yandex.cloud/ru/docs/ai-studio/conversations/listConversationItems
func (cc *ConversationsClient) ListItems(conversationID string, opts *ListItemsOptions) (*ConversationItemsList, error) {
	u := fmt.Sprintf("%s/%s/items", conversationsBaseURL, conversationID)

	if opts != nil {
		params := url.Values{}
		if opts.Limit != nil {
			params.Set("limit", strconv.Itoa(*opts.Limit))
		}
		if opts.Order != nil {
			params.Set("order", *opts.Order)
		}
		if opts.After != nil {
			params.Set("after", *opts.After)
		}
		if encoded := params.Encode(); encoded != "" {
			u += "?" + encoded
		}
	}

	var result ConversationItemsList
	if err := cc.sendRequest("GET", u, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetItem retrieves a single item from a conversation.
//
// See https://yandex.cloud/ru/docs/ai-studio/conversations/getConversationItem
func (cc *ConversationsClient) GetItem(conversationID, itemID string) (*ConversationItem, error) {
	var item ConversationItem
	if err := cc.sendRequest("GET", fmt.Sprintf("%s/%s/items/%s", conversationsBaseURL, conversationID, itemID), nil, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

// DeleteItem deletes an item from a conversation.
//
// See https://yandex.cloud/ru/docs/ai-studio/conversations/deleteConversationItem
func (cc *ConversationsClient) DeleteItem(conversationID, itemID string) (*Conversation, error) {
	var conversation Conversation
	if err := cc.sendRequest("DELETE", fmt.Sprintf("%s/%s/items/%s", conversationsBaseURL, conversationID, itemID), nil, &conversation); err != nil {
		return nil, err
	}

	return &conversation, nil
}

func (cc *ConversationsClient) sendRequest(method, requestURL string, body interface{}, result interface{}) error {
	iamToken, err := cc.client.getValidIAMToken()
	if err != nil {
		return err
	}

	var reqBody io.Reader
	if body != nil && (method == "POST" || method == "PUT" || method == "PATCH") {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return NewAPIError("failed to marshal request", 0, err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, requestURL, reqBody)
	if err != nil {
		return NewAPIError("failed to create request", 0, err)
	}

	req.Header.Set("Authorization", "Bearer "+iamToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := cc.client.httpClient.Do(req)
	if err != nil {
		return NewAPIError("failed to send request", 0, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return NewAPIError("failed to read response", resp.StatusCode, err)
	}

	if resp.StatusCode != http.StatusOK {
		var errorResp map[string]interface{}
		if err := json.Unmarshal(respBody, &errorResp); err == nil {
			if msg, ok := errorResp["message"].(string); ok {
				return NewAPIError(msg, resp.StatusCode, nil)
			}
		}
		return NewAPIError(fmt.Sprintf("Conversations API request failed: %s", string(respBody)), resp.StatusCode, nil)
	}

	if err := json.Unmarshal(respBody, result); err != nil {
		return NewAPIError("failed to decode response", resp.StatusCode, err)
	}

	return nil
}
