package rainbow_rest

import "time"

type GetRoomsResponses struct {
	Data   []Room `json:"data"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Total  int    `json:"total"`
}

type Room struct {
	Name                 string        `json:"name"`
	Topic                string        `json:"topic"`
	Visibility           string        `json:"visibility"`
	History              string        `json:"history"`
	DisableNotifications bool          `json:"disableNotifications"`
	IsActive             bool          `json:"isActive"`
	AutoRegister         string        `json:"autoRegister"`
	LastAvatarUpdateDate interface{}   `json:"lastAvatarUpdateDate"`
	GuestEmails          []interface{} `json:"guestEmails"`
	Jid                  string        `json:"jid"`
	Creator              string        `json:"creator"`
	CreationDate         time.Time     `json:"creationDate"`
	Users                []struct {
		Privilege    string    `json:"privilege"`
		Status       string    `json:"status"`
		UserID       string    `json:"userId"`
		JidIm        string    `json:"jid_im"`
		AdditionDate time.Time `json:"additionDate"`
	} `json:"users"`
	ConfEndpoints []struct {
		ConfEndpointID string `json:"confEndpointId"`
		UserID         string `json:"userId"`
		MediaType      string `json:"mediaType"`
	} `json:"confEndpoints"`
	Conference struct {
		Scheduled        bool          `json:"scheduled"`
		DisableTimeStats bool          `json:"disableTimeStats"`
		GuestEmails      []interface{} `json:"guestEmails"`
		MediaType        string        `json:"mediaType"`
	} `json:"conference,omitempty"`
	AutoAcceptInvitation bool      `json:"autoAcceptInvitation"`
	ID                   string    `json:"id"`
	LastActivityDate     time.Time `json:"lastActivityDate"`
	Tags                 []struct {
		Tag   string `json:"tag"`
		Color string `json:"color"`
		Emoji string `json:"emoji,omitempty"`
	} `json:"tags"`
	ContainerID        string `json:"containerId"`
	ContainerName      string `json:"containerName"`
	ActiveUsersCounter int    `json:"activeUsersCounter"`
}
