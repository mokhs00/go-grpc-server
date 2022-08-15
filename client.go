package go_grpc_server

import (
	gogrpcserver "go-grpc-server/protos/apis/v1/content"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
)

//go:generate mockgen -package go_grpc_server -destination ./mock_client.go -mock_names ContentClient=MockContentClient -source protos/apis/v1/content/content_grpc.pb.go

const serviceConfig = `{"loadBalancingPolicy":"round_robin"}`

var (
	once sync.Once
	cli  gogrpcserver.ContentClient

	_ gogrpcserver.ContentClient = (*MockContentClient)(nil)
)

func GetContentClient(serviceHost string) gogrpcserver.ContentClient {
	once.Do(func() {
		conn, _ := grpc.Dial(
			serviceHost,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(serviceConfig),
		)

		cli = gogrpcserver.NewContentClient(conn)
	})

	return cli
}
