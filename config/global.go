package config

var (
	API_URL   = "https://api.openai.com/v1/completions"
	API_TOKEN string
)

var (
	Prompt      string
	Temperature float64 = 0.9
	// 解除封印
	Seal bool
	// 选择是否使用一问一答连续对话
	//Stream bool
)

var (
	Version    = "1.2"
	ConfigPath = ""
)
