package server

import (
	"log"
	"net"

	grpcAPI "github.com/Archisman-Mridha/blockchain-from-scratch/api/grpc"
	"github.com/Archisman-Mridha/blockchain-from-scratch/api/grpc/proto/generated"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	generated.RegisterNodeServer(server, grpcAPI.NewNode())

	address := ":3000"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed creating TCP listener at %s", address)
	}

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Server error occurred : %v", err)
	}
}
