package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarcetic/gorchase/server/accounts"
)

// InitRouter register accounts API routes
func InitRouter(r *gin.RouterGroup, svc accounts.Service) *gin.RouterGroup {

	handler := NewHandler(svc)

	accounts := r.Group("/accounts")
	// Public
	accounts.PUT("/register", handler.Singup)
	accounts.POST("/login", AuthMiddleware(svc).LoginHandler)
	// Protected
	accounts.GET("/me", AuthMiddleware(svc).MiddlewareFunc(), handler.Get)
	accounts.GET("/refresh_token", AuthMiddleware(svc).RefreshHandler)

	return r
}
