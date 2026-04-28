package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/httpserver"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/logger"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/otel"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/postgres"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/redis"
)

type App struct {
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App      App
	HTTP     httpserver.Config
	Logger   logger.Config
	OTEL     otel.Config
	Postgres postgres.Config
	Redis    redis.Config
}

func New() (Config, error) {
	var config Config

	err := godotenv.Load(".env")
	if err != nil {
		return config, fmt.Errorf("godotenv.Load: %w", err)
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return config, fmt.Errorf("envconfig.Process: %w", err)
	}

	return config, nil
}
