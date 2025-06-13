package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"nota.shared/pagination"
	s "nota.shared/status"
	"nota.snippet/internal/converter"
	"nota.snippet/internal/repository"
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
	if req.Content == "" || req.Title == "" {
		return nil, status.Errorf(codes.InvalidArgument, "content and title are required")
	}

	languageHint := "text"
	if req.LanguageHint != nil {
		languageHint = *req.LanguageHint
	}

	isPublic := false
	if req.IsPublic != nil {
		isPublic = *req.IsPublic
	}

	tags := []string{}
	if req.Tags != nil {
		tags = req.Tags
	}

	id, err := h.service.Snippet.Create(ctx, req.Title, req.Content, languageHint, isPublic, tags)
	if err != nil {
		return nil, s.StatusInternal
	}

	return &pb.CreateSnippetResponse{
		Id: id.String(),
	}, nil
}

func (h *SnippetServiceHandler) GetSnippet(
	ctx context.Context,
	req *pb.GetSnippetRequest,
) (*pb.GetSnippetResponse, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id is required")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	snippet, err := h.service.Snippet.GetSnippet(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrSnippetNotFound):
			return nil, status.Errorf(codes.NotFound, "snippet not found")
		default:
			return nil, s.StatusInternal
		}
	}

	return &pb.GetSnippetResponse{
		Snippet: converter.ToPbSnippet(snippet),
	}, nil
}

func (h *SnippetServiceHandler) ListMySnippets(
	ctx context.Context,
	req *pb.ListMySnippetsRequest,
) (*pb.ListMySnippetsResponse, error) {
	page := 1
	if req.Page != nil {
		page = int(*req.Page)
	}

	size := 10
	if req.Size != nil {
		size = int(*req.Size)
	}

	pagination := pagination.Pagination{
		Page: page,
		Size: size,
	}

	snippets, total, err := h.service.Snippet.ListByOwnerID(ctx, pagination)
	if err != nil {
		return nil, s.StatusInternal
	}

	snippetsPb := make([]*pb.Snippet, len(snippets))
	for i, snippet := range snippets {
		snippetsPb[i] = converter.ToPbSnippet(snippet)
	}
	return &pb.ListMySnippetsResponse{
		Snippets:   snippetsPb,
		Pagination: converter.ToPbPagination(pagination, total),
	}, nil
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
