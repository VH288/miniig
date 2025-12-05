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

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id from database")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikePostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error count like by id from database")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get comment by id from database")
		return nil, err
	}

	response := posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:           postDetail.ID,
			UserID:       postDetail.UserID,
			Username:     postDetail.Username,
			PostTitle:    postDetail.PostTitle,
			PostContent:  postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked:      postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comment:   comments,
	}

	return &response, nil
}
