package main

import (
	"fmt"
	"net/http"

	"github.com/monochromegane/chat_sample"
)

func main() {
	http.Handle("/", chat_sample.ChatHandler(chat_sample.NewPubSub()))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		fmt.Printf("%v\n", err)
	}
}
