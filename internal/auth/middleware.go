package auth

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const (
	UserIDKey   contextKey = "user_id"
	IsMemberKey contextKey = "is_member"
)

type Middleware struct {
	jwt *JWTManager
}

func NewMiddleware(jwt *JWTManager) *Middleware {
	return &Middleware{jwt: jwt}
}

// RequireAuth 인증 필수 미들웨어
func (m *Middleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := m.extractClaims(r)
		if err != nil {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, IsMemberKey, claims.IsMember)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireMember 회원 전용 미들웨어
func (m *Middleware) RequireMember(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := m.extractClaims(r)
		if err != nil {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
			return
		}

		if !claims.IsMember {
			http.Error(w, `{"error":"membership required"}`, http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, IsMemberKey, claims.IsMember)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// OptionalAuth 인증 선택적 미들웨어 (비회원도 접근 가능하지만 토큰이 있으면 파싱)
func (m *Middleware) OptionalAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := m.extractClaims(r)
		if err == nil {
			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, IsMemberKey, claims.IsMember)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) extractClaims(r *http.Request) (*Claims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, ErrInvalidToken
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, ErrInvalidToken
	}

	return m.jwt.ValidateAccessToken(parts[1])
}

// GetUserID 컨텍스트에서 사용자 ID 추출
func GetUserID(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDKey).(int64)
	return userID, ok
}

// IsMember 컨텍스트에서 회원 여부 추출
func IsMember(ctx context.Context) bool {
	isMember, ok := ctx.Value(IsMemberKey).(bool)
	return ok && isMember
}
