package main

import (
	"fmt"
	"os"

	"github.com/JulioZittei/go-i18n-app/locale/messages"
)


func main () {
	if len(os.Args) < 3 {
		fmt.Println("Use: go run main.go <key> <lang> <args...>")
		os.Exit(1)
	}

	key := os.Args[1]
	lang := os.Args[2]
	args := os.Args[3:]

	message, err := messages.GetMessage(key, lang, args...)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(message)
}