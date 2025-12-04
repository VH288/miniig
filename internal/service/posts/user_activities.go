package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/VH288/miniig/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {
	now := time.Now()
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	userActivity, err := s.postRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error get user activity from database")
		return err
	}

	if userActivity == nil {
		if !request.IsLiked {
			return errors.New("you havent liked before")
		} else {
			err = s.postRepo.CreateUserActivity(ctx, model)
			if err != nil {
				log.Error().Err(err).Msg("error create user activity to database")
				return err
			}
		}
	} else {
		if userActivity.IsLiked == model.IsLiked {
			if model.IsLiked {
				return errors.New("status is liked before")
			} else {
				return errors.New("status is disliked before")
			}
		}
		err = s.postRepo.UpdateUserActivity(ctx, model)
		if err != nil {
			log.Error().Err(err).Msg("error update user activity to database")
			return err
		}
	}
	return nil
}
