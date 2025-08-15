package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// Redis
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	// Server
	ServerPort string

	// External API
	DragonBallAPIURL     string
	DragonBallAPITimeout int
}

func Load() (*Config, error) {
	return &Config{
		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "dragon_ball_user"),
		DBPassword: getEnv("DB_PASSWORD", "dragon_ball_pass"),
		DBName:     getEnv("DB_NAME", "dragon_ball_db"),

		// Redis
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", "dragon_ball_redis_pass"),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),

		// Server
		ServerPort: getEnv("API_PORT", "8080"),

		// External API
		DragonBallAPIURL:     getEnv("DRAGONBALL_API_URL", "https://dragonball-api.com/api"),
		DragonBallAPITimeout: getEnvAsInt("DRAGONBALL_API_TIMEOUT", 10),
	}, nil
}

func (c *Config) PostgresConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

func (c *Config) RedisAddress() string {
	return fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
