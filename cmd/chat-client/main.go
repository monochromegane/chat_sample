package main

import (
	"flag"

	"github.com/monochromegane/chat_sample"
)

var name string

func init() {
	flag.StringVar(&name, "name", "unknown", "your name")
	flag.Parse()
}

func main() {
	client := chat_sample.Client{
		Host: "localhost",
		Port: 1234,
		Name: name,
	}
	client.Start()
}
