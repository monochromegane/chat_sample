package chat_sample

type callback func(m Message)

func NewPubSub() *pubSub {
	ps := &pubSub{
		subChan: make(chan callback),
		pubChan: make(chan Message),
	}
	go func() {
		for {
			select {
			case callback := <-ps.subChan:
				ps.callbacks = append(ps.callbacks, callback)
			case message := <-ps.pubChan:
				for _, callback := range ps.callbacks {
					go callback(message)
				}
			}
		}
	}()
	return ps
}

type pubSub struct {
	subChan   chan callback
	pubChan   chan Message
	callbacks []callback
}

func (p *pubSub) subscribe(c callback) {
	p.subChan <- c
}

func (p *pubSub) publish(m Message) {
	p.pubChan <- m
}
