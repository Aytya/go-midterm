package handler

import (
	"context"
	"errors"
	"myproject/backend/domain"
	"net/http"
)

func (a *App) AuthMiddleware(tokenString string) (context.Context, error) {
	if tokenString == "" {
		return nil, errors.New("authorization header missing")
	}

	claims, err := a.ValidateJWT(tokenString)
	if err != nil {
		return nil, err
	}

	userRole := claims["role"].(string)
	ctx := context.WithValue(context.Background(), "user", claims["user"])
	ctx = context.WithValue(ctx, "role", userRole)

	return ctx, nil
}

func (a *App) RolePermissionMiddleware(requiredPermission string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole, ok := r.Context().Value("role").(string)
			if !ok || !CheckPermission(userRole, requiredPermission) {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func CheckPermission(role string, permission string) bool {
	r, exists := domain.Roles[role]
	if !exists {
		return false
	}

	for _, v := range r.Permissions {
		if v == permission {
			return true
		}
	}
	return false
}
