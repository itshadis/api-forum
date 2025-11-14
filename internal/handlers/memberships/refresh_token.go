package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itshadis/api-forum/internal/models/memberships"
)

func (h *Handler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")
	accessToken, err := h.membershipSvc.ValidateRefershToken(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, memberships.RefreshTokenResponse{
		AccessToken: accessToken,
	})
}
