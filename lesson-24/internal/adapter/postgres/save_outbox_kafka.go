package postgres

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-24/internal/domain"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-24/pkg/otel/tracer"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-24/pkg/transaction"
)

func (p *Postgres) SaveOutboxKafka(ctx context.Context, events ...domain.Event) error {
	ctx, span := tracer.Start(ctx, "adapter postgres SaveOutboxKafka")
	defer span.End()

	if len(events) == 0 {
		return nil
	}

	batch := make([]any, 0, len(events))

	for _, e := range events {
		if e.Topic == "" {
			return domain.ErrEmptyTopic
		}

		batch = append(batch, goqu.Record{
			"topic": e.Topic,
			"key":   e.Key,
			"value": e.Value,
		})
	}

	sql, _, err := goqu.Insert("outbox").Rows(batch...).ToSQL()
	if err != nil {
		return fmt.Errorf("goqu.Insert.ToSQL: %w", err)
	}

	txOrPool := transaction.TryExtractTX(ctx)

	_, err = txOrPool.Exec(ctx, sql)
	if err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	return nil
}
