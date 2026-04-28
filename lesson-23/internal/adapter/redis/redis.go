package redis

import (
	"time"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/pkg/redis"
)

const (
	idempotencyPrefix = "idempotency:"
	ttl               = time.Hour
)

type Redis struct {
	redis *redis.Client
}

func New(client *redis.Client) *Redis {
	return &Redis{
		redis: client,
	}
}
