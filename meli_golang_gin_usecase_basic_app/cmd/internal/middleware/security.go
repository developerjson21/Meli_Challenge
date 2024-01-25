package middleware

import (
	"meli_golang_gin_basic_app/cmd/handler"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// **************** AUTHENTICATION BASIC ******************** |
// ******** HEADER-TOKEN vs TOKEN .ENV ******************* |
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.JSON(http.StatusForbidden, handler.AuthFailedError("Token Not Found"))
			c.Abort()
			return
		}
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, handler.AuthUnauthorizedError("Token Invalid") )
			c.Abort()
			return
		}
		c.Next()
	}
}