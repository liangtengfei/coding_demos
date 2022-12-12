package gapi

import (
	"context"
	"demo-grpc/pb"
	"demo-grpc/utils"
	"log"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
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

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := utils.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := utils.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	return violations
}

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	violations := validateLoginUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
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

func validateLoginUserRequest(req *pb.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := utils.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	return violations
}
