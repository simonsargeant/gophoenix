package gophoenix

type Client struct {
	t Transport
	mr *messageRouter
	cr ConnectionReceiver
}

func NewWebsocketClient(cr ConnectionReceiver) *Client {

	mr := NewMessageRouter()
	return &Client {
		t: &socketTransport{},
		cr: cr,
		mr: mr,
	}
}

func (c *Client) Connect(url string) {
	c.t.Connect(url, c.mr, c.cr)
}

func (c *Client) Close() {
	c.t.Close()
}

func (c *Client) Join(callbacks ChannelReceiver, topic string, payload interface{}) Channel {
	var start int64
	rr := newReplyRouter()
	ch := Channel{topic: topic, t: c.t, rc: &atomic_ref{ref: &start}, rr: rr, ln: func() {c.mr.unsubscribe(topic)}}
	c.mr.subscribe(topic, callbacks, rr)
	ch.Join(payload)
	return ch
}
