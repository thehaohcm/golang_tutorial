package token

import (
	"encoding/json"
	"fmt"
	"net/http"
	"https://github.com/thehaohcm/simple-onedrive/config"
	"https://github.com/thehaohcm/simple-onedrive/models"
	"strings"
	"time"
)

func init() {
	RefreshToken()
}

func RefreshToken() {
	payload := strings.NewReader("grant_type=refresh_token" +
		"&client_id=" + config.ClientID +
		"&client_secret=" + config.ClientSecret +
		"&scope=" + config.Scope +
		"&redirect_uri=" + config.RedirectURL +
		"&refresh_token=" + config.SavedToken.RefreshToken)
	req, _ := http.NewRequest("POST", config.RefreshAPIEndPoint, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	var jsonResult models.RefreshTokenResponse
	err := json.NewDecoder(res.Body).Decode(&jsonResult)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if jsonResult.AccessToken != config.SavedToken.AccessToken {
		saveToken(&jsonResult)
		fmt.Println("saved a new token")
	} else {
		fmt.Println("nothing changed")
	}

}

func saveToken(tokenJSON *models.RefreshTokenResponse) {
	config.SavedToken.AccessToken = tokenJSON.AccessToken
	config.SavedToken.RefreshToken = tokenJSON.RefreshToken
	config.SavedToken.TokenType = tokenJSON.TokenType

	//assign the new refreshTokenStartTime
	config.ExpiredTokenTime = time.Now().Add(3000 * time.Second)
}
