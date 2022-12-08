package gapi

import "demo-grpc/pb"

type Server struct {
	pb.UnimplementedUserServiceServer
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}
