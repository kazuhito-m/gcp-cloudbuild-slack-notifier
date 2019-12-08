package pubsub

type PubSubMessage struct {
	Data []byte `json:"data"`
}
type Info struct {
	Name  string `json:"name"`
	Place string `json:"place"`
}

// func (i PubSubMessage) IsCloudBuildMessage() bool {
// 	messageJson := string(i.Data)
// 	r := regexp.MustCompile(`"status": ".*"`)

// }

func PickUpStatusText(json string) string {
	return "FAILURE"
}
