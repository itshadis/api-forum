package posts

import (
	"context"

	"github.com/itshadis/api-forum/internal/models/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostById(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostById(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get post from database")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikeByPostId(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get like count from database")
		return nil, err
	}
	comments, err := s.postRepo.GetCommentByPostId(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get comments from database")
		return nil, err
	}

	return &posts.GetPostResponse{
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
		Comments:  comments,
	}, nil
}
