package gophoenix


type Message struct {
	Topic string `json:"topic"`
	Event event `json:"event"`
	Payload interface{} `json:"payload"`
	Ref int64 `json:"ref"`
}
