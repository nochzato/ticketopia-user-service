package grpcapi

import (
	"net"

	db "github.com/nochzato/ticketopia-user-service/internal/db/sqlc"
	pb "github.com/nochzato/ticketopia-user-service/pkg/pb/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	grpcServer *grpc.Server
	queries    *db.Queries
}

func NewServer(queries *db.Queries) *Server {
	grpcServer := grpc.NewServer()

	server := &Server{
		grpcServer: grpcServer,
		queries:    queries,
	}

	pb.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	return server
}

func (s *Server) Start(listener net.Listener) error {
	return s.grpcServer.Serve(listener)
}
