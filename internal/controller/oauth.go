package controller

type GoogleOauthService interface {
	RefreshAccessToken() error
	AccessToken() string
}
