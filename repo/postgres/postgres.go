package pg

import (
	"context"
	"database/sql"
	"fmt"

	model "github.com/SaeedSH02/payment_gateway/Models"
	"github.com/SaeedSH02/payment_gateway/config"
	logger "github.com/SaeedSH02/payment_gateway/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	bunDebug "github.com/uptrace/bun/extra/bundebug"
	// "github.com/uptrace/bun/migrate"
)

type postgres struct {
	db *bun.DB
}

func NewPostgres(cfg config.Postgres) (*postgres, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
		cfg.SSLMode,
		// cfg.Timezone, // TODO: Timezone should be specified.
	)

	// dsn := "host=localhost user=myuser password=mypass dbname=mydb port=5432 sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	if err := sqldb.Ping(); err != nil {
		return nil, fmt.Errorf("cannot connect to database: %w", err)
	}

	db.AddQueryHook(bunDebug.NewQueryHook(
		bunDebug.WithVerbose(true),
		bunDebug.FromEnv("DEBUG"),
	))
	


	createTables(db,context.Background())
	return &postgres{
		db: db,
	}, nil
}



func createTables(pg *bun.DB,ctx context.Context) error {
	_, err := pg.NewCreateTable().
		Model((*model.Merchant)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		logger.LogError(ctx, "cant create merchant table", err)
		return err
	}

	_, err = pg.NewCreateTable().
		Model((*model.ApiKey)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		logger.LogError(ctx, "cant create api_key table", err)
		return err
	}

	return nil
}
