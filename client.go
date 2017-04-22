package gophoenix

type Client struct {
	t Transport
	mr messageRouter
}

func NewWebsocketClient(cr ConnectionReceiver) *Client {
	mr := make(messageRouter)
	return &Client {
		t: &socketTransport{
			cr: cr,
			mr: mr,
		},
		mr: mr,
	}
}

func (c *Client) Connect(url string) {
	c.t.Connect(url)
}

func (c *Client) Join(topic string, payload interface{}) Channel {
	var start int64
	ch := Channel{topic: topic, t: c.t, rc: &atomic_ref{ref: &start}}
	ch.Join(payload)
	return ch
}
