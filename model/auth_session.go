package model

import (
	"context"
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type (
	AuthSession struct {
		Id           string
		UserId       string
		RefreshToken string
		Revoked      bool
		UserAgent    string
		ExpiresAt    time.Time
	}

	AccessToken struct {
		Token     string `json:"token"`
		ExpiresAt int64  `json:"expires_at"`
	}

	RefreshToken struct {
		Token     string `json:"token"`
		ExpiresAt int64  `json:"expires_at"`
	}

	TokenResponse struct {
		AccessToken  AccessToken  `json:"access_token"`
		RefreshToken RefreshToken `json:"refresh_token"`
	}

	LoginRequest struct {
		Identifier string `validate:"required" json:"identifier"`
		Password   string `validate:"required" json:"password"`
	}

	AuthSessionRepository interface {
		SaveSession(ctx context.Context, tx *sql.Tx, data *AuthSession)
		RevokeSession(ctx context.Context, tx *sql.Tx, userID string, userAgent string)
		CheckRefreshToken(ctx context.Context, tx *sql.Tx, userID string, userAgent string) (*AuthSession, error)
	}

	AuthSessionService interface {
		Login(ctx context.Context, request *LoginRequest, userAgent string) TokenResponse
		Logout(ctx context.Context, claims jwt.MapClaims, userAgent string)
		GetAccessToken(ctx context.Context, claims jwt.MapClaims, userAgent string) TokenResponse
	}

	AuthSessionHandler interface {
		Login() http.HandlerFunc
		Logout() http.HandlerFunc
		GetAccessToken() http.HandlerFunc
	}
)
