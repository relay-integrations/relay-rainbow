package rainbow_rest

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const EndUserPrefix = "api/rainbow/enduser/v1.0/"

func (r *Rest) GetBubbles(userID string, extraParams map[string]string) (*resty.Response, *GetRoomsResponses, error) {
	var result GetRoomsResponses
	if extraParams == nil {
		extraParams = map[string]string{}
	}
	extraParams["userId"] = userID
	extraParams["format"] = "full"

	response, err := r.get(EndUserPrefix+"rooms", nil, extraParams, &result)
	if response != nil {
		fmt.Printf("GetBubbles %s -> %s\n", userID, string(response.Body()))
	}
	return response, &result, err
}

func (r *Rest) GetAllBubbles(userID string) ([]Room, error) {

	var result []Room
	var err error

	limit := 100

	for i := 0; ; i++ {
		extraParams := map[string]string{}

		extraParams["offset"] = fmt.Sprintf("%d", i*limit)
		extraParams["limit"] = fmt.Sprintf("%d", limit)

		_, tmpResult, err := r.GetBubbles(userID, extraParams)

		if err != nil {
			break
		}

		result = append(result, tmpResult.Data...)
		if tmpResult.Offset+limit >= tmpResult.Total {
			break
		}
	}

	return result, err
}
