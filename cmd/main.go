package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dylanconnolly/home-forecast/internal/http"
	"github.com/joho/godotenv"
)

const DefaultConfigPath = "./env"

func main() {
	if err := LoadEnvFile(); err != nil {
		log.Fatal(err)
	}
	gc := http.NewGoogleOauthClient()
	gc.RefreshAccessToken()
	fmt.Printf("refreshed token\n %+v", gc.Config.AccessToken)
	nc := http.NewNestClient()
	nc.GetDevices(gc.Config.AccessToken)
}

func LoadEnvFile() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file. Be sure you have created one in the root of this directory: %s", err)
	}
	if s := os.Getenv("CLIENT_ID"); s == "" {
		return fmt.Errorf("no CLIENT_ID found in .env file. Please add your client ID from https://console.cloud.google.com/apis/credentials.")
	}
	if s := os.Getenv("CLIENT_SECRET"); s == "" {
		return fmt.Errorf("no CLIENT_SECRET found in .env file. Please add your client secret from https://console.cloud.google.com/apis/credentials.")
	}
	if s := os.Getenv("REFRESH_TOKEN"); s == "" {
		return fmt.Errorf("no REFRESH_TOKEN found in .env file. Please add your refresh token after generating an access token from https://developers.google.com/nest/device-access/authorize.")
	}
	if s := os.Getenv("PROJECT_ID"); s == "" {
		return fmt.Errorf("no PROJECT_ID found in .env file. Please add your project ID from https://console.nest.google.com/device-access/.")
	}
	return nil
}
