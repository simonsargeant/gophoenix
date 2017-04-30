package gophoenix

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type socketTransport struct {
	socket *websocket.Conn
	cr ConnectionReceiver
	mr messageReceiver
	close chan struct{}
	done chan struct{}
}

func (st *socketTransport) Connect(url string) error {
	// TODO Add origin header, handle resp from dial
	conn, _, err := websocket.DefaultDialer.Dial(url, http.Header{})

	if err != nil {
		return err
	}

	st.socket = conn
	go st.listen()
	st.cr.NotifyConnect()

	return err
}

func (st *socketTransport) Push(data interface{}) error {
	return st.socket.WriteJSON(data)
}

func (st *socketTransport) Close() {
	st.close<- struct {}{}
	<-st.done
}

func (st *socketTransport) listen() {
	defer st.stop()
	for {
		select {
		case <-st.close:
			return
		default:
		}

		var msg *Message
		err := st.socket.ReadJSON(msg)

		if err != nil {
			return
		}

		st.mr.notifyMessage(msg)
	}
}

func (st *socketTransport) stop() {
	st.socket.Close()
	st.cr.NotifyDisconnect()
	func() {st.done<- struct{}{}} ()
}
