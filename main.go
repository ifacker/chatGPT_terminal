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
	control.Init()
}
