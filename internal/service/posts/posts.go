package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/VH288/miniig/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error {
	postHashtags := strings.Join(req.PostHashtags, ",")

	now := time.Now()
	model := posts.PostModel{
		UserID:       userID,
		PostTitle:    req.PostTitle,
		PostContent:  req.PostContent,
		PostHashtags: postHashtags,
		CreatedAt:    now,
		UpdatedAt:    now,
		CreatedBy:    strconv.FormatInt(userID, 10),
		UpdatedBy:    strconv.FormatInt(userID, 10),
	}

	err := s.postRepo.CreatePost(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error create post to repository")
		return err
	}
	return nil
}

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error) {
	limit := pageSize
	offset := pageSize * (pageIndex - 1)
	post, err := s.postRepo.GetAllPosts(ctx, limit, offset)
	if err != nil {
		return posts.GetAllPostResponse{}, err
	}
	response := posts.GetAllPostResponse{
		Data: post,
		Pagination: posts.Pagination{
			Limit:  limit,
			Offset: offset,
		},
	}
	return response, nil
}
