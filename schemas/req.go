package schemas

type AiRequest struct {
	Model    string      `json:"model"`
	Stream   bool        `json:"stream"`
	Messages []AiMessage `json:"messages"`
}

type AiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Image   []byte `json:"images,omitempty"`
}
