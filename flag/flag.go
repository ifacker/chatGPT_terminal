package flag

import (
	"chatGPT/config"
	"flag"
)

func Init() {
	flag.StringVar(&config.Prompt, "p", "", "输入要询问的问题")
	flag.Float64Var(&config.Temperature, "t", config.Temperature, "输入 Temperature，范围：0-2")
	flag.BoolVar(&config.Seal, "f", false, "按 F 解除封印")
	//flag.BoolVar(&config.Stream, "s", false, "进入一问一答模式")
	flag.StringVar(&config.ConfigPath, "l", config.ConfigPath, "加载配置文件（仅支持 yaml 格式文件），默认在用户根目录创建\"chatGPT.yaml\"文件")

	flag.Parse()
}
