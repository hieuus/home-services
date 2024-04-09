package must

import (
	"fmt"
	"github.com/hieuus/home-services/config"
	"github.com/hieuus/home-services/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ll = log.New()

func ConnectPostgres(cfg *config.Postgres) *gorm.DB {
	if cfg == nil {
		return nil
	}

	ll.Info().Msgf("connecting to postgres at %s", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	db, err := gorm.Open(postgres.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		ll.Fatal().Err(err).Msg("failed to connect to postgres")
	}

	ll.Info().Msgf("connected to postgres")
	return db
}
