# chatGPT 问答机器人
## 使用方法：
```shell
# 获取帮助信息
./chatGPT -h

# 进入一问一答模式（也可以上 -t 一并使用）
./chatGPT

# 设置 Temperature （该选项可以让机器人的回答更接近人类，数值越高越接近，但是太高了会导致机器人胡言乱语，官方建议最高 0.9，本程序默认设置 0.9，有需要可自行设置）
./chatGPT -t 0.9 -p "提出的问题"

# -p 表示要提出的问题
./chatGPT -p "提出的问题"

# -f 解除机器人封印，机器人将会回答所有问题
./chatGPT -f -p "提出的问题"

# -n 不显示 logo 使用
./chatGPT -n -p "提出的问题"

# -proxy 使用代理模式
./chatGPT -proxy "http://localhost:1080" -p "提出的问题"
```
## 演示：
<img src="img/gifmk3015.gif" width="530"/>
<img src="img/1.png" width="530"/>
<img src="img/2.png" width="530"/>

## 版本信息

### v1.3 🐹
1. 新增 -n 功能，关闭显示 logo
2. 新增代理功能（由于一些不可抗力...），新增代理功能可由两种方式实现，分别是：
   1. 参数形式，如：-proxy socks5://192.168.1.1:1080 或 -proxy http://192.168.1.1:8080
   2. 配置形式，如：在 ~/chatGPT.yaml 中进行配置，格式如下（或者可以把原来的配置文件删了，然后重新生成一个，重新配置）：
   3. 参数形式 的优先级比 配置形式 的优先级要高，所以如果 参数 和 配置 都有代理，那么代理将会走 参数 上的代理
```text
apiToken: "sk-xxxxxxx"
orgID: "org-xxxxxxx"
proxy: "http://localhost:7890"
```

### v1.2.1 🐭
#### 修复了API有时无法使用的情况，以下为操作步骤
1. 更新最新版 v1.2.1 版本的 chatGPT
2. 删除原来的 chatGPT.yaml 配置文件（在用户根目录下）
3. 重新配置 API key ，还有本次新增加的 Org-ID，之后即可解决发送请求无响应的问题

### v1.2 🐭
#### 新增加的功能：
1. 支持显示程序体温（Temperature）
2. 一问一答显示更美观（仅 linux、mac 可见，windows 无色，但是格式相比之前的版本显示的更加友好🧑‍🤝‍🧑 ）
3. 支持一问一答模式设置 -t 参数
#### 已知的问题：
1. 在 linux 和 mac 的终端上，在一问一答模式中，输入汉子，一旦输入错误，使用退格键删除文字，就会发现只删除了一半！并没有完全删干净，推测是 utf8 格式的编码占2个字节，但是退格一次只能删除1个字节，所以会出现半个字的情况，但是 Windows 上没有这个问题。我是用的是 "bufio.NewReader(os.Stdin)"，目前存在这个问题，我也不知道该怎么解决，如果有大佬知道，还请大佬赐教🙏 ～

### v1.1 🐱
#### 该版本修复了之前一个版本的 bug：
1. 修复了 -s 参数无法使用一问一答的 bug
#### 使用方式变更：
1. 去除了 -s 参数，改为直接运行 "chatGPT" 即可实现一问一答的模式
#### 新增加的功能：
1. 现在一问一答已支持记忆模式，AI 再也不用担心失忆啦～ 只要你不 quit，AI 就不会失忆！  

[在这里感谢 "P-TNPC" 师傅的不吝赐教，使我们的 AI 从此不会失忆 🌹 ！]  
[再次感谢 🎉 ！]

### v1.0 🐶
#### 该版本暂时支持的功能仅有：
1. 支持手动设置 Temperature 数值
2. 支持一键解封机器人道德底线
3. 支持只回答但个问题
4. 支持一问一答形式  

#### ~~该版本暂不支持但未来有可能上线的功能：~~
1. ~~暂时不支持机器人提问记忆功能，也就是说机器人暂时没有记忆能力，你上一句问的问题，它下一句就忘记了，暂时不知道需要什么参数才能实现，如有大佬知道，感谢🙏 告知，谢谢！~~