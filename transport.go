package gophoenix

type Transport interface {
	Connect(url string, messageReceiver MessageReceiver, connectionReceiver ConnectionReceiver) error
	Push(data interface{}) error
	Close()
}

type notificationReceiver interface {
	ConnectionReceiver
	messageReceiver
}

type ConnectionReceiver interface {
	NotifyConnect()
	NotifyDisconnect()
}

type messageReceiver interface {
	notifyMessage(msg *Message)
}
