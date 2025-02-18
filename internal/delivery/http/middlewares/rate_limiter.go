package middlewares

import (
	"net/http"

	httpErrors "github.com/idmaksim/url-shortener-api/internal/delivery/http/errors"
	rateLimiter "github.com/idmaksim/url-shortener-api/internal/delivery/http/rate_limiter"
	"github.com/labstack/echo/v4"
)

func ThrottleMiddleware(limiter *rateLimiter.IPRateLimiter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ip := c.RealIP()

			lim := limiter.GetLimiter(ip)

			if !lim.Allow() {
				return c.JSON(http.StatusTooManyRequests, httpErrors.NewHTTPError("Too many requests", http.StatusTooManyRequests))
			}

			return next(c)
		}
	}
}
