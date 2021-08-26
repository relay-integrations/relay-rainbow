package rainbow_rest

type loginS2SBody struct {
	Connection loginS2SBodyConnection `json:"connection"`
}

type loginS2SBodyConnection struct {
	CallbackURL string `json:"callback_url"`
}

type S2SLoginConnection struct {
	Resource    string `json:"resource"`
	ID          string `json:"id"`
	CallbackURL string `json:"callback_url"`
}

type S2SLoginResponse struct {
	Data S2SLoginConnection `json:"data"`
}

type sendPresenceParam struct {
	Presence sendPresenceParamPresence `json:"presence"`
}

type sendPresenceParamPresence struct {
	Show   string `json:"show"`
	Status string `json:"status"`
}

type joinBubbleParam struct {
	Role string `json:"role"`
}

type sendMessageParam struct {
	Message sendMessageParamMessage `json:"message"`
}

type sendMessageParamMessage struct {
	Body string `json:"body"`
}
