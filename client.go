package gophoenix

// Client is the entry point for a phoenix channel connection
type Client struct {
	t Transport
	mr *messageRouter
	cr ConnectionReceiver
}

// NewWebsocketClient creates the default connection using a websocket as the transport.
func NewWebsocketClient(cr ConnectionReceiver) *Client {
	return &Client {
		t: &socketTransport{},
		cr: cr,
	}
}

// Connect should be called to establish the connection through the transport.
func (c *Client) Connect(url string) {
	mr := newMessageRouter()
	c.t.Connect(url, mr, c.cr)
}

// Close closes the connection via the transport
func (c *Client) Close() {
	c.t.Close()
}

// Join subscribes to a channel via the transport and returns a reference to the channel.
func (c *Client) Join(callbacks ChannelReceiver, topic string, payload interface{}) *Channel {
	var start int64
	rr := newReplyRouter()
	ch := &Channel{topic: topic, t: c.t, rc: &atomicRef{ref: &start}, rr: rr, ln: func() {c.mr.unsubscribe(topic)}}
	c.mr.subscribe(topic, callbacks, rr)
	ch.join(payload)
	return ch
}
