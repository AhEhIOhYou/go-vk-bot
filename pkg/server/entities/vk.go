package entities

type Event struct {
	GroupID int                    `json:"group_id"`
	Type    string                 `json:"type"`
	EventID string                 `json:"event_id"`
	Version string                 `json:"v"`
	Object  map[string]interface{} `json:"object"`
	Secret  string                 `json:"secret"`
}

type MessageNew struct {
	Message    MessageRequest `json:"message"`
	ClientInfo ClientInfo     `json:"client_info"`
}

type MessageEvent struct {
	UserID                int    `json:"user_id"`
	PeerID                int    `json:"peer_id"`
	EventID               string `json:"event_id"`
	Payload               string `json:"payload"`
	ConversationMessageId int    `json:"conversation_message_id"`
}

type MessageRequest struct {
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

type MessageResponse struct {
	Message     string   `json:"message" qs:"message"`
	UserID      int      `json:"user_id" qs:"user_id"`
	AccessToken string   `json:"access_token" qs:"access_token"`
	Version     string   `json:"v" qs:"v"`
	RandomID    int      `json:"random_id" qs:"random_id"`
	Keyboard    Keyboard `json:"keyboard" qs:"keyboard"`
}

type Keyboard struct {
	OneTime bool     `json:"one_time" qs:"one_time"`
	Buttons [][]Button `json:"buttons" qs:"buttons"`
	Inline  bool     `json:"inline" qs:"inline"`
}

type Button struct {
	Color  string       `json:"color" qs:"color"`
	Action ButtonAction `json:"action" qs:"action"`
}

type ButtonAction struct {
	Type    string `json:"type" qs:"type"`
	Label   string `json:"label" qs:"label"`
	Payload string `json:"payload" qs:"payload"`
}
