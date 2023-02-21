package messaging

type MessageProcessor interface {
	ProcessInput([]byte)
}
