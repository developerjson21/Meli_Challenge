package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)
//********** GENERA LOGS DE LOS REQUEST ************* |
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		time := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		fmt.Printf("method: [%s] path: %s   time: %v \n", method, path, time.Format("2006-01-02 15:04:05"))
		c.Next()
	}
}