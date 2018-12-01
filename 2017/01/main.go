package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kieron-pivotal/advent2017/01/captcha"
)

func main() {
	usage := fmt.Sprintf("%s <captchaFile>", os.Args[0])
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	contents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(captcha.Decode(string(contents)))
}
