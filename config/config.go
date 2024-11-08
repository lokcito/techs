package config

import (
    "log"
    "github.com/joho/godotenv"
    "os"
)

func Load() {
    env := os.Getenv("ENV") // Obtén el entorno (por ejemplo, "development" o "production")
    
    if env == "development" {
        err := godotenv.Load(".env")
        if err != nil {
            log.Fatal("Error al cargar el archivo .env")
        }
    } else {
        log.Println("Modo producción: cargando variables de entorno del sistema operativo")
    }
}

func GetEnv(key string) string {
    return os.Getenv(key)
}
