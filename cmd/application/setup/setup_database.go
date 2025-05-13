package setup

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Migrations struct {
	Path   string
	Schema string
}

func SetupDatabase(migrations Migrations) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: os.Getenv("DATABASE_POSTGRES"),
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	_, err = RunMigration(db, &MigrateConfig{
		Path:   migrations.Path,
		Schema: migrations.Schema,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
