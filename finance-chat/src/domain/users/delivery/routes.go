package users

import (
	"finance-chat/src/domain/users"
	"github.com/gin-gonic/gin"
)

func AddRoutes(engine *gin.Engine) {
	// Configure login route
	engine.PUT("/api/users/login", users.HandleLogin)

	// Configure sign up route
	engine.POST("/api/users", users.HandleSignup)
}