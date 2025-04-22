package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/services"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// AuthContext keys
type contextKey string

const (
	UserIDKey contextKey = "userID"
	RoleKey   contextKey = "role"
)

// JWTAuth middleware validates the JWT token in the Authorization header
func JWTAuth(jwtSecret []byte, roleService *services.RoleService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			// Check that it's a Bearer token
			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
				return
			}

			// Extract the token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Parse and validate the token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Validate the signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return jwtSecret, nil
			})

			if err != nil {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			// Check if the token is valid
			if !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Extract the claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			// Extract user ID
			userIDStr, ok := claims["sub"].(string)
			if !ok {
				http.Error(w, "Invalid token subject", http.StatusUnauthorized)
				return
			}

			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
				return
			}

			// Create a new context with the user ID
			ctx := context.WithValue(r.Context(), UserIDKey, userID)

			// Get user role (if role service is provided)
			if roleService != nil {
				role, err := roleService.GetUserRole(r.Context(), userID)
				if err == nil && role != nil {
					ctx = context.WithValue(ctx, RoleKey, role.Name)
				}
			}

			// Call the next handler with the new context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RequireRole middleware checks if the user has the required role
func RequireRole(requiredRoles []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get role from context
			role, ok := r.Context().Value(RoleKey).(string)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusForbidden)
				return
			}

			// Check if the user has one of the required roles
			hasRole := false
			for _, requiredRole := range requiredRoles {
				if role == requiredRole {
					hasRole = true
					break
				}
			}

			if !hasRole {
				http.Error(w, "Insufficient permissions", http.StatusForbidden)
				return
			}

			// User has the required role, proceed
			next.ServeHTTP(w, r)
		})
	}
}

// GetUserID extracts the user ID from the request context
func GetUserID(r *http.Request) (uuid.UUID, error) {
	userID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("user ID not found in context")
	}
	return userID, nil
}

// GetUserRole extracts the user role from the request context
func GetUserRole(r *http.Request) (string, error) {
	role, ok := r.Context().Value(RoleKey).(string)
	if !ok {
		return "", errors.New("role not found in context")
	}
	return role, nil
}
