package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bilibili-live/cowsay/cowsay"
)

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	// https://stackoverflow.com/a/51125922
	data, _ := ioutil.ReadAll(reader)
	var result = string(data)

	cowsay.RenderCowsay(result)
	var echo = cowsay.RenderCowsay(result)
	fmt.Println(echo)
}
