package services

import (
	"context"
	"github.com/hieuus/home-services/config"
	"github.com/hieuus/home-services/internal/repositories"
	pb "github.com/hieuus/home-services/pb"
	"github.com/rs/zerolog"
)

type serviceInterface interface {
	pb.HomeServicesAdminServiceServer
}

var _ serviceInterface = &Service{}

type Service struct {
	log  zerolog.Logger
	cfg  *config.Config
	repo repositories.Repository
}

func (s *Service) UserHealthz(ctx context.Context, request *pb.UserHealthzRequest) (*pb.UserHealthzResponse, error) {
	return &pb.UserHealthzResponse{Message: "OK"}, nil
}

func (s *Service) Healthz(ctx context.Context, request *pb.HealthzRequest) (*pb.HealthzResponse, error) {
	return &pb.HealthzResponse{Message: "OK"}, nil
}

func New(
	log zerolog.Logger,
	cfg *config.Config,
	repo repositories.Repository,
) *Service {
	return &Service{
		log:  log,
		cfg:  cfg,
		repo: repo,
	}
}
