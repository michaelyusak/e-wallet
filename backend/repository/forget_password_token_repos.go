package repository

import (
	"context"
	"database/sql"
	"e-wallet/entity"
)

type RPTRepository interface {
	DeleteExisting(ctx context.Context, userId int) error
	GetTokenExpiredAt(ctx context.Context, token string) (*entity.ResetPasswordToken, error)
	PostOneRPT(ctx context.Context, rptReq entity.ResetPasswordToken, userId int) error
}

type rptRepositoryPostgres struct {
	db *sql.DB
}

func NewRPTRepositoryPostgres(db *sql.DB) rptRepositoryPostgres {
	return rptRepositoryPostgres{
		db: db,
	}
}

func (r *rptRepositoryPostgres) PostOneRPT(ctx context.Context, rptReq entity.ResetPasswordToken, userId int) error {
	query := `
		INSERT INTO reset_password_tokens (user_id, reset_password_token, expired_at, created_at)
		VALUES ($1, $2, $3, NOW())
	`

	_, err := r.db.ExecContext(ctx, query, userId, rptReq.Token, rptReq.ExpiredAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *rptRepositoryPostgres) DeleteExisting(ctx context.Context, userId int) error {
	query := `
		UPDATE reset_password_tokens
		SET deleted_at = NULL
		WHERE user_id = $1
	`

	_, err := r.db.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *rptRepositoryPostgres) GetTokenExpiredAt(ctx context.Context, token string) (*entity.ResetPasswordToken, error) {
	query := `
		SELECT expired_at
		FROM reset_password_tokens
		WHERE reset_password_token = $1
			AND deleted_at IS NULL
	`

	var rpt entity.ResetPasswordToken

	err := r.db.QueryRowContext(ctx, query, token).Scan(&rpt.ExpiredAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &rpt, nil
}
