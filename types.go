package yandexgpt

type Message struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

type ReasoningOptions struct {
	Mode   string  `json:"mode,omitempty"`
	Effort *string `json:"effort,omitempty"`
}

type CompletionOptions struct {
	Stream           bool               `json:"stream"`
	Temperature      float64            `json:"temperature"`
	MaxTokens        int                `json:"maxTokens"`
	ReasoningOptions *ReasoningOptions  `json:"reasoningOptions,omitempty"`
}

type CompletionRequest struct {
	ModelURI          string            `json:"modelUri"`
	CompletionOptions CompletionOptions `json:"completionOptions"`
	Messages          []Message         `json:"messages"`
}

type Alternative struct {
	Message Message `json:"message"`
	Status  string  `json:"status"`
}

type Usage struct {
	InputTextTokens  int `json:"inputTextTokens"`
	CompletionTokens int `json:"completionTokens"`
	TotalTokens      int `json:"totalTokens"`
	ReasoningTokens  int `json:"reasoningTokens,omitempty"`
}

type Result struct {
	Alternatives []Alternative `json:"alternatives"`
	Usage        Usage         `json:"usage"`
	ModelVersion string        `json:"modelVersion"`
}

type CompletionResponse struct {
	Result Result `json:"result"`
}

type GenerationOptions struct {
	Seed        *int    `json:"seed,omitempty"`
	AspectRatio *string `json:"aspectRatio,omitempty"`
	MimeType    *string `json:"mimeType,omitempty"`
}

type ARTMessage struct {
	Text   string `json:"text"`
	Weight int    `json:"weight"`
}

type ImageGenerationRequest struct {
	ModelURI          string            `json:"modelUri"`
	GenerationOptions GenerationOptions `json:"generationOptions"`
	Messages          []ARTMessage      `json:"messages"`
}

type OperationMetadata struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	CreatedBy   string `json:"createdBy"`
	ModifiedAt  string `json:"modifiedAt"`
}

type OperationError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ImageResponse struct {
	Image string `json:"image"`
}

type Operation struct {
	ID          string             `json:"id"`
	Description string             `json:"description"`
	CreatedAt   string             `json:"createdAt"`
	CreatedBy   string             `json:"createdBy"`
	ModifiedAt  string             `json:"modifiedAt"`
	Done        bool               `json:"done"`
	Metadata    *OperationMetadata `json:"metadata,omitempty"`
	Error       *OperationError    `json:"error,omitempty"`
	Response    *ImageResponse     `json:"response,omitempty"`
}

type ImageGenerationResult struct {
	OperationID string
	ImageBase64 string
}
