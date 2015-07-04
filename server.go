package chat_sample

import "golang.org/x/net/websocket"

func ChatHandler(ps *pubSub) websocket.Handler {
	return func(ws *websocket.Conn) {
		ps.subscribe(func(m Message) {
			websocket.JSON.Send(ws, m)
		})

		for {
			var m Message
			if err := websocket.JSON.Receive(ws, &m); err != nil {
				break
			}
			ps.publish(m)
		}
	}
}
