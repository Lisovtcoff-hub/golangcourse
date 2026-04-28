package postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/internal/domain"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/otel/tracer"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/transaction"
)

func (p *Postgres) CreateProfile(ctx context.Context, profile domain.Profile) error {
	ctx, span := tracer.Start(ctx, "adapter postgres CreateProfile")
	defer span.End()

	const sql = `INSERT INTO profile (id, name, age, status, verified, contacts)
                    VALUES ($1, $2, $3, $4, $5, $6)`

	contacts, err := json.Marshal(profile.Contacts)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	args := []any{
		profile.ID,
		profile.Name,
		profile.Age,
		profile.Status.String(),
		profile.Verified,
		contacts,
	}

	txOrPool := transaction.TryExtractTX(ctx)

	_, err = txOrPool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	return nil
}
