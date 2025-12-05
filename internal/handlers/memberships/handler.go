// Package memberships
package memberships

import (
	"context"

	"github.com/VH288/miniig/internal/middleware"
	"github.com/VH288/miniig/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine

	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")

	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)

	route.Use(middleware.AuthMiddleware())
	route.GET("/ping", h.Ping)

	routeRefresh := h.Group("memberships")
	routeRefresh.Use(middleware.AuthRefreshMiddleware())
	routeRefresh.POST("/refresh", h.RefreshToken)
}
