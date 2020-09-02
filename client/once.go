package client

type onceID struct {
	playerID     byte
	playerIDChan chan struct{}
}

func newOnceID() *onceID {
	return &onceID{playerIDChan: make(chan struct{}, 1)}
}

func (o *onceID) get() byte {
	if o.playerIDChan != nil {
		<-o.playerIDChan
		close(o.playerIDChan)
		o.playerIDChan = nil
	}

	return o.playerID
}

func (o *onceID) set(b byte) {
	if o.playerIDChan != nil {
		o.playerID = b
		o.playerIDChan <- struct{}{}
	}
}
