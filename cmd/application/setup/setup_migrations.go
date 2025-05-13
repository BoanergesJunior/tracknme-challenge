package setup

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

type MigrateConfig struct {
	Path   string
	Schema string
}

type MigrateResult struct {
	Version uint
	Dirty   bool
}

func RunMigration(orm *gorm.DB, config *MigrateConfig) (*MigrateResult, error) {
	db, err := orm.DB()
	if err != nil {
		return nil, err
	}

	if err := orm.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", config.Schema)).Error; err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTableQuoted: true,
		MigrationsTable:       fmt.Sprintf(`"%s"."schema_migrations"`, config.Schema),
	})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", config.Path),
		"postgres",
		driver,
	)

	if err != nil {
		return nil, err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	version, dirty, err := m.Version()
	if err != nil {
		return nil, err
	}

	return &MigrateResult{
		Version: version,
		Dirty:   dirty,
	}, nil
}
