package postgres

import (
	"context"
	"github.com/hieuus/home-services/config"
	"github.com/hieuus/home-services/internal/models"
	"github.com/hieuus/home-services/internal/models/store"
	"github.com/hieuus/home-services/internal/must"
	ll "github.com/hieuus/home-services/pkg/log"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"testing"
)

func TestQueries_RecordUser(t *testing.T) {
	l := ll.New()
	cfg := config.Load()
	db := must.ConnectPostgres(cfg.Postgres)

	type fields struct {
		db  *gorm.DB
		cfg *config.Config
		log zerolog.Logger
	}
	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				db:  db,
				cfg: cfg,
				log: l,
			},
			args: args{
				ctx: context.Background(),
				user: &models.User{
					User: store.User{
						Code:          "ABCD",
						FullName:      "Hieu Nguyen",
						Email:         "nvhieu.us@gmail.com",
						PhoneNumber:   "84931902050",
						Password:      "123",
						Salt:          "123",
						IsVerifyPhone: false,
						IsVerifyEmail: false,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db:  tt.fields.db,
				cfg: tt.fields.cfg,
				log: tt.fields.log,
			}
			if err := q.RecordUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("RecordUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
