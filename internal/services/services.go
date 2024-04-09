package services

import (
	"context"
	"github.com/hieuus/home-services/config"
	pb "github.com/hieuus/home-services/pb"
)

type serviceInterface interface {
	pb.HomeServicesAdminServiceServer
}

var _ serviceInterface = &Service{}

type Service struct {
	cfg *config.Config
}

func (s *Service) UserHealthz(ctx context.Context, request *pb.UserHealthzRequest) (*pb.UserHealthzResponse, error) {
	return &pb.UserHealthzResponse{Message: "OK"}, nil
}

func (s *Service) Healthz(ctx context.Context, request *pb.HealthzRequest) (*pb.HealthzResponse, error) {
	return &pb.HealthzResponse{Message: "OK"}, nil
}

func New(cfg *config.Config) *Service {
	return &Service{cfg: cfg}
}
