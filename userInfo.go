package harvestovertimelib

import (
	"encoding/json"
	"net/http"

	"github.com/sknutsen/harvestovertimelib/v2/models"
)

func GetUserInfo(client *http.Client, settings models.Settings) (models.UserInfo, error) {
	var url string = "https://api.harvestapp.com/api/v2/users/me"

	userInfo, err := getUserInfo(client, url, settings)
	if err != nil {
		return userInfo, err
	}

	return userInfo, nil
}

func getUserInfo(client *http.Client, url string, settings models.Settings) (models.UserInfo, error) {
	var userInfo models.UserInfo

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		println("Error creating request: " + err.Error())
		return userInfo, err
	}

	req.Header.Add("Harvest-Account-ID", settings.AccountId)
	req.Header.Add("Authorization", "Bearer "+settings.AccessToken)
	req.Header.Add("User-Agent", "Harvest overtime")

	resp, err := client.Do(req)
	if err != nil {
		println("Error sending request: " + err.Error())
		return userInfo, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		println("Error decoding response: " + err.Error())
		return userInfo, err
	}

	return userInfo, nil
}
