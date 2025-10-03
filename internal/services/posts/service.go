package posts

import (
	"context"

	"github.com/itshadis/api-forum/internal/configs"
	"github.com/itshadis/api-forum/internal/models/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
}

type service struct {
	cfg      *configs.Config
	postRepo postRepository
}

func NewService(cfg *configs.Config, postRepo postRepository) *service {
	return &service{
		cfg:      cfg,
		postRepo: postRepo,
	}
}
