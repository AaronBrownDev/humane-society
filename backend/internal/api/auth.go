package api

import (
	"context"
	"encoding/json"
	"github.com/AaronBrownDev/HumaneSociety/internal/services"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// registerHandler processes user registration requests.
// It validates the input, registers the user via the auth service,
// and returns the created user ID on success.
func (a *api) registerHandler(w http.ResponseWriter, r *http.Request) {
	var req services.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	// Basic validation
	if req.FirstName == "" || req.LastName == "" || req.EmailAddress == "" ||
		req.PhysicalAddress == "" || req.MailingAddress == "" || req.Password == "" {
		a.respondError(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	// Call the auth service to register the user
	userAccount, err := a.authService.Register(ctx, req)
	if err != nil {
		a.logger.Printf("Error registering user: %v", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to register user")
		return
	}

	a.respondJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "User registered successfully",
		"userId":  userAccount.UserID,
	})
}

// loginHandler processes user login requests.
// It validates credentials via the auth service and returns an access token,
// setting a secure HTTP-only cookie with the refresh token.
func (a *api) loginHandler(w http.ResponseWriter, r *http.Request) {
	var req services.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	// Basic validation
	if req.Email == "" || req.Password == "" {
		a.respondError(w, http.StatusBadRequest, "Email and password are required")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	// Call the auth service to authenticate the user
	authResp, err := a.authService.Login(ctx, req)
	if err != nil {
		a.logger.Printf("Login failed: %v", err)
		a.respondError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Set refresh token as an HTTP-only cookie for security
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    authResp.RefreshToken,
		HttpOnly: true,
		Secure:   true, // Set to true in production with HTTPS
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 7), // 7 days
	})

	// Return only the access token in the response body
	a.respondJSON(w, http.StatusOK, map[string]interface{}{
		"accessToken": authResp.AccessToken,
		"expiresAt":   authResp.ExpiresAt,
		"userId":      authResp.UserID,
	})
}

// refreshTokenHandler processes token refresh requests.
// It validates the refresh token cookie and issues a new access token,
// rotating the refresh token for security.
func (a *api) refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Get refresh token from cookie
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, "Refresh token not found")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	// Call the auth service to refresh the token
	refreshReq := services.RefreshTokenRequest{
		RefreshTokenID: cookie.Value,
	}

	authResp, err := a.authService.RefreshAccessToken(ctx, refreshReq)
	if err != nil {
		a.logger.Printf("Token refresh failed: %v", err)
		a.respondError(w, http.StatusUnauthorized, "Invalid refresh token")
		return
	}

	// Update the refresh token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    authResp.RefreshToken,
		HttpOnly: true,
		Secure:   true, // Set to true in production with HTTPS
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 7), // 7 days
	})

	// Return only the access token in the response body
	a.respondJSON(w, http.StatusOK, map[string]interface{}{
		"accessToken": authResp.AccessToken,
		"expiresAt":   authResp.ExpiresAt,
		"userId":      authResp.UserID,
	})
}

// logoutHandler processes user logout requests.
// It clears the refresh token cookie and invalidates the token server-side.
func (a *api) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get refresh token from cookie to invalidate it server-side
	cookie, err := r.Cookie("refresh_token")
	if err == nil && cookie.Value != "" {
		// If we have a valid token, revoke it server-side
		tokenID, parseErr := uuid.Parse(cookie.Value)
		if parseErr == nil {
			ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
			defer cancel()

			if revokeErr := a.authService.RevokeToken(ctx, tokenID); revokeErr != nil {
				// Log the error but continue with logout
				a.logger.Printf("Error revoking token: %v", revokeErr)
			}
		}
	}

	// Clear the refresh token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   -1, // Delete the cookie
	})

	a.respondJSON(w, http.StatusOK, map[string]string{
		"message": "Logged out successfully",
	})
}
