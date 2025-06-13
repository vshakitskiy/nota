package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"nota.snippet/internal/model"
	pb "nota.snippet/pkg/pb/v1"
)

func ToPbSnippet(snippet *model.Snippet) *pb.Snippet {
	return &pb.Snippet{
		Id:           snippet.ID.String(),
		Title:        snippet.Title,
		Content:      snippet.Content,
		LanguageHint: snippet.LanguageHint,
		IsPublic:     snippet.IsPublic,
		Tags:         snippet.Tags,
		CreatedAt:    timestamppb.New(snippet.CreatedAt),
		UpdatedAt:    timestamppb.New(snippet.UpdatedAt),
	}
}
