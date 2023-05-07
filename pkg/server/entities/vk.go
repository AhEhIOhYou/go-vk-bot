package entities

type Event struct {
	Type    string      `json:"type"`
	Object  interface{} `json:"object"`
	GroupID int         `json:"group_id"`
}

type MessageNew struct {
	Message    `json:"message"`
	ClientInfo `json:"client_info"`
}

type Message struct {
	ID                    int    `json:"id"`
	Date                  int    `json:"date"`
	PeerID                int    `json:"peer_id"`
	FromID                int    `json:"from_id"`
	Text                  string `json:"text"`
	RandomID              int    `json:"random_id"`
	Payload               string `json:"payload"`
	ConversationMessageId int    `json:"conversation_message_id"`
}

type ClientInfo struct {
	ButtonActions  []string `json:"button_actions"`
	Keyboard       bool     `json:"keyboard"`
	InlineKeyboard bool     `json:"inline_keyboard"`
	Carousel       bool     `json:"carousel"`
	LangID         int      `json:"lang_id"`
}
