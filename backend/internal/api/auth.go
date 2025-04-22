// auth.go - Updated authentication API handlers

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
		a.logger.Printf("Invalid request format: %v", err)
		a.respondError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	// Basic validation
	if req.FirstName == "" || req.LastName == "" || req.EmailAddress == "" ||
		req.PhysicalAddress == "" || req.MailingAddress == "" || req.Password == "" {
		a.logger.Printf("Missing required fields in registration request")
		a.respondError(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	// Call the auth service to register the user
	userAccount, err := a.authService.Register(ctx, req)
	if err != nil {
		a.logger.Printf("Error registering user: %v", err)
		// Return more specific error messages
		if err.Error() == "email address already in use" {
			a.respondError(w, http.StatusConflict, "Email address already in use")
		} else {
			a.respondError(w, http.StatusInternalServerError, "Failed to register user")
		}
		return
	}

	a.logger.Printf("User registered successfully: %s", userAccount.UserID)
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
		a.logger.Printf("Invalid login request format: %v", err)
		a.respondError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	// Basic validation
	if req.Email == "" || req.Password == "" {
		a.logger.Printf("Missing email or password in login request")
		a.respondError(w, http.StatusBadRequest, "Email and password are required")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	a.logger.Printf("Login attempt for email: %s", req.Email)

	// Call the auth service to authenticate the user
	authResp, err := a.authService.Login(ctx, req)
	if err != nil {
		a.logger.Printf("Login failed: %v", err)

		// Return appropriate HTTP status based on the error
		switch err.Error() {
		case "invalid email":
			a.respondError(w, http.StatusUnauthorized, "Invalid email")
		case "no user account found":
			a.respondError(w, http.StatusUnauthorized, "User account not found")
		case "invalid password":
			a.respondError(w, http.StatusUnauthorized, "Invalid password")
		case "account is locked":
			a.respondError(w, http.StatusForbidden, "Account is locked")
		case "account is not active":
			a.respondError(w, http.StatusForbidden, "Account is not active")
		default:
			a.respondError(w, http.StatusInternalServerError, "Authentication failed")
		}
		return
	}

	// Set refresh token as an HTTP-only cookie for security
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    authResp.RefreshToken,
		HttpOnly: true,
		Secure:   true,                  // Set to true in production with HTTPS
		SameSite: http.SameSiteNoneMode, // Required for cross-origin requests
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 7), // 7 days
	})

	a.logger.Printf("Login successful for user: %s", authResp.UserID)

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
		a.logger.Printf("Refresh token not found in cookie: %v", err)
		a.respondError(w, http.StatusBadRequest, "Refresh token not found")
		return
	}

	a.logger.Printf("Token refresh request received with token ID: %s", cookie.Value)

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
		Secure:   true,                  // Set to true in production with HTTPS
		SameSite: http.SameSiteNoneMode, // Required for cross-origin requests
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 7), // 7 days
	})

	a.logger.Printf("Token refresh successful for user: %s", authResp.UserID)

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
			} else {
				a.logger.Printf("Successfully revoked token: %s", tokenID)
			}
		} else {
			a.logger.Printf("Error parsing token ID: %v", parseErr)
		}
	} else {
		a.logger.Printf("No refresh token cookie found for logout")
	}

	// Clear the refresh token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		Path:     "/",
		MaxAge:   -1, // Delete the cookie
	})

	a.logger.Printf("User logged out successfully")
	a.respondJSON(w, http.StatusOK, map[string]string{
		"message": "Logged out successfully",
	})
}
