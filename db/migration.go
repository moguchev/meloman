package db

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/multierr"
)

const (
	sourceURL = "file://db/migrations"
)

func Migrate(databaseURL string) error {
	const api = "Migrate"

	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		return fmt.Errorf("%s.New: %w", api, err)
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		} else if errors.Is(err, migrate.ErrDirty{}) {
			err = multierr.Append(err, m.Down())
		}
		return fmt.Errorf("%s.Up: %w", api, err)
	}

	return nil
}
