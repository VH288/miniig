package memberships

import (
	"net/http"

	"github.com/VH288/miniig/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")
	accessToken, err := h.membershipSvc.ValidateRefreshToken(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	response := memberships.RefreshResponse{
		AccessToken: accessToken,
	}
	c.JSON(http.StatusOK, response)
}
