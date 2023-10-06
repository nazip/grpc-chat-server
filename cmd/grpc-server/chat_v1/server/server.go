package server

import (
	"context"

	"github.com/brianvoe/gofakeit"

	desc "github.com/nazip/grpc-chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	desc.UnimplementedChatV1Server
}

// NewServer returns *server
func NewServer() *server {
	return &server{}
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id := gofakeit.Int64()

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
