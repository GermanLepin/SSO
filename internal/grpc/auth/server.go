package auth

import (
	"context"

	ssov1 "github.com/germanlepin/sso_protos/gen/go/sso"

	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	loginResponse := &ssov1.LoginResponse{
		Token: req.GetEmail(),
	}

	return loginResponse, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	registerResponse := &ssov1.RegisterResponse{
		UserId: 1112,
	}

	return registerResponse, nil
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	panic("emplement me")
}
