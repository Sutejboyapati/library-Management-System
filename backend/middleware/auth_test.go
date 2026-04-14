package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"library/backend/apiutil"

	"github.com/golang-jwt/jwt/v5"
)

func signedToken(t *testing.T, role string) string {
	t.Helper()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": 11,
		"role":   role,
	})

	tokenString, err := token.SignedString([]byte(apiutil.JWTSecret()))
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}

	return tokenString
}

func TestVerifyTokenAcceptsDefaultSecretTokens(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
	req.Header.Set("Authorization", "Bearer "+signedToken(t, "user"))

	recorder := httptest.NewRecorder()
	handler := VerifyToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if userID, ok := r.Context().Value("userID").(int); !ok || userID != 11 {
			t.Fatalf("expected userID 11 in context, got %v", r.Context().Value("userID"))
		}
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}
}

func TestRequireAdminRejectsNonAdminUsers(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
	req = req.WithContext(context.WithValue(req.Context(), "role", "user"))

	recorder := httptest.NewRecorder()
	handler := RequireAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusForbidden {
		t.Fatalf("expected status 403, got %d", recorder.Code)
	}
}
