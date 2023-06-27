package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

const DefaultConfigPath = "./env"

func main() {
	if err := LoadEnvFile(); err != nil {
		log.Fatal(err)
	}

}

func LoadEnvFile() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file. Be sure you have created one in the root of this directory: %s", err)
	}
	return nil
}
