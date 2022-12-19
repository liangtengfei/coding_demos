package main

import (
	context "context"
	"log"
	"math/rand"
	"net"
	"quickstart3/pb"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type LaptopService struct {
	pb.UnimplementedLaptopServiceServer
}

// DetailLaptop implements pb.LaptopServiceServer
func (*LaptopService) DetailLaptop(ctx context.Context, req *pb.LaptopRequest) (*pb.Laptop, error) {
	log.Printf("请求参数V2：%s", req.Id)
	return randomLaptop()
}

func main() {
	log.Println("开始启动服务")
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	serverOptions := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(serverOptions...)
	pb.RegisterLaptopServiceServer(grpcServer, &LaptopService{})
	reflection.Register(grpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	log.Println("gRPC服务已启动")
}

func randomLaptop() (*pb.Laptop, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
	}
	brand := randomStringFromSet("Apple", "Dell", "Lenovo")
	laptop := &pb.Laptop{
		Id:        id.String(),
		Brand:     brand,
		Name:      randomLaptopName(brand),
		CreatedAt: ptypes.TimestampNow(),
	}
	return laptop, nil
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
	}
}
