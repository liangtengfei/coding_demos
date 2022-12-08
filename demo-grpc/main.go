package main

import (
	"demo-grpc/gapi"
	"demo-grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var DB_SOURCE string

func main() {
	grpcServerRun()
}

func grpcServerRun() {
	server, err := gapi.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("启动gRPC服务：%s", "127.0.0.1:8088")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("启动gRPC服务失败")
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
