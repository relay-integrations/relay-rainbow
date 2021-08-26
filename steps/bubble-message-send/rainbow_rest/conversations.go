package rainbow_rest

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func (r *Rest) GetConversations(userID string, extraParams map[string]string) (*resty.Response, []Conversation, error) {
	var result GetConversationsResponse
	if extraParams == nil {
		extraParams = map[string]string{}
	}
	extraParams["userId"] = userID
	extraParams["format"] = "full"

	response, err := r.get(EndUserPrefix+"users/{userID}/conversations", map[string]string{"userID": userID}, extraParams, &result)

	if response != nil {
		fmt.Printf("GetConversations %s -> %s\n", userID, string(response.Body()))
	}

	return response, result.Data, err
}
