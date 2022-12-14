package main

import (
	"chatGPT/config"
	"chatGPT/control"
	"chatGPT/flag"
	"github.com/gookit/color"
)

func init() {
	color.C256(46).Println(config.Logo)
	flag.Init()
	control.LoadConfig()
}

func main() {
	//inputReader := bufio.NewReader(os.Stdin)
	//fmt.Println("Please enter some input: ")
	//input, err := inputReader.ReadString('\n')
	//if err == nil {
	//	fmt.Printf("The input was: %s\n", input)
	//}
	control.Init()
}
