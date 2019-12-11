package slack

type SlackNotify struct {
	Username    string       `json:"username"`
	IconURL     string       `json:"icon_url"`
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	Mrkdwn      bool         `json:"mrkdwn"`
	Attachments []Attachment `json:"attachments"`
	Blocks      []Block      `json:"blocks"`
}

type Attachment struct {
	Title     string            `json:"title"`
	TitleLink string            `json:"title_link"`
	Color     string            `json:"color"`
	Fields    []AttachmentField `json:"fields"`
	Blocks    []Block           `json:"blocks"`
}

type AttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type Block struct {
	Type    string       `json:"type"`
	Text    BlockText    `json:"text,omitempty"`
	BlockID string       `json:"block_id,omitempty"`
	Fields  []BlockField `json:"fields,omitempty"`
}

type BlockField struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type BlockText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
