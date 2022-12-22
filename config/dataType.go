package config

type ReqJson struct {
	// 要使用的模型的 ID。可以使用列表模型API 查看所有可用模型，或参阅模型概述了解它们的描述。
	Model string `json:"model"`
	// 用于生成完成、编码为字符串、字符串数组、标记数组或标记数组数组的提示。
	// 请注意，<|endoftext|> 是模型在训练期间看到的文档分隔符，因此如果未指定提示，模型将生成，就像从新文档的开头一样。
	Prompt string `json:"prompt"`
	// 完成时要生成的最大令牌数。
	// 提示加号的令牌计数不能超过模型的上下文长度。大多数模型的上下文长度为 2048 个令牌（最新模型除外，它支持 4096）。
	Max_tokens int `json:"max_tokens"`
	// 使用什么采样温度。值越高意味着模型将承担更多风险。对于更具创造性的应用程序，请尝试 0.9，对于具有明确定义答案的应用程序，请尝试 0（argmax 采样）。
	// 我们通常建议更改此设置，但不要同时更改两者。
	Temperature float64 `json:"temperature"`
	// 使用温度采样的替代方法称为核心采样，其中模型考虑具有top_p概率质量的令牌的结果。因此，0.1 意味着只考虑包含前 10% 概率质量的代币。
	// 我们通常建议更改此设置，但不要同时更改两者。 top_p
	Top_p int `json:"top_p"`
	// 为每个提示生成的完成次数。
	// 注意：由于此参数会生成许多完成，因此它会快速消耗令牌配额。小心使用并确保您有合理的设置。 temperature
	N int `json:"n"`
	// 是否流式传输回部分进度。如果设置，令牌将在可用时作为纯数据服务器发送的事件发送，流由消息终止。data: [DONE]
	Stream bool `json:"stream"`
	// 包括最可能的令牌的日志概率，以及所选令牌。例如，ifis 5，API 将返回 5 个最有可能的令牌的列表。API 将始终返回采样令牌，因此响应中最多可能存在元素。logprobslogprobslogproblogprobs+1
	// 最大值为 5。如果您需要更多，请通过我们的帮助中心与我们联系并描述您的使用案例。logprobs
	//Logprobs any `json:"logprobs"`
	// 最多 4 个序列，其中 API 将停止生成更多令牌。返回的文本将不包含停止序列。
	Stop string `json:"stop"`
}

type RespJson struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// 配置文件
type YamlConfig struct {
	ApiToken string `yaml:"apiToken"`
	Org      string `yaml:"orgID"`
}
