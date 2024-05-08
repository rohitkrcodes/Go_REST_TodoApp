package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecrets             string
	JWTExpirationInSeconds int64
}

func initConfig() Config {

	godotenv.Load()

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost:0000"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "12345"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "0.0.0"), getEnv("DB_PORT", "0000")),
		DBName:                 getEnv("DB_NAME", "ecom"),
		JWTSecrets:             getEnv("JWT_SECRET", "Secrets revealed"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 60*1*1),
	}
}

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
