//go:build integration

package test

import (
	"context"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/config"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/internal/app"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/httpclient"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/httpserver"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/otel"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/postgres"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/redis"
)

// Prepare:  make up
// Run test: make integration-test

var ctx = context.Background()

func Test_Integration(t *testing.T) {
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite
	*require.Assertions

	profile     *httpclient.Client
	kafkaWriter *kafka.Writer
	kafkaReader *kafka.Reader
}

func (s *Suite) SetupSuite() { // В начале всех тестов
	s.Assertions = s.Require()

	s.ResetMigrations()

	// Config
	c := config.Config{
		App: config.App{
			Name:    "my-app",
			Version: "test",
		},
		HTTP: httpserver.Config{
			Port: "8080",
		},
		Postgres: postgres.Config{
			Host:     "localhost",
			Port:     "5432",
			User:     "login",
			Password: "pass",
			DBName:   "postgres",
		},
		Redis: redis.Config{
			Addr: "localhost:6379",
		},
	}

	// Logger and OTEL disable
	log.Logger = zerolog.Nop()
	otel.SilentModeInit()

	// Server
	go func() {
		err := app.Run(context.Background(), c)
		s.NoError(err)
	}()

	// API client
	s.profile = httpclient.New(httpclient.Config{Host: "localhost", Port: "8080"})

	time.Sleep(time.Second)
}

func (s *Suite) TearDownSuite() {} // В конце всех тестов

func (s *Suite) SetupTest() {} // Перед каждым тестом

func (s *Suite) TearDownTest() {} // После каждого теста
