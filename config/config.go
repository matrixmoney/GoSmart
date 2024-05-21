package config

import (
    "log"
    "github.com/joho/godotenv"
    "os"
)

var (
    SmartAPIKey     string
    SmartAPIClientID string
    SmartAPIMpin    string
)

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    SmartAPIKey = os.Getenv("SMARTAPI_KEY")
    SmartAPIClientID = os.Getenv("SMARTAPI_CLIENT_ID")
    SmartAPIMpin = os.Getenv("SMARTAPI_MPIN")
}
