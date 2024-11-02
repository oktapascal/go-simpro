package auth

import (
	"context"
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oktapascal/go-simpro/config"
	"github.com/oktapascal/go-simpro/exception"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
	"strings"
	"time"
)

type Service struct {
	rpo  model.AuthSessionRepository
	urpo model.UserRepository
	prpo model.PermissionRepository
	db   *sql.DB
}

func (svc *Service) Login(ctx context.Context, request *model.LoginRequest, userAgent string) model.TokenResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	var identifier string
	if strings.Contains(request.Identifier, "@") {
		identifier = "email"
	} else {
		identifier = "username"
	}

	var user *model.User
	if identifier == "email" {
		user, err = svc.urpo.GetUserByEmail(ctx, tx, request.Identifier)
	} else {
		user, err = svc.urpo.GetUserByUsername(ctx, tx, request.Identifier)
	}

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	checkPassword := helper.CheckHash(request.Password, user.Password)
	if !checkPassword {
		panic(exception.NewNotMatchedError("password is not matched"))
	}

	permissionsByID := svc.prpo.GetRolePermissionsByID(ctx, tx, user.IDRole)

	var permissions []int

	for _, value := range *permissionsByID {
		permissions = append(permissions, value.IDPermission)
	}

	jwtParams := config.JwtParameters{
		Id:          user.ID,
		Username:    user.Username,
		GroupMenu:   user.IDMenuGroup,
		Role:        user.IDRole,
		Permissions: permissions,
	}

	accessToken, expiresAccessToken, err := helper.GenerateAccessToken(&jwtParams)
	if err != nil {
		panic(err)
	}

	refreshToken, expiresRefreshToken, err := helper.GenerateRefreshToken(&jwtParams)
	if err != nil {
		panic(err)
	}

	session := new(model.AuthSession)
	session.UserId = user.ID
	session.RefreshToken = refreshToken
	session.UserAgent = userAgent

	unixFormat := expiresRefreshToken
	t := time.Unix(unixFormat, 0)
	formattedDateTime := t.Format("2006-01-02 15:04:05")
	expiresAt, _ := time.Parse("2006-01-02 15:04:05", formattedDateTime)
	session.ExpiresAt = expiresAt

	svc.rpo.SaveSession(ctx, tx, session)

	accessTokenModel := model.AccessToken{
		Token:     accessToken,
		ExpiresAt: expiresAccessToken,
	}

	refreshTokenModel := model.RefreshToken{
		Token:     refreshToken,
		ExpiresAt: expiresRefreshToken,
	}

	return model.TokenResponse{
		AccessToken:  accessTokenModel,
		RefreshToken: refreshTokenModel,
	}
}

func (svc *Service) Logout(ctx context.Context, claims jwt.MapClaims, userAgent string) {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	userID, ok := claims["id"].(string)
	if !ok {
		panic("Something wrong when extracting username from jwt token")
	}

	svc.rpo.RevokeSession(ctx, tx, userID, userAgent)
}

func (svc *Service) GetAccessToken(ctx context.Context, claims jwt.MapClaims, userAgent string) model.TokenResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	userID, ok := claims["id"].(string)
	if !ok {
		panic("Something wrong when extracting username from jwt token")
	}

	checkRefreshToken, err := svc.rpo.CheckRefreshToken(ctx, tx, userID, userAgent)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user, err := svc.urpo.GetUserByID(ctx, tx, userID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	permissions, ok := claims["permissions"].([]int)

	jwtParams := config.JwtParameters{
		Id:          user.ID,
		Username:    user.Username,
		GroupMenu:   user.IDMenuGroup,
		Role:        user.IDRole,
		Permissions: permissions,
	}

	accessToken, expiresAccessToken, err := helper.GenerateAccessToken(&jwtParams)
	if err != nil {
		panic(err)
	}

	accessTokenModel := model.AccessToken{
		Token:     accessToken,
		ExpiresAt: expiresAccessToken,
	}

	refreshTokenModel := model.RefreshToken{
		Token:     checkRefreshToken.RefreshToken,
		ExpiresAt: checkRefreshToken.ExpiresAt.Unix(),
	}

	return model.TokenResponse{
		AccessToken:  accessTokenModel,
		RefreshToken: refreshTokenModel,
	}
}
