# chatGPT 问答机器人
## 使用方法：
```shell
# 获取帮助信息
./chatGPT -h

# 设置 Temperature （该选项可以让机器人的回答更接近人类，数值越高越接近，但是太高了会导致机器人胡言乱语，官方建议最高 0.9，本程序默认设置 0.9，有需要可自行设置）
./chatGPT -t 0.9 -p "提出的问题"

# -p 表示要提出的问题
./chatGPT -p "提出的问题"

# -f 解除机器人封印，机器人将会回答所有问题
./chatGPT -f -p "提出的问题"

# -s 一问一答模式，该模式程序不会结束，直至输入 quit 为止，注意⚠️ ：该模式仅可与 -t 同时使用，与 -f -p 同时使用均无效
./chatGPT -s 
```
## 版本信息
### v1.0
该版本暂时支持的功能仅有：
1. 支持手动设置 Temperature 数值
2. 支持一键解封机器人道德底线
3. 支持只回答但个问题
4. 支持一问一答形式  

该版本暂不支持但未来有可能上线的功能：
1. 暂时不支持机器人提问记忆功能，也就是说机器人暂时没有记忆能力，你上一句问的问题，它下一句就忘记了，暂时不知道需要什么参数才能实现，如有大佬知道，感谢🙏 告知，谢谢！