package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hieuus/home-services/config"
	"github.com/hieuus/home-services/internal/services"
	pb "github.com/hieuus/home-services/pb"

	ll "github.com/hieuus/home-services/pkg/log"
	"log"

	"net/http"
	"os"
)

func run(_ []string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	l := ll.New()
	cfg := config.Load()
	grpcMux := runtime.NewServeMux()

	service := services.New(cfg)
	if err := registerWithMuxServer(ctx, grpcMux, service); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	httpServer := &http.Server{
		Addr:    cfg.Server.Http.String(),
		Handler: mux,
	}

	l.Info().Msgf("start HTTP gateway server at %s", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func registerWithMuxServer(ctx context.Context, grpcMux *runtime.ServeMux, service *services.Service) error {
	if err := pb.RegisterHomeServicesAdminServiceHandlerServer(ctx, grpcMux, service); err != nil {
		return err
	}
	if err := pb.RegisterHomeServicesUserServiceHandlerServer(ctx, grpcMux, service); err != nil {
		return err
	}
	registerWithRestfullHandler(grpcMux)
	return nil
}

func registerWithRestfullHandler(grpcMux *runtime.ServeMux) {
	_ = grpcMux.HandlePath(http.MethodGet, "/health", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Health OK"))
	})
}
