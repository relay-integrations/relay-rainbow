package rainbow_rest

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func b2a(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func setSecurityHeaders(r *http.Request, loginEmail string, password string, appID string, appSecretKey string) {
	auth := b2a(loginEmail + ":" + password)
	appSum := sha256.Sum256([]byte(appSecretKey + password))
	appAuth := b2a(appID + ":" + hex.EncodeToString(appSum[:]))
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Authorization", "Basic "+auth)
	r.Header.Set("x-rainbow-app-auth", "Basic "+appAuth)

}

func (r *Rest) GetLogin(loginEmail string, password string, appID string, appSecretKey string) (*LoginResponse, error) {
	url := fmt.Sprintf("https://%s/api/rainbow/authentication/v1.0/login", r.host)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	setSecurityHeaders(req, loginEmail, password, appID, appSecretKey)

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(" Error: ", err)
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(" Error: ", err)
		return nil, err
	}

	var lr LoginResponse
	err = json.Unmarshal(data, &lr)
	if err != nil {
		fmt.Println(" Error: ", err)
		return nil, err
	}
	r.SetToken(lr.Token)
	return &lr, nil
}
