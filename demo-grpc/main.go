package main

import (
	"context"
	"demo-grpc/gapi"
	"demo-grpc/pb"
	"net"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"

	_ "demo-grpc/doc/statik"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

var DB_SOURCE string

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// 启动HTTP网关服务，支持普通http请求
	go grpcGatewayServerRun()

	// 启动gRPC服务
	grpcServerRun()
}

func grpcGatewayServerRun() {
	server, err := gapi.NewServer()
	if err != nil {
		log.Fatal().Err(err).Msg("启动gRPC网关错误")
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			// 使用小写驼峰字段名
			UseProtoNames: false,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterUserServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("注册服务处理失败")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("创建statikFS失败")
	}
	// fs := http.FileServer(http.Dir("./doc/swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(statikFS)))

	listener, err := net.Listen("tcp", "127.0.0.1:8089")
	if err != nil {
		log.Fatal().Err(err)
	}

	log.Printf("启动 HTTP网关服务：%s", "127.0.0.1:8089")
	err = http.Serve(listener, gapi.HttpLogger(mux))
	if err != nil {
		log.Fatal().Err(err).Msg("启动 HTTP 网关服务失败")
	}
}

func grpcServerRun() {
	server, err := gapi.NewServer()
	if err != nil {
		log.Fatal().Err(err)
	}

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		log.Fatal().Err(err)
	}

	log.Printf("启动 gRPC服务：%s", "127.0.0.1:8088")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("启动gRPC服务失败")
	}
}

// func ginServerRun() {
// 	r := gin.Default()

// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	flag.StringVar(&DB_SOURCE, "DB_SOURCE", "postgresql://postgres:Xiaohuozhi2022.@localhost:5432/postgres_api?sslmode=disable", "数据库连接")
// 	flag.Parse()

// 	conn, err := pgx.Connect(context.Background(), DB_SOURCE)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close(context.Background())

// 	var greeting string
// 	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 		os.Exit(1)
// 	}

// 	r.GET("/hello", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": greeting,
// 		})
// 	})

// 	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }
