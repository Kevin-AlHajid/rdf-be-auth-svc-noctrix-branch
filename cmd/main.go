package cmd

import (
	"log"
	"net"
	"os"

	"github.com/Kevin-AlHajid/rdf-be-auth-svc-noctrix-branch/domain"
	"github.com/Kevin-AlHajid/rdf-be-auth-svc-noctrix-branch/infrastructure"
	"github.com/Kevin-AlHajid/rdf-be-auth-svc-noctrix-branch/proto"

	"github.com/Kevin-AlHajid/rdf-be-auth-svc-noctrix-branch/repository"
	"github.com/Kevin-AlHajid/rdf-be-auth-svc-noctrix-branch/usecase"
	gogrpc "google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	db := infrastructure.ConnectDB()

	if err := db.AutoMigrate(&domain.Role{}); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	roleRepo := repository.NewRoleRepository(db)
	roleUC := usecase.NewRoleUsecase(roleRepo)
	//roleGRPC := grpc.NewRoleGRPC(roleUC)

	gRPCServer := gogrpc.NewServer()
	proto.RegisterRoleServiceServer(gRPCServer, roleGRPC)

	log.Println("gRPC server is running on port 9000")
	err = gRPCServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
