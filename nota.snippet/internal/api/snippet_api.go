package api

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
	"nota.snippet/internal/service"
	pb "nota.snippet/pkg/pb/v1"
)

type SnippetServiceHandler struct {
	pb.UnimplementedSnippetServiceServer
	service *service.Service
}

func NewSnippetServiceHandler(service *service.Service) *SnippetServiceHandler {
	return &SnippetServiceHandler{service: service}
}

func (h *SnippetServiceHandler) CreateSnippet(
	ctx context.Context,
	req *pb.CreateSnippetRequest,
) (*pb.CreateSnippetResponse, error) {
	panic("not implemented")
}

func (h *SnippetServiceHandler) GetSnippet(
	ctx context.Context,
	req *pb.GetSnippetRequest,
) (*pb.GetSnippetResponse, error) {
	panic("not implemented")
}

func (h *SnippetServiceHandler) ListMySnippets(
	ctx context.Context,
	req *pb.ListMySnippetsRequest,
) (*pb.ListMySnippetsResponse, error) {
	panic("not implemented")
}

func (h *SnippetServiceHandler) ListPublicSnippets(
	ctx context.Context,
	req *pb.ListPublicSnippetsRequest,
) (*pb.ListPublicSnippetsResponse, error) {
	fmt.Println("ListPublicSnippets")
	panic("not implemented")
}

func (h *SnippetServiceHandler) UpdateSnippet(
	ctx context.Context,
	req *pb.UpdateSnippetRequest,
) (*pb.UpdateSnippetResponse, error) {
	panic("not implemented")
}

func (h *SnippetServiceHandler) DeleteSnippet(
	ctx context.Context,
	req *pb.DeleteSnippetRequest,
) (*emptypb.Empty, error) {
	panic("not implemented")
}
