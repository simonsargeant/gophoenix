package gophoenix

type Transport interface {
	Connect(url string) error
	Push(data interface{}) error
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
