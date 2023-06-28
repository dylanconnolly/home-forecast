package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	accessTokenURL  = "https://www.googleapis.com/oauth2/v4/token?grant_type=authorization_code&redirect_uri=https://www.google.com"
	refreshTokenURL = "https://www.googleapis.com/oauth2/v4/token?grant_type=refresh_token"
)

type GoogleOauthConfig struct {
	AccessToken     string
	AccessTokenURL  string
	ClientID        string
	ClientSecret    string
	RefreshTokenURL string
	RefreshToken    string
}

type GoogleOauthClient struct {
	Config     *GoogleOauthConfig
	HttpClient *http.Client
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int16  `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func NewGoogleOauthClient() *GoogleOauthClient {
	config := &GoogleOauthConfig{
		AccessToken:     "",
		AccessTokenURL:  accessTokenURL,
		ClientID:        os.Getenv("CLIENT_ID"),
		ClientSecret:    os.Getenv("CLIENT_SECRET"),
		RefreshTokenURL: refreshTokenURL,
		RefreshToken:    os.Getenv("REFRESH_TOKEN"),
	}
	return &GoogleOauthClient{
		Config:     config,
		HttpClient: NewHttpClient(5),
	}
}

func (gc *GoogleOauthClient) RefreshAccessToken() error {
	var oauthResp TokenResponse

	url := fmt.Sprintf("%s&client_id=%s&client_secret=%s&refresh_token=%s", gc.Config.RefreshTokenURL, gc.Config.ClientID, gc.Config.ClientSecret, gc.Config.RefreshToken)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("error creating refresh token request: %s", err)
	}

	resp, err := gc.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error requesting refresh token: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %s", err)
	}
	if err := json.Unmarshal(body, &oauthResp); err != nil {
		return fmt.Errorf("error unmarshalling response: %s", err)
	}

	gc.Config.AccessToken = oauthResp.AccessToken

	return nil
}

func (gc *GoogleOauthClient) AccessToken() string {
	return gc.Config.AccessToken
}
