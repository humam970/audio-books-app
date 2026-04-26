package env

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type config struct {
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresPort     string
	PostgresHost     string

	RedisUser     string
	RedisPassword string
	RedisPort     string
	RedisHost     string

	S3Secret string
	S3Key    string
	S3Url    string

	AppEnv string
	Port   string

	JwtKey        []byte
	JwtExpiration time.Duration
	AdminToken    string
	FrontEndUrl   string
	BackEndUrl    string

	TurnstileSecret string
	ResendApiKey    string
	SupportEmail    string
}

func (c *config) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.PostgresUser, c.PostgresPassword, c.PostgresHost, c.PostgresPort, c.PostgresDB)
}

var EnvFilePath = ""

var Config = func() config {
	if EnvFilePath != "" {
		_ = godotenv.Load(EnvFilePath)
	}

	cfg := config{
		PostgresDB:       getRequired("POSTGRES_DB"),
		PostgresUser:     getRequired("POSTGRES_USER"),
		PostgresPassword: getRequired("POSTGRES_PASSWORD"),
		PostgresPort:     getRequired("POSTGRES_PORT"),
		PostgresHost:     getRequired("POSTGRES_HOST"),

		RedisUser:     getRequired("REDIS_USER"),
		RedisPassword: getRequired("REDIS_PASSWORD"),
		RedisPort:     getRequired("REDIS_PORT"),
		RedisHost:     getRequired("REDIS_HOST"),

		AppEnv:        "",
		Port:          "",
		JwtKey:        []byte{},
		JwtExpiration: 0,
		AdminToken:    "",

		FrontEndUrl: getRequired("FRONT_END_URL"),
		BackEndUrl:  getRequired("BACK_END_URL"),

		TurnstileSecret: "",
		ResendApiKey:    "",
		SupportEmail:    "",
	}

	return cfg
}()

func getRequired(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Errorf("missing required environment variable: %s", key))
	}
	return val
}
