package config

import (
    "os"
    "fmt"
	"github.com/joho/godotenv"
)

func Config(key string) string{

    err := godotenv.Load(".env")
    if err != nil {
        fmt.Print("Error al leer el archivo .env ")
    }
    return os.Getenv(key)

}