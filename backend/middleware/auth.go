package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
<<<<<<< HEAD
	"strings"

	"library/backend/apiutil"

=======
	"os"
	"strings"

>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
<<<<<<< HEAD
			apiutil.WriteError(w, http.StatusUnauthorized, "No token provided")
=======
			http.Error(w, "No token provided", http.StatusUnauthorized)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
			return
		}
		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
<<<<<<< HEAD
			apiutil.WriteError(w, http.StatusUnauthorized, "Invalid token format")
=======
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
			return
		}
		token, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
<<<<<<< HEAD
			return []byte(apiutil.JWTSecret()), nil
		})
		if err != nil || !token.Valid {
			log.Println("Invalid token:", err)
			apiutil.WriteError(w, http.StatusForbidden, "Invalid token")
=======
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			log.Println("Invalid token:", err)
			http.Error(w, "Invalid token", http.StatusForbidden)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
<<<<<<< HEAD
			apiutil.WriteError(w, http.StatusUnauthorized, "Invalid token payload")
=======
			http.Error(w, "Invalid token payload", http.StatusUnauthorized)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
			return
		}
		userIDFloat, ok := claims["userId"].(float64)
		if !ok {
<<<<<<< HEAD
			apiutil.WriteError(w, http.StatusUnauthorized, "Invalid token payload")
=======
			http.Error(w, "Invalid token payload", http.StatusUnauthorized)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
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
<<<<<<< HEAD
			apiutil.WriteError(w, http.StatusForbidden, "Admin required")
=======
			http.Error(w, "Admin required", http.StatusForbidden)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
			return
		}
		next.ServeHTTP(w, r)
	})
}
