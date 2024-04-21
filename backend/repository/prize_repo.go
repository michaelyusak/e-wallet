package repository

import (
	"context"
	"database/sql"
	"e-wallet/entity"
)

type PrizeRepository interface {
	GetAllBox(ctx context.Context) ([]entity.GachaBox, error)
}

type prizeRepositoryPostgres struct {
	db *sql.DB
}

func NewPrizeRepositoryPostgres(db *sql.DB) prizeRepositoryPostgres {
	return prizeRepositoryPostgres{
		db: db,
	}
}

func (r *prizeRepositoryPostgres) GetAllBox(ctx context.Context) ([]entity.GachaBox, error) {
	query := `
		SELECT prize_id, amount
		FROM gacha_prizes
		WHERE deleted_at
			IS NULL
		ORDER BY
			RANDOM()
	`

	var gachaBoxes []entity.GachaBox

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var gachaBox entity.GachaBox

		err := rows.Scan(&gachaBox.PrizeId, &gachaBox.Prize)
		if err != nil {
			return nil, err
		}

		gachaBoxes = append(gachaBoxes, gachaBox)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return gachaBoxes, nil
}
