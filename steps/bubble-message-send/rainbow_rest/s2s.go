package rainbow_rest

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const S2SPrefix = "/api/rainbow/ucs/v1.0/"

func (r *Rest) LoginS2S(callbackUrl string) (*S2SLoginConnection, error) {
	var result S2SLoginResponse

	body := loginS2SBody{
		Connection: loginS2SBodyConnection{
			CallbackURL: callbackUrl,
		},
	}

	response, err := r.post("api/rainbow/ucs/v1.0/connections", nil, nil, &body, &result)
	if response != nil {
		fmt.Println("loginS2S ->", string(response.Body()))
	} else {
		fmt.Println("loginS2S -> NIL", err)
	}

	if err == nil && !response.IsSuccess() {
		err = fmt.Errorf("%d: %s", response.StatusCode(), response.Status())
	}

	return &result.Data, err
}

func (r *Rest) SendPresence(cnxId string) (*resty.Response, error) {
	body := sendPresenceParam{
		sendPresenceParamPresence{
			Show:   "",
			Status: "",
		},
	}

	response, err := r.put(S2SPrefix+"connections/{cnxId}/presences", map[string]string{"cnxId": cnxId}, nil, body, nil)
	if response != nil {
		fmt.Printf("SendPresence %s -> %s\n", cnxId, string(response.Body()))
	}
	if err == nil && !response.IsSuccess() {
		err = fmt.Errorf("%d: %s", response.StatusCode(), response.Status())
	}
	return response, err
}

func (r *Rest) JoinBubble(cnxId string, roomID string) (*resty.Response, error) {
	body := joinBubbleParam{
		Role: "member",
	}

	response, err := r.post(S2SPrefix+"connections/{cnxId}/rooms/{roomId}/join", map[string]string{"cnxId": cnxId, "roomId": roomID}, nil, body, nil)
	if response != nil {
		fmt.Printf("JoinBubble %s, %s -> %s\n", cnxId, roomID, string(response.Body()))
	}

	if err == nil && !response.IsSuccess() {
		err = fmt.Errorf("%d: %s", response.StatusCode(), response.Status())
	}

	return response, err
}

func (r *Rest) SendMessageToConversation(cnxId string, convId string, message string) (*resty.Response, error) {
	body := sendMessageParam{
		sendMessageParamMessage{
			Body: message,
		},
	}

	response, err := r.post(S2SPrefix+"connections/{cnxId}/conversations/{convId}/messages", map[string]string{"cnxId": cnxId, "convId": convId}, nil, body, nil)
	if response != nil {
		fmt.Printf("SendMessageToConversation( %s, %s, %s ) -> %s\n", cnxId, convId, message, string(response.Body()))
	}
	if err == nil && !response.IsSuccess() {
		err = fmt.Errorf("%d: %s", response.StatusCode(), response.Status())
	}
	return response, err
}

func (r *Rest) DeleteS2SConnection(cnxID string) (*resty.Response, error) {

	response, err := r.delete(S2SPrefix+"connections/{cnxID}", map[string]string{"cnxID": cnxID}, nil)
	if response != nil {
		fmt.Printf("DeleteS2SConnection( %s ) -> %s\n", cnxID, string(response.Body()))
	}
	if err == nil && !response.IsSuccess() {
		err = fmt.Errorf("%d: %s", response.StatusCode(), response.Status())
	}
	return response, err
}
