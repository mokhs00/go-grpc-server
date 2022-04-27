package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-grpc-server/config"
	"go-grpc-server/protos/apis/v1/content"
	"go-grpc-server/server/db"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ContentServer struct {
	content.ContentServer
}

func (s *ContentServer) ListContent(ctx context.Context, req *content.ListContentRequest) (*content.ListContentResponse, error) {
	return &content.ListContentResponse{Contents: []*content.ContentMessage{
		{
			ContentId: "1",
			Title:     "title",
			Body:      "body",
		},
	}}, nil
}

func main() {
	setting := config.NewSetting()

	db := db.MustGetDB(db.DBSetting{
		Host:     setting.DBHost,
		Port:     setting.DBPort,
		User:     setting.DBUser,
		Password: setting.DBPassword,
		Name:     setting.DBName,
	})
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("db.Close() error: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", setting.GRPCServerPort))
	if err != nil {
		log.Fatalf("fail net.Listen error: %v", err)
	}
	grpcServer := grpc.NewServer()

	content.RegisterContentServer(grpcServer, &ContentServer{})

	log.Printf("start gRPC server port: %s", setting.GRPCServerPort)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail Serve error: %v", err)
	}
}
