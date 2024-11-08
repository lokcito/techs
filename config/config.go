package config

import (
    "log"
    "github.com/joho/godotenv"
    "os"
)

func Load() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func GetEnv(key string) string {
    return os.Getenv(key)
}
