package generation

type Request struct {
	Prompt                   string   `json:"prompt"`
	MaxTokens                int64    `json:"max_tokens"`
	Temperature              *float64 `json:"temperature"`
	TopProbability           *float64 `json:"top_p"`
	DesiredNumberOfResponses *float64 `json:"n"`
}

type Response struct {
	Id          string       `json:"id"`
	Generations []Generation `json:"generations"`
	Usage       Usage        `json:"usage"`
}

type Generation struct {
	Text   string `json:"text"`
	Tokens int    `json:"tokens"`
}

type Usage struct {
	PromptTokens    int `json:"prompt_tokens"`
	GeneratedTokens int `json:"generated_tokens"`
	TotalTokens     int `json:"total_tokens"`
}
