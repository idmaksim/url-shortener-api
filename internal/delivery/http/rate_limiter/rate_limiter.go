package rateLimiter

import (
	"time"

	"github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"
)

type IPRateLimiter struct {
	rps   rate.Limit
	burst int
	c     *cache.Cache
}

func NewIPRateLimiter(rps rate.Limit, burst int) *IPRateLimiter {
	return &IPRateLimiter{
		rps:   rps,
		burst: burst,
		c:     cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	item, found := i.c.Get(ip)
	if found {
		return item.(*rate.Limiter)
	}

	limiter := rate.NewLimiter(i.rps, i.burst)
	i.c.Set(ip, limiter, cache.DefaultExpiration)
	return limiter
}
