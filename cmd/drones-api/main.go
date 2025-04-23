package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/VeneLooool/drones-api/internal/app/api/v1/drones"
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

	go func() {
		if err := runGRPC(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err := runHTTPGateway(ctx); err != nil {
		log.Fatal(err)
	}
}

func runGRPC(ctx context.Context) error {
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	dbAdapter, err := db.New(ctx)
	if err != nil {
		return err
	}
	defer dbAdapter.Close(ctx)

	dronesServer, err := newServices(ctx, dbAdapter)
	if err != nil {
		return err
	}
	pb.RegisterDronesServer(grpcServer, dronesServer)

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	log.Println("gRPC server listening on :50051")
	if err = grpcServer.Serve(grpcListener); err != nil {
		return err
	}
	return nil
}

func runHTTPGateway(ctx context.Context) error {
	mux := runtime.NewServeMux()
	err := pb.RegisterDronesHandlerFromEndpoint(ctx, mux, "localhost:50051", []grpc.DialOption{
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

	// gRPC → REST mux
	http.Handle("/", mux)

	log.Println("HTTP gateway listening on :8080")
	if err = http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}

func newServices(ctx context.Context, dbAdapter db.DataBase) (*drones.Implementation, error) {
	dronesRepo := drones_repo.New(dbAdapter)
	dronesUC := drones_uc.New(dronesRepo)

	return drones.NewService(dronesUC), nil
}
