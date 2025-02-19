package middlewares

import (
	"net/http"

	rateLimiter "github.com/idmaksim/url-shortener-api/internal/delivery/http/rate_limiter"
	"github.com/idmaksim/url-shortener-api/internal/errors"
	"github.com/labstack/echo/v4"
)

func ThrottleMiddleware(limiter *rateLimiter.IPRateLimiter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ip := c.RealIP()

			lim := limiter.GetLimiter(ip)

			if !lim.Allow() {
				return errors.NewHttpError(
					http.StatusTooManyRequests,
					errors.ErrCodeTooManyRequests,
					"Too many requests",
					nil,
				)
			}

			return next(c)
		}
	}
}
