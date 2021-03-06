package types


type RBase struct {
	Ok          bool    `json:"ok"`
	ErrorCode   *int    `json:"error_code,omitempty"`
	Description *string `json:"description,omitempty"`
}
// Update object.
type Update struct {
	UpdateId   float64 `json:"update_id"`
	Message    *Message `json:"message"`
}

type RGetUpdates struct {
	RBase
	Result []Update `json:"result"`
}

// User object.
type User struct {
	Id         float64 `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Username   string `json:"username"`
}


type RGetUser struct {
	RBase
	Result User `json:"result,omitempty"`
}

// GroupChat object.
type GroupChat struct {
	Id      float64 `json:"id"`
	Title   string `json:"title"`
}

// Chat object. 
type Chat struct {
	Id        float64 `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Username  string  `json:"username"`
	Title     string  `json:"title"`
}

// Message object.
type Message struct {
	Message_Id          int64 `json:"message_id"`
	From                *User `json: "from"`
	Date                float64 `json:"date"`
	Chat                *Chat `json:"chat"`
	ForwardFrom         *User `json:"forward_from,omitempty"`
	ForwardDate         float64 `json:"forward_date,omitempty"`
	Text                string `json:"text,omitempty"`
	Caption             string `json:"caption,omitempty"`
	NewChatParticipant  *User `json:"new_chat_participant,omitempty"`
	LeftChatParticipant *User `json:"left_chat_participant,omitempty"`
	NewChatTitle        string `json:"new_chat_title,omitempty"`
	GroupChatCreated    bool `json:"group_chat_created,omitempty"`
}
