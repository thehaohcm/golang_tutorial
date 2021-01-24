package config

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var (
	ClientID         string
	ClientSecret     string
	Scope            string
	RedirectURL      string
	OauthStateString string
	UploadFolderPath string

	oauthConf *oauth2.Config
	TenantID  string
	// AccessToken      string
	RefreshToken     string
	Expiry           time.Time
	ExpiredTokenTime time.Time
	TokenType        string

	SavedToken  *oauth2.Token
	StaticToken oauth2.TokenSource

	//link stuffs
	RefreshAPIEndPoint string
	UploadAPIEndPoint  string
	ShareAPIEndPoint   string

	//stuffs
	FragSize       int
	ShareBodyJson  string
	UploadBodyJson string
)

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
		panic(err)
	}

	ClientID = viper.GetString("CLIENT_ID")
	ClientSecret = viper.GetString("CLIENT_SECRET")
	Scope = viper.GetString("SCOPE")
	RedirectURL = viper.GetString("REDIRECT_URL")
	TenantID = viper.GetString("TENANT_ID")
	// AccessToken = viper.GetString("AccessToken")
	RefreshToken = viper.GetString("REFRESH_TOKEN")
	Expiry = time.Now().Add(time.Duration(viper.GetInt("EXPIRY")) * time.Second)
	TokenType = viper.GetString("TOKEN_TYPE")
	UploadFolderPath = viper.GetString("UPLOAD_FOLDER_PATH")
	RefreshAPIEndPoint = strings.Replace(viper.GetString("REFESH_API_ENDPOINT"), "{TENANT_ID}", TenantID, 1)
	UploadAPIEndPoint = strings.Replace(viper.GetString("UPLOAD_API_ENDPOINT"), "{UPLOAD_FOLDER_PATH}", UploadFolderPath, 1)
	ShareAPIEndPoint = viper.GetString("SHARE_API_ENDPOINT")
	FragSize = viper.GetInt("FRAG_SIZE")

	ShareBodyJson = viper.GetString("SHARE_BODY_JSON")

	UploadBodyJson = viper.GetString("UPLOAD_BODY_JSON")

	SavedToken = &oauth2.Token{
		AccessToken:  "",
		RefreshToken: RefreshToken,
		Expiry:       Expiry,
		TokenType:    TokenType,
	}

	StaticToken = oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: SavedToken.AccessToken},
	)
}
