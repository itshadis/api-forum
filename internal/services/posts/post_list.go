package posts

import (
	"context"

	"github.com/itshadis/api-forum/internal/models/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, PageIndex int) (posts.GetAllPostResponse, error) {
	limit := pageSize
	offset := pageSize * (PageIndex - 1)
	response, err := s.postRepo.GetAllPost(ctx, limit, offset)

	if err != nil {
		log.Error().Err(err).Msg("err get all post from database")
		return response, err
	}

	return response, nil
}
