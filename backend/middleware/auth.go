package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"library/backend/apiutil"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			apiutil.WriteError(w, http.StatusUnauthorized, "No token provided")
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			apiutil.WriteError(w, http.StatusUnauthorized, "Invalid token format")
			return
		}

		token, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(apiutil.JWTSecret()), nil
		})
		if err != nil || !token.Valid {
			log.Println("Invalid token:", err)
			apiutil.WriteError(w, http.StatusForbidden, "Invalid token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			apiutil.WriteError(w, http.StatusUnauthorized, "Invalid token payload")
			return
		}

		userIDFloat, ok := claims["userId"].(float64)
		if !ok {
			apiutil.WriteError(w, http.StatusUnauthorized, "Invalid token payload")
			return
		}

		userID := int(userIDFloat)
		role, _ := claims["role"].(string)
		if role == "" {
			role = "user"
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		ctx = context.WithValue(ctx, "role", role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, _ := r.Context().Value("role").(string)
		if role != "admin" {
			apiutil.WriteError(w, http.StatusForbidden, "Admin required")
			return
		}
		next.ServeHTTP(w, r)
	})
}
