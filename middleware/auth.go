package middleware

import (
	"context"
	"go.quick.start/apiresponse"
	"net/http"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "requestTime", time.Now().Format(time.RFC3339))
		r = r.WithContext(ctx)
		apiresponse.ErrorResponse(http.StatusUnauthorized, "E101", map[string]string{}, "Unauthorized", w)
		//next.ServeHTTP(w, r)
	})
}
