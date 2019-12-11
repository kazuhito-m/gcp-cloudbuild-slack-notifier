package slack

type SlackNotify struct {
	Username    string       `json:"username"`
	IconURL     string       `json:"icon_url"`
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	Mrkdwn      bool         `json:"mrkdwn"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Title     string       `json:"title"`
	TitleLink string       `json:"title_link"`
	Color     string       `json:"color"`
	Fields    []SlackField `json:"fields"`
}

type SlackField struct {
	Title string `json:"title"`
	Value string `json:"value"`
}
