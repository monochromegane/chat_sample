package chat_sample

import "fmt"

type Message struct {
	Name string
	Text string
}

func (m Message) String() string {
	return fmt.Sprintf("%s: %s", m.Name, m.Text)
}
