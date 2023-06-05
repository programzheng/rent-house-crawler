package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/programzheng/rent-house-crawler/config"
	"github.com/programzheng/rent-house-crawler/ent"

	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	client *ent.Client
	tx     *ent.Tx
}

func getOpenDatabaseClient() *ent.Client {
	client, err := ent.Open(config.Cfg.GetString("database.driver"), config.Cfg.GetString("database.source-name"))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func NewRepository() (*Repository, error) {
	client := getOpenDatabaseClient()
	return &Repository{client: client}, nil
}

func (rp *Repository) StartTransaction(ctx context.Context) error {
	tx, err := rp.client.Tx(ctx)
	if err != nil {
		return err
	}
	rp.client = tx.Client()

	return nil
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
