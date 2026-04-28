package config

import (
	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/internal/adapter/kafka_produce"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/internal/adapter/postgres"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/internal/adapter/redis"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/internal/controller/kafka_consume"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/pkg/httpserver"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/pkg/logger"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/pkg/otel"
)

// Конфиг приложения
type App struct {
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App           App
	HTTP          httpserver.Config
	Logger        logger.Config
	OTEL          otel.Config
	Postgres      postgres.Config
	Redis         redis.Config
	KafkaProducer kafka_produce.Config
	KafkaConsumer kafka_consume.Config
}

func InitConfig() (Config, error) {
	// Парсим и валидируем конфиг
	return Config{}, nil
}
