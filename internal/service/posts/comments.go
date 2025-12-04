package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/VH288/miniig/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreateComment(ctx context.Context, postID, userID int64, req posts.CreateCommentRequest) error {
	now := time.Now()
	model := posts.CommentModel{
		PostID:         postID,
		UserID:         userID,
		CommentContent: req.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userID, 10),
		UpdatedBy:      strconv.FormatInt(userID, 10),
	}

	err := s.postRepo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error create post to repository")
		return err
	}
	return nil
}
