package config

import "net/http"

var (
	API_URL   = "https://api.openai.com/v1/completions"
	API_TOKEN string
	ORG_ID    string
	PROXY     string
)

var (
	Prompt      string
	Temperature float64 = 0.9
	// 解除封印
	Seal bool
	// 选择是否使用一问一答连续对话
	//Stream bool

	// 不显示 logo
	NotLogo bool

	Tr *http.Transport
)

var (
	Version    = "1.3"
	ConfigPath = ""
)
