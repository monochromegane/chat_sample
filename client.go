package chat_sample

import (
	"encoding/json"
	"fmt"

	"github.com/monochromegane/chat_sample_view"
	"golang.org/x/net/websocket"
)

type Client struct {
	Host string
	Port int
	Name string
}

func (c Client) url() string {
	return fmt.Sprintf("ws://%s:%d", c.Host, c.Port)
}

func (c Client) originUrl() string {
	return fmt.Sprintf("http://%s", c.Host)
}

func (c Client) Start() error {
	ws, err := websocket.Dial(c.url(), "", c.originUrl())
	if err != nil {
		return err
	}

	chat := chat_sample_view.NewChatView("CHAT>")
	chat.EnterAction = func(text string) {
		msg := Message{Name: c.Name, Text: text}
		data, _ := json.Marshal(msg)
		ws.Write(data)
	}

	chat.Init()
	defer chat.Close()

	go func() {
		for {
			rcv := make([]byte, 512)
			n, _ := ws.Read(rcv)
			var msg Message
			json.Unmarshal(rcv[0:n], &msg)
			chat.Message(&msg)
		}
	}()

	chat.Draw()
	chat.PollInput()

	return nil
}
