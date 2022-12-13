package main

import (
	"context"
	"demo-grpc/pb"
	"flag"
	"time"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddress := flag.String("address", "0.0.0.0:8080", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc1, err := grpc.Dial(*serverAddress, transportOption)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot dial server")
	}

	service := pb.NewUserServiceClient(cc1)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := service.LoginUser(ctx, &pb.LoginUserRequest{
		Username: "admin",
		Password: "123456",
	})
	if err != nil {
		log.Fatal().Err(err).Msg("访问登陆接口失败")
	}

	log.Info().Str("用户名", res.User.Username).Msg("请求结果")
}
