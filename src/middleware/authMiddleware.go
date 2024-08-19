package middleware

import "github.com/gin-gonic/gin"

type AuthMiddleware struct {
}

func NewAuthMiddleware() AuthMiddleware {
	return AuthMiddleware{}
}

func (a *AuthMiddleware) AuthorizeRequest(c *gin.Context) {
	c.Next()
}
