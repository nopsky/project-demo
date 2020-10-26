/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/9/29 10:31
 */
package http

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.jiayunhui.com/service/api/pkg/logger"
)

func ginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		end = end.Local()

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			logger.Info(path,
				" status:", c.Writer.Status(),
				" method:", c.Request.Method,
				" path:", path,
				" query:", query,
				" ip:", c.ClientIP(),
				" user-agent:", c.Request.UserAgent(),
				" time:", end.Format(time.RFC3339),
				" latency:", latency.Seconds(),
			)
		}
	}
}
