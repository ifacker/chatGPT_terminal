package config

var (
	API_URL   = "https://api.openai.com/v1/completions"
	API_TOKEN string
)

var (
	Prompt      string
	Temperature float64 = 0.9
	// 解除封印
	Seal = false
	// 选择是否使用一问一答连续对话
	Stream = false
)

var (
	Version    = "1.0"
	ConfigPath = ""
)
