package schemas

import "time"

type AiRequest struct {
	Model    string      `json:"model"`
	Stream   bool        `json:"stream"`
	Messages []AiMessage `json:"messages"`
}

type AiResponse struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Message            AiMessage `json:"message"`
	DoneReason         string    `json:"done_reason"`
	Done               bool      `json:"done"`
	TotalDuration      int64     `json:"total_duration"`
	LoadDuration       int64     `json:"load_duration"`
	PromptEvalCount    int       `json:"prompt_eval_count"`
	PromptEvalDuration int64     `json:"prompt_eval_duration"`
	EvalCount          int       `json:"eval_count"`
	EvalDuration       int64     `json:"eval_duration"`
}
type AiMessage struct {
	Role    string   `json:"role"`
	Content string   `json:"content"`
	Image   []string `json:"images,omitempty"`
}
