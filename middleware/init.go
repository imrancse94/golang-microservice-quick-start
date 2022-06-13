package middleware

import (
	"go.quick.start/api"
	"golang.org/x/time/rate"
	"net/http"
)

type RateLimiterMiddleware struct {
	Name        string
	Description string
}

// Handle set a limit of request allowed in a specific time
func (RateLimiterMiddleware) Handle(next http.Handler) http.Handler {
	var limiter = rate.NewLimiter(1, 3)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		responseData := api.Response{
			Status:  "E001",
			Message: "Too many requests",
			Data:    "",
		}
		api.ErrorResponse(responseData, w) //http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
		//next.ServeHTTP(w, r)
	})
}

// GetName returns the middleware name
func (m RateLimiterMiddleware) GetName() string {
	return m.Name
}

// GetDescription returns the middleware description
func (m RateLimiterMiddleware) GetDescription() string {
	return m.Description
}

func NewRateLimiterMiddleware() RateLimiterMiddleware {
	return RateLimiterMiddleware{
		Name:        "RateLimiter",
		Description: "Provides rate limit over HTTP requests",
	}
}
