package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/VeneLooool/drones-api/internal/app/api/v1/drones"
	drone_client "github.com/VeneLooool/drones-api/internal/clients/drone-client"
	"github.com/VeneLooool/drones-api/internal/config"
	"github.com/VeneLooool/drones-api/internal/handlers/external_drone_events"
	"github.com/VeneLooool/drones-api/internal/kafka/drone-events/publisher"
	"github.com/VeneLooool/drones-api/internal/kafka/external-drone-events/subscriber"
	pb "github.com/VeneLooool/drones-api/internal/pb/api/v1/drones"
	"github.com/VeneLooool/drones-api/internal/pkg/db"
	drones_repo "github.com/VeneLooool/drones-api/internal/repository/drones"
	drones_uc "github.com/VeneLooool/drones-api/internal/usecase/drones"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.New(ctx)
	if err != nil {
		log.Fatalf("failed to create new config: %s", err.Error())
	}

	go func() {
		if err := runGRPC(ctx, cfg); err != nil {
			log.Fatal(err)
		}
	}()

	if err := runHTTPGateway(ctx, cfg); err != nil {
		log.Fatal(err)
	}
}

func runGRPC(ctx context.Context, cfg *config.Config) error {
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	dbAdapter, err := db.New(ctx)
	if err != nil {
		return err
	}
	defer dbAdapter.Close(ctx)

	dronesServer, err := newServices(ctx, dbAdapter, cfg)
	if err != nil {
		return err
	}
	pb.RegisterDronesServer(grpcServer, dronesServer)

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	log.Printf("gRPC server listening on :%s\n", cfg.GrpcPort)
	if err = grpcServer.Serve(grpcListener); err != nil {
		return err
	}
	return nil
}

func runHTTPGateway(ctx context.Context, cfg *config.Config) error {
	mux := runtime.NewServeMux()
	err := pb.RegisterDronesHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%s", cfg.GrpcPort), []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		log.Fatalf("failed to register gateway: %s", err.Error())
	}

	// Serve Swagger JSON and Swagger UI
	fs := http.FileServer(http.Dir("./swagger-ui")) // директория со статикой UI
	http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", fs))

	// Serve Swagger JSON файл
	http.HandleFunc("/swagger/drones.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./internal/pb/api/v1/drones/drones.swagger.json")
	})

	withCORS := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			// Для preflight-запросов
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			h.ServeHTTP(w, r)
		})
	}

	// gRPC → REST mux
	http.Handle("/", withCORS(mux))

	log.Printf("HTTP gateway listening on :%s\n", cfg.HttpPort)
	if err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.HttpPort), nil); err != nil {
		return err
	}

	return nil
}

func newServices(ctx context.Context, dbAdapter db.DataBase, cfg *config.Config) (*drones.Implementation, error) {
	droneEventsPublisher := publisher.New(ctx, cfg.GetKafkaConfig())

	droneClient, err := drone_client.New(ctx, cfg.GetDroneClientConfig())
	if err != nil {
		return nil, err
	}

	dronesRepo := drones_repo.New(dbAdapter)
	dronesUC := drones_uc.New(dronesRepo, droneEventsPublisher, droneClient)

	newHandlers(ctx, dronesUC, cfg)

	return drones.NewService(dronesUC), nil
}

func newHandlers(ctx context.Context, dronesUC *drones_uc.UseCase, cfg *config.Config) {
	handler := external_drone_events.New(dronesUC)
	sub := subscriber.New(ctx, handler, cfg.GetKafkaConfig())
	sub.Subscribe(ctx)
}
