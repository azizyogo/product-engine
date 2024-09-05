package middleware

import (
	"context"
	"net/http"
	"product-engine/common/auth"
	constanta "product-engine/common/const"
	"product-engine/common/utils/response"
	"product-engine/config"
	"strings"
	"time"
)

func Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			response.WriteErrorResponse(w, http.StatusUnauthorized, "Missing auth token")
			return
		}

		tokenParts := strings.Split(tokenHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			response.WriteErrorResponse(w, http.StatusUnauthorized, "Invalid auth token")
			return
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			response.WriteErrorResponse(w, http.StatusInternalServerError, "Internal; Server Error")
			return
		}
		claims, err := auth.GetJWTClaims(tokenParts[1], cfg.JWT.Secret)
		if err != nil {
			response.WriteErrorResponse(w, http.StatusUnauthorized, "Invalid auth token")
			return
		}

		expirationTime := time.Unix(claims.ExpiresAt.Unix(), 0)
		if expirationTime.Before(time.Now()) {
			response.WriteErrorResponse(w, http.StatusUnauthorized, "Token Expired")
			return
		}

		ctx := context.WithValue(r.Context(), constanta.CLAIMS_CONTEXT_KEY, claims)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
