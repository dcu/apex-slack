package slack

// AttachmentField is used to add a table-like structure to the attachment. See: https://api.slack.com/docs/message-attachments#fields
type AttachmentField struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

// Attachment allows to add attachments to the message. See https://api.slack.com/docs/message-attachments
type Attachment struct {
	Fallback      string             `json:"fallback,omitempty"`
	Color         string             `json:"color,omitempty"`
	Pretext       string             `json:"pretext,omitempty"`
	AuthorName    string             `json:"author_name,omitempty"`
	AuthorLink    string             `json:"author_link,omitempty"`
	AuthorIconURL string             `json:"author_icon,omitempty"`
	Title         string             `json:"title,omitempty"`
	TitleLink     string             `json:"title_link,omitempty"`
	Text          string             `json:"text,omitempty"`
	Fields        []*AttachmentField `json:"fields,omitempty"`
	ImageURL      string             `json:"image_url,omitempty"`
	ThumbURL      string             `json:"thumb_url,omitempty"`
	Footer        string             `json:"footer,omitempty"`
	FooterIconURL string             `json:"footer_icon,omitempty"`
	Timestamp     int64              `json:"ts,omitempty"`
}
