package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/VH288/miniig/internal/model/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsetUserActivity(c *gin.Context) {
	ctx := c.Request.Context()
	var request posts.UserActivityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("postID not valid").Error(),
		})
		return
	}

	userID := c.GetInt64("userID")

	err = h.postSvc.UpsertUserActivity(ctx, postID, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}
