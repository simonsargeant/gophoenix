package gophoenix

type messageRouter map[string]ChannelReceiver

func (mr messageRouter) notifyMessage(msg *Message) {
	ch, ok := mr[msg.Topic]
	if !ok {
		return
	}

	switch msg.Event {
	case joinEvent:
		ch.OnJoin(msg.Payload)
	case errorEvent:
		ch.OnJoinError(msg.Payload)
	case closeEvent:
		ch.OnChannelClose(msg.Payload)
	case replyEvent:
		ch.OnReply(msg.Ref, msg.Payload)
	default:
		ch.OnMessage(msg.Event, msg.Payload)
	}
}


