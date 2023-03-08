package main

import (
	"chatGPT/config"
	"chatGPT/control"
	"chatGPT/flag"
	"crypto/tls"
	"github.com/gookit/color"
	"github.com/ifacker/myutil"
	"log"
	"net/http"
)

func init() {
	config.LogConfigInit()

	control.LoadConfig()
	flag.Init()

	if !config.NotLogo {
		color.C256(46).Println(config.Logo)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	err := myutil.InitProxy(tr, config.PROXY)
	if err != nil {
		log.Println(err)
	}
	config.Tr = tr

}

func main() {
	//inputReader := bufio.NewReader(os.Stdin)
	//fmt.Println("Please enter some input: ")
	//input, err := inputReader.ReadString('\n')
	//if err == nil {
	//	fmt.Printf("The input was: %s\n", input)
	//}

	//log.SetFlags(log.Lshortfile)
	control.Init()
}
