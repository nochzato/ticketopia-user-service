package main

import (
	"context"
	"log"
	"net"

	"github.com/jackc/pgx/v5"
	"github.com/nochzato/ticketopia-user-service/internal/config"
	dbmigration "github.com/nochzato/ticketopia-user-service/internal/db/migration"
	db "github.com/nochzato/ticketopia-user-service/internal/db/sqlc"
	"github.com/nochzato/ticketopia-user-service/internal/grpcapi"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot read config file: %s", err)
	}

	conn, err := pgx.Connect(context.Background(), config.Database.URL)
	if err != nil {
		log.Fatalf("cannot connect to the db: %s", err)
	}

	if err = dbmigration.RunMigrations(config); err != nil {
		log.Fatalf("failed to run migrations: %s", err)
	}

	queries := db.New(conn)

	grpcServer := grpcapi.NewServer(queries)

	lis, err := net.Listen("tcp", config.GRPCServer.Addr)
	if err != nil {
		log.Fatalf("failed to create a listener: %s", err)
	}

	log.Printf("starting grpc server at %s\n", lis.Addr().String())
	if err = grpcServer.Start(lis); err != nil && err != grpc.ErrServerStopped {
		log.Fatalf("grpc server crashed: %s", err)
	}
}
