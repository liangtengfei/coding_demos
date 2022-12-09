package gapi

import (
	"context"
	"demo-grpc/pb"
	"fmt"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {
		return nil, fmt.Errorf("用户名或密码不能为空")
	}

	rsp := &pb.CreateUserResponse{
		User: &pb.UserMessage{
			Username:          req.Username,
			Email:             req.Email,
			PasswordChangedAt: timestamppb.New(time.Now()),
			CreatedAt:         timestamppb.New(time.Now()),
		},
	}

	return rsp, nil
}

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	if len(req.Password) == 0 || len(req.Username) == 0 {
		return nil, fmt.Errorf("用户名或密码不能为空")
	}
	meta := s.extractMetadata(ctx)
	rsp := &pb.LoginUserResponse{
		User: &pb.UserMessage{
			Username:          req.Username,
			Email:             "2022@gmail.com",
			PasswordChangedAt: timestamppb.New(time.Now()),
			CreatedAt:         timestamppb.New(time.Now()),
		},
		SessionId:             "123456789",
		AccessToken:           "access_token",
		RefreshToken:          "refresh_token",
		AccessTokenExpiresAt:  timestamppb.New(time.Now()),
		RefreshTokenExpiresAt: timestamppb.New(time.Now().Add(24 * time.Hour)),
	}
	log.Printf("客户端信息：%v", meta)
	return rsp, nil
}
