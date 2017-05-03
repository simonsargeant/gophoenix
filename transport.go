package gophoenix

// Transport is used to establish the connection.
type Transport interface {
	Connect(url string, messageReceiver MessageReceiver, connectionReceiver ConnectionReceiver) error
	Push(data interface{}) error
	Close()
}

// MessageReceiver is called by the Transport to notify of a message.
type MessageReceiver interface {
	NotifyMessage(msg *Message)
}

// ConnectionReceiver is called by the Transport to notify of a connection change.
type ConnectionReceiver interface {
	NotifyConnect()
	NotifyDisconnect()
}

type messageReceiver interface {
	notifyMessage(msg *Message)
}
