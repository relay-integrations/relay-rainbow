package rainbow_rest

import "time"

type GetConversationsResponse struct {
	Data  []Conversation `json:"data"`
	Total int            `json:"total"`
}

type Conversation struct {
	ID                            string      `json:"id"`
	CreationDate                  time.Time   `json:"creationDate"`
	PeerID                        string      `json:"peerId"`
	Type                          string      `json:"type"`
	Mute                          bool        `json:"mute"`
	LoginEmail                    string      `json:"loginEmail,omitempty"`
	FirstName                     string      `json:"firstName,omitempty"`
	LastName                      string      `json:"lastName,omitempty"`
	DisplayName                   string      `json:"displayName,omitempty"`
	JidIm                         string      `json:"jid_im"`
	UseScreenSharingCustomisation string      `json:"useScreenSharingCustomisation,omitempty"`
	UseWebRTCAudioCustomisation   string      `json:"useWebRTCAudioCustomisation,omitempty"`
	UseWebRTCVideoCustomisation   string      `json:"useWebRTCVideoCustomisation,omitempty"`
	InstantMessagesCustomisation  string      `json:"instantMessagesCustomisation,omitempty"`
	ReadReceiptsCustomisation     string      `json:"readReceiptsCustomisation,omitempty"`
	LastAvatarUpdateDate          time.Time   `json:"lastAvatarUpdateDate"`
	LastMessageText               string      `json:"lastMessageText"`
	LastMessageSender             string      `json:"lastMessageSender"`
	LastMessageDate               time.Time   `json:"lastMessageDate"`
	UnreceivedMessageNumber       string      `json:"unreceivedMessageNumber"`
	UnreadMessageNumber           string      `json:"unreadMessageNumber"`
	Call                          interface{} `json:"call,omitempty"`
	Geoloc                        bool        `json:"geoloc,omitempty"`
	ContentType                   string      `json:"contentType,omitempty"`
	Name                          string      `json:"name,omitempty"`
	Topic                         string      `json:"topic,omitempty"`
	AvatarID                      string      `json:"avatarId,omitempty"`
}
