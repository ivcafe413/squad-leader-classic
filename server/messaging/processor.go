package messaging

type MessageProcessor interface {
	ProcessInput(interface{}) []byte
}
