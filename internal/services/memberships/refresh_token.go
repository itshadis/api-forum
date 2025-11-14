package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/itshadis/api-forum/internal/models/memberships"
	"github.com/itshadis/api-forum/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefershToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get refresh token from database")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has expired")
	}

	// missmastch between token database and token request
	if existingRefreshToken.RefreshToken != request.Token {
		return "", errors.New("invalid refresh token")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exist")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}

	return token, nil
}
