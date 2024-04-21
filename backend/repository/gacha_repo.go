package repository

import (
	"context"
	"database/sql"
)

type GachaRepository interface {
	PostOneUserGacha(ctx context.Context, prizeId int, walletId int) error
}

type gachaRepositoryPostgres struct {
	db *sql.DB
}

func NewGachaRepositoryPostgres(db *sql.DB) gachaRepositoryPostgres {
	return gachaRepositoryPostgres{
		db: db,
	}
}

func (r *gachaRepositoryPostgres) PostOneUserGacha(ctx context.Context, prizeId int, walletId int) error {
	query := `
		INSERT INTO user_gachas (wallet_id, prize_id, created_at)
		VALUES ($1, $2, NOW())
	`

	_, err := r.db.ExecContext(ctx, query, walletId, prizeId)
	if err != nil {
		return err
	}

	return nil
}
