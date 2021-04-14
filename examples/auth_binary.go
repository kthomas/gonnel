package main

import (
	"fmt"

	"github.com/kthomas/gonnel"
)

func main() {
	opt := gonnel.Options{
		BinaryPath: "../ngrok-bin/ngrok_linux",
		AuthToken:  "your token string here",
	}

	err := opt.AuthTokenCommand()
	if err != nil {
		fmt.Println(err)
	}
}
