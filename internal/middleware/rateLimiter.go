package middleware

import (
	"fmt"
	"github.com/dot123/gin-gorm-admin/internal/config"
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
	"golang.org/x/time/rate"
	"strconv"
	"time"
)

func RateLimiterMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := config.C.RateLimiter
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	rc := config.C.Redis
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": rc.Addr,
		},
		Password: rc.Password,
		DB:       cfg.RedisDB,
	})

	limiter := redis_rate.NewLimiter(ring)
	limiter.Fallback = rate.NewLimiter(rate.Inf, 0)

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userID := contextx.FromUserID(c.Request.Context())
		if userID != 0 {
			limit := cfg.Count
			count, delay, allowed := limiter.AllowMinute(fmt.Sprintf("%d", userID), limit)
			if !allowed {
				h := c.Writer.Header()
				h.Set("X-RateLimit-Limit", strconv.FormatInt(limit, 10))
				h.Set("X-RateLimit-Remaining", strconv.FormatInt(limit-count, 10))
				delaySec := int64(delay / time.Second)
				h.Set("X-RateLimit-Delay", strconv.FormatInt(delaySec, 10))
				ginx.ResError(c, errors.ErrTooManyRequests)
				return
			}
		}

		c.Next()
	}
}
