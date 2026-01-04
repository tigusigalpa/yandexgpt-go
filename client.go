package yandexgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tigusigalpa/yandexgpt-go/models"
)

const (
	CompletionEndpoint           = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"
	ImageGenerationAsyncEndpoint = "https://llm.api.cloud.yandex.net/foundationModels/v1/imageGenerationAsync"
	OperationsEndpoint           = "https://operation.api.cloud.yandex.net/operations"
)

type Client struct {
	httpClient  *http.Client
	oauthToken  string
	folderID    string
	iamToken    string
	tokenExpiry time.Time
}

func NewClient(oauthToken, folderID string) (*Client, error) {
	if oauthToken == "" {
		return nil, NewAuthenticationError("OAuth token cannot be empty", nil)
	}
	if folderID == "" {
		return nil, NewAuthenticationError("Folder ID cannot be empty", nil)
	}

	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		oauthToken: oauthToken,
		folderID:   folderID,
	}, nil
}

func NewClientWithHTTPClient(oauthToken, folderID string, httpClient *http.Client) (*Client, error) {
	if oauthToken == "" {
		return nil, NewAuthenticationError("OAuth token cannot be empty", nil)
	}
	if folderID == "" {
		return nil, NewAuthenticationError("Folder ID cannot be empty", nil)
	}

	return &Client{
		httpClient: httpClient,
		oauthToken: oauthToken,
		folderID:   folderID,
	}, nil
}

func (c *Client) getValidIAMToken() (string, error) {
	if c.iamToken != "" && time.Now().Before(c.tokenExpiry) {
		return c.iamToken, nil
	}

	type iamRequest struct {
		YandexPassportOauthToken string `json:"yandexPassportOauthToken"`
	}

	type iamResponse struct {
		IAMToken  string    `json:"iamToken"`
		ExpiresAt time.Time `json:"expiresAt"`
	}

	reqBody, err := json.Marshal(iamRequest{
		YandexPassportOauthToken: c.oauthToken,
	})
	if err != nil {
		return "", NewAuthenticationError("failed to marshal IAM request", err)
	}

	req, err := http.NewRequest("POST", "https://iam.api.cloud.yandex.net/iam/v1/tokens", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", NewAuthenticationError("failed to create IAM request", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", NewAuthenticationError("failed to get IAM token", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", NewAuthenticationError(fmt.Sprintf("IAM token request failed with status %d: %s", resp.StatusCode, string(body)), nil)
	}

	var iamResp iamResponse
	if err := json.NewDecoder(resp.Body).Decode(&iamResp); err != nil {
		return "", NewAuthenticationError("failed to decode IAM response", err)
	}

	c.iamToken = iamResp.IAMToken
	c.tokenExpiry = iamResp.ExpiresAt.Add(-5 * time.Minute)

	return c.iamToken, nil
}

func (c *Client) GenerateText(prompt, model string, options *CompletionOptions) (*CompletionResponse, error) {
	if !models.IsValidModel(model) {
		return nil, NewAPIError(fmt.Sprintf("invalid model: %s", model), 0, nil)
	}

	iamToken, err := c.getValidIAMToken()
	if err != nil {
		return nil, err
	}

	modelURI := models.GetModelURI(model, c.folderID)

	if options == nil {
		options = &CompletionOptions{
			Stream:      false,
			Temperature: 0.6,
			MaxTokens:   2000,
		}
	}

	request := CompletionRequest{
		ModelURI:          modelURI,
		CompletionOptions: *options,
		Messages: []Message{
			{
				Role: "user",
				Text: prompt,
			},
		},
	}

	return c.sendCompletionRequest(iamToken, request)
}

func (c *Client) GenerateFromMessages(messages []Message, model string, options *CompletionOptions) (*CompletionResponse, error) {
	if !models.IsValidModel(model) {
		return nil, NewAPIError(fmt.Sprintf("invalid model: %s", model), 0, nil)
	}

	iamToken, err := c.getValidIAMToken()
	if err != nil {
		return nil, err
	}

	modelURI := models.GetModelURI(model, c.folderID)

	if options == nil {
		options = &CompletionOptions{
			Stream:      false,
			Temperature: 0.6,
			MaxTokens:   2000,
		}
	}

	request := CompletionRequest{
		ModelURI:          modelURI,
		CompletionOptions: *options,
		Messages:          messages,
	}

	return c.sendCompletionRequest(iamToken, request)
}

func (c *Client) sendCompletionRequest(iamToken string, request CompletionRequest) (*CompletionResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, NewAPIError("failed to marshal request", 0, err)
	}

	req, err := http.NewRequest("POST", CompletionEndpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, NewAPIError("failed to create request", 0, err)
	}

	req.Header.Set("Authorization", "Bearer "+iamToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, NewAPIError("failed to send request", 0, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, NewAPIError("failed to read response", resp.StatusCode, err)
	}

	if resp.StatusCode != http.StatusOK {
		var errorResp map[string]interface{}
		if err := json.Unmarshal(body, &errorResp); err == nil {
			if msg, ok := errorResp["message"].(string); ok {
				return nil, NewAPIError(msg, resp.StatusCode, nil)
			}
		}
		return nil, NewAPIError(fmt.Sprintf("API request failed: %s", string(body)), resp.StatusCode, nil)
	}

	var response CompletionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, NewAPIError("failed to decode response", resp.StatusCode, err)
	}

	return &response, nil
}

func (c *Client) GenerateImageAsync(messages interface{}, options *GenerationOptions, catalogID *string) (*Operation, error) {
	iamToken, err := c.getValidIAMToken()
	if err != nil {
		return nil, err
	}

	folderID := c.folderID
	if catalogID != nil {
		folderID = *catalogID
	}

	modelURI := models.GetARTModelURI(folderID)

	if options == nil {
		options = &GenerationOptions{}
	}

	artMessages := normalizeARTMessages(messages)

	request := ImageGenerationRequest{
		ModelURI:          modelURI,
		GenerationOptions: *options,
		Messages:          artMessages,
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, NewAPIError("failed to marshal request", 0, err)
	}

	req, err := http.NewRequest("POST", ImageGenerationAsyncEndpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, NewAPIError("failed to create request", 0, err)
	}

	req.Header.Set("Authorization", "Bearer "+iamToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, NewAPIError("failed to send request", 0, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, NewAPIError("failed to read response", resp.StatusCode, err)
	}

	if resp.StatusCode != http.StatusOK {
		var errorResp map[string]interface{}
		if err := json.Unmarshal(body, &errorResp); err == nil {
			if msg, ok := errorResp["message"].(string); ok {
				return nil, NewAPIError(msg, resp.StatusCode, nil)
			}
		}
		return nil, NewAPIError(fmt.Sprintf("API request failed: %s", string(body)), resp.StatusCode, nil)
	}

	var operation Operation
	if err := json.Unmarshal(body, &operation); err != nil {
		return nil, NewAPIError("failed to decode response", resp.StatusCode, err)
	}

	return &operation, nil
}

func (c *Client) GetOperation(operationID string) (*Operation, error) {
	iamToken, err := c.getValidIAMToken()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", OperationsEndpoint, operationID), nil)
	if err != nil {
		return nil, NewAPIError("failed to create request", 0, err)
	}

	req.Header.Set("Authorization", "Bearer "+iamToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, NewAPIError("failed to send request", 0, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, NewAPIError("failed to read response", resp.StatusCode, err)
	}

	if resp.StatusCode != http.StatusOK {
		var errorResp map[string]interface{}
		if err := json.Unmarshal(body, &errorResp); err == nil {
			if msg, ok := errorResp["message"].(string); ok {
				return nil, NewAPIError(msg, resp.StatusCode, nil)
			}
		}
		return nil, NewAPIError(fmt.Sprintf("API request failed: %s", string(body)), resp.StatusCode, nil)
	}

	var operation Operation
	if err := json.Unmarshal(body, &operation); err != nil {
		return nil, NewAPIError("failed to decode response", resp.StatusCode, err)
	}

	return &operation, nil
}

func (c *Client) GenerateImage(messages interface{}, options *GenerationOptions, catalogID *string) (*ImageGenerationResult, error) {
	operation, err := c.GenerateImageAsync(messages, options, catalogID)
	if err != nil {
		return nil, err
	}

	if operation.ID == "" {
		return nil, NewAPIError("operation ID not found in response", 0, nil)
	}

	pollInterval := 2 * time.Second
	maxWait := 10 * time.Minute
	elapsed := time.Duration(0)

	for elapsed < maxWait {
		time.Sleep(pollInterval)
		elapsed += pollInterval

		op, err := c.GetOperation(operation.ID)
		if err != nil {
			return nil, err
		}

		if op.Done {
			if op.Error != nil {
				return nil, NewAPIError(fmt.Sprintf("operation error: %s", op.Error.Message), op.Error.Code, nil)
			}

			if op.Response == nil || op.Response.Image == "" {
				return nil, NewAPIError("image data not found in operation response", 0, nil)
			}

			return &ImageGenerationResult{
				OperationID: operation.ID,
				ImageBase64: op.Response.Image,
			}, nil
		}
	}

	return nil, NewAPIError("operation timed out", 0, nil)
}

func (c *Client) GetAvailableModels() []string {
	return models.GetAllModels()
}

func (c *Client) GetModelDescriptions() map[string]string {
	return models.GetModelDescriptions()
}

func (c *Client) GetFolderID() string {
	return c.folderID
}

func (c *Client) SetFolderID(folderID string) {
	c.folderID = folderID
}

func normalizeARTMessages(messages interface{}) []ARTMessage {
	switch v := messages.(type) {
	case string:
		return []ARTMessage{{Text: v, Weight: 1}}
	case []string:
		result := make([]ARTMessage, len(v))
		for i, msg := range v {
			result[i] = ARTMessage{Text: msg, Weight: 1}
		}
		return result
	case []ARTMessage:
		return v
	case []map[string]interface{}:
		result := make([]ARTMessage, len(v))
		for i, msg := range v {
			text, _ := msg["text"].(string)
			weight := 1
			if w, ok := msg["weight"].(int); ok {
				weight = w
			} else if w, ok := msg["weight"].(float64); ok {
				weight = int(w)
			}
			result[i] = ARTMessage{Text: text, Weight: weight}
		}
		return result
	default:
		return []ARTMessage{{Text: fmt.Sprintf("%v", messages), Weight: 1}}
	}
}
