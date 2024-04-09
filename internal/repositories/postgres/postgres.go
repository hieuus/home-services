package postgres

import (
	"github.com/hieuus/home-services/config"
	"github.com/hieuus/home-services/internal/repositories"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Queries struct {
	db  *gorm.DB
	cfg *config.Config
	log zerolog.Logger
}

func New(log zerolog.Logger, cfg *config.Config, db *gorm.DB) *Queries {
	return &Queries{
		db:  db,
		cfg: cfg,
		log: log,
	}
}

func (q *Queries) Ping() error {
	sqlDB, err := q.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func (q *Queries) WithTx(tx *gorm.DB) *Queries {
	nr := *q
	nr.db = tx
	return &nr
}

func (q *Queries) Transaction(txFunc func(repositories.Repository) error) (err error) {
	// start new transaction
	tx := q.db.Begin()
	defer func() {
		p := recover()
		switch {
		case p != nil:
			execErr := tx.Rollback().Error
			if execErr != nil {
				q.log.Error().Err(execErr)
			}
			panic(p) // re-throw panic after Rollback
		case err != nil:
			execErr := tx.Rollback().Error // err is non-nil; don't change it
			if execErr != nil {
				q.log.Error().Err(execErr)
			}
		default:
			err = tx.Commit().Error // err is nil; if Commit returns error update err
		}
	}()
	return txFunc(q.WithTx(tx))
}
