package oauth

import (
	"fmt"
	"net/http"
	"os"
)

type GoogleOauth struct {
	AccessToken  string
	refreshToken string
}

const googleRefreshBaseUrl = "https://www.googleapis.com/oauth2/v4/token?"

func (g *GoogleOauth) RefreshToken() (*GoogleOauth, error) {
	url := fmt.Sprintf("%sclient_id=%s&client_secret=%s&refresh_token=%s&grant_type=refresh_token",
		googleRefreshBaseUrl,
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		g.refreshToken)

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return nil, fmt.Errorf("error with refresh token request: %s", err)
	}
	defer resp.Body.Close()

	fmt.Printf("response body: %+v", resp.Body)
	return nil, nil
}
