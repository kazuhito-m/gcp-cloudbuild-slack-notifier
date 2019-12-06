package pubsub

type PubSubMessage struct {
	Data []byte `json:"data"`
}
type Info struct {
	Name  string `json:"name"`
	Place string `json:"place"`
}
