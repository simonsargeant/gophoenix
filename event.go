package gophoenix

type event string

const (
	joinEvent event = "phx_join"
	closeEvent event = "phx_close"
	errorEvent event = "phx_error"
	replyEvent event = "phx_reply"
	leaveEvent event = "phx_leave"
)

