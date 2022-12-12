package control

import (
	"chatGPT/config"
	"chatGPT/util/file"
	"github.com/gookit/color"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/user"
)

// 加载配置文件
func LoadConfig() {
	home, err := user.Current()
	if err != nil {
		log.Println(err)
	}
	tmpPath := home.HomeDir + "/chatGPT.yaml"
	if config.ConfigPath == "" {
		body, err := file.ReadFile(tmpPath)
		if err != nil {
			log.Println(err)
			if os.IsNotExist(err) {
				out, err := yaml.Marshal(config.YamlConfig{ApiToken: ""})
				if err != nil {
					log.Println(err)
				}
				ok, err := file.WriteFile(tmpPath, out)
				if err != nil {
					log.Println(err)
				}
				if ok {
					color.Infof("配置文件\"chatGPT.yaml\"已创建，目录在：%s, 请前去配置该文件的 apiToken\n", tmpPath)
					os.Exit(1)
				}
			}
		}
		// 对读取的内容进行反序列化
		var configYaml = &config.YamlConfig{}
		err = yaml.Unmarshal(body, configYaml)
		if err != nil {
			log.Println(err)
		}
		if configYaml.ApiToken == "" {
			color.Errorf("未检测到 apiToken，请重新配置～")
			os.Exit(1)
		}
		config.API_TOKEN = configYaml.ApiToken
	}

}
