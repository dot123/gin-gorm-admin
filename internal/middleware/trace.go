package middleware

import (
	"fmt"
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"os"
	"sync/atomic"
	"time"
)

var (
	version string
	incrNum uint64
	pid     = os.Getpid()
)

func NewTraceID() string {
	return fmt.Sprintf("trace-id-%d-%s-%d", os.Getpid(), time.Now().Format("2006.01.02.15.04.05.999"), atomic.AddUint64(&incrNum, 1))
}

func TraceMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		traceID := c.GetHeader("X-Request-Id")
		if traceID == "" {
			traceID = NewTraceID()
		}

		ctx := contextx.NewTraceID(c.Request.Context(), traceID)
		ctx = logger.NewTraceIDContext(ctx, traceID)
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set("X-Trace-Id", traceID)

		c.Next()
	}
}
