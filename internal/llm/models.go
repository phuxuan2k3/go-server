package llm

type LLMRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type LLMResponse struct {
	Model       string `json:"model"`
	Created_at  string `json:"created_at"`
	Response    string `json:"response"`
	Done_reason string `json:"done_reason"`
	Done        bool   `json:"done"`
}
