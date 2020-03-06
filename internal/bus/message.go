package bus

type Message struct {
	Topic   string
	Pattern string
	Payload interface{}
}
