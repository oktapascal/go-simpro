package auth

import (
	"context"
	"database/sql"
	"errors"
	"github.com/oktapascal/go-simpro/model"
)

type Repository struct{}

func (rpo *Repository) SaveSession(ctx context.Context, tx *sql.Tx, data *model.AuthSession) {
	query := `insert into auth_session (user_id,refresh_token,user_agent,revoked,expired_at)
	values (?,?,?,false,?)`

	_, err := tx.ExecContext(ctx, query, data.UserId, data.RefreshToken, data.UserAgent, data.ExpiresAt)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) RevokeSession(ctx context.Context, tx *sql.Tx, userID string, userAgent string) {
	query := "update auth_session set revoked = true where user_id = ? and user_agent = ?"

	_, err := tx.ExecContext(ctx, query, userID, userAgent)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) CheckRefreshToken(ctx context.Context, tx *sql.Tx, userID string, userAgent string) (*model.AuthSession, error) {
	query := "select id, user_id, refresh_token, expired_at from auth_session where user_id = ? and user_agent = ? and revoked = false"

	rows, err := tx.QueryContext(ctx, query, userID, userAgent)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	session := new(model.AuthSession)
	if rows.Next() {
		err = rows.Scan(&session.Id, &session.UserId, &session.RefreshToken, &session.ExpiresAt)
		if err != nil {
			panic(err)
		}

		return session, nil
	}

	return session, errors.New("session not found")
}
