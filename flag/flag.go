package flag

import (
	"chatGPT/config"
	"flag"
	"github.com/gookit/color"
)

func Init() {
	flag.StringVar(&config.Prompt, "p", "", "输入要询问的问题")
	flag.Float64Var(&config.Temperature, "t", config.Temperature, "输入 Temperature，范围：0-2")
	flag.BoolVar(&config.Seal, "f", false, "按 F 解除封印")
	//flag.BoolVar(&config.Stream, "s", false, "进入一问一答模式")
	flag.StringVar(&config.ConfigPath, "l", config.ConfigPath, "加载配置文件（仅支持 yaml 格式文件），默认在用户根目录创建\"chatGPT.yaml\"文件")
	flag.BoolVar(&config.NotLogo, "n", false, "不显示logo")
	flag.StringVar(&config.PROXY, "proxy", config.PROXY, "设置代理，如：-proxy socks5://192.168.1.1:1080 或 -proxy http://192.168.1.1:8080")

	flag.Usage = func() {
		color.C256(46).Println(config.Logo)
		flag.PrintDefaults()
	}

	flag.Parse()
}
