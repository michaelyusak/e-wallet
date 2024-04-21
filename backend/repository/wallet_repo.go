package repository

import (
	"context"
	"database/sql"
	"e-wallet/entity"
	"fmt"

	"github.com/shopspring/decimal"
)

type WalletRepository interface {
	GetCashFlow(ctx context.Context, wallet *entity.Wallet) error
	GetWalletByNum(ctx context.Context, walletNum string) (*entity.Wallet, error)
	GetWalletByUserId(ctx context.Context, userId int) (*entity.Wallet, error)
	PostOneTransaction(ctx context.Context, transferReq entity.Transaction) error
	PostOneWallet(ctx context.Context, userId int) (*entity.Wallet, error)
	UpdateBalance(ctx context.Context, walletId int, amount decimal.Decimal) error
	UpdateGachaTrial(ctx context.Context, userId int, getTrial decimal.Decimal) error
	WithTx(ctx context.Context, fn func(repo WalletRepository) error) error
}

type walletRepositoryPostgres struct {
	db   *sql.DB
	dbtx DBTX
}

func NewWalletRepositoryPostgres(db *sql.DB) walletRepositoryPostgres {
	return walletRepositoryPostgres{
		db:   db,
		dbtx: db,
	}
}

func (r *walletRepositoryPostgres) injectDBTX(dbtx DBTX) {
	r.dbtx = dbtx
}

func (r *walletRepositoryPostgres) PostOneWallet(ctx context.Context, userId int) (*entity.Wallet, error) {
	query := `
		INSERT INTO wallets (user_id, gacha_trial, created_at)
		VALUES ($1, 0, NOW())
		RETURNING wallet_number, balance
	`

	var wallet entity.Wallet

	err := r.dbtx.QueryRowContext(ctx, query, userId).Scan(&wallet.Number, &wallet.Balance)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (r *walletRepositoryPostgres) GetWalletByNum(ctx context.Context, walletNum string) (*entity.Wallet, error) {
	query := `
		SELECT wallet_id, user_id, balance, gacha_trial
		FROM wallets
		WHERE wallet_number = $1
		AND deleted_at
			IS NULL
	`

	var wallet entity.Wallet

	err := r.dbtx.QueryRowContext(ctx, query, walletNum).Scan(&wallet.Id, &wallet.UserId, &wallet.Balance, &wallet.GachaTrial)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	wallet.Number = walletNum

	return &wallet, nil
}

func (r *walletRepositoryPostgres) UpdateBalance(ctx context.Context, walletId int, amount decimal.Decimal) error {
	query := `
		UPDATE wallets
		SET balance = balance + $1
		WHERE wallet_id = $2
		AND deleted_at
			IS NULL
	`

	var wallet entity.Wallet

	_, err := r.dbtx.ExecContext(ctx, query, amount, walletId)
	if err != nil {
		return err
	}

	wallet.Id = walletId

	return nil
}

func (r *walletRepositoryPostgres) GetCashFlow(ctx context.Context, wallet *entity.Wallet) error {
	query := `
	SELECT COALESCE(income.total_income, 0) AS income, COALESCE(expense.total_expense, 0) AS expense
		FROM wallets w
		LEFT JOIN (
    		SELECT recepient_wallet_number, SUM(amount) AS total_income
    			FROM transactions
    			GROUP 
					BY recepient_wallet_number) 
			AS income 
				ON w.wallet_number = income.recepient_wallet_number
		LEFT JOIN (
    		SELECT sender_wallet_number, SUM(amount) AS total_expense
    			FROM transactions
    			GROUP 
					BY sender_wallet_number)
			AS expense 
				ON w.wallet_number = expense.sender_wallet_number
		WHERE w.wallet_id = $1
	`

	err := r.dbtx.QueryRowContext(ctx, query, wallet.Id).Scan(&wallet.Income, &wallet.Expense)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	return nil
}

func (r *walletRepositoryPostgres) GetWalletByUserId(ctx context.Context, userId int) (*entity.Wallet, error) {
	query := `
	SELECT w.wallet_id, w.wallet_number, w.balance, w.gacha_trial
		FROM wallets w
		JOIN users u 
			ON w.user_id = u.user_id
				WHERE u.user_id = $1
	`

	var wallet entity.Wallet

	err := r.dbtx.QueryRowContext(ctx, query, userId).Scan(&wallet.Id, &wallet.Number, &wallet.Balance, &wallet.GachaTrial)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	wallet.UserId = userId

	return &wallet, nil
}

func (r *walletRepositoryPostgres) UpdateGachaTrial(ctx context.Context, userId int, getTrial decimal.Decimal) error {
	query := `
		UPDATE wallets
		SET gacha_trial = gacha_trial + $1
		WHERE user_id = $2
			AND deleted_at
				IS NULL
	`

	_, err := r.dbtx.ExecContext(ctx, query, getTrial, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *walletRepositoryPostgres) PostOneTransaction(ctx context.Context, transferReq entity.Transaction) error {
	query := `
		INSERT INTO transactions (sender_wallet_number, recepient_wallet_number, amount, source_of_fund, description, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
	`

	_, err := r.dbtx.ExecContext(ctx, query, transferReq.SenderWalletNum, transferReq.RecepientWalletNum, transferReq.Amount, transferReq.SourceOfFunds, transferReq.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *walletRepositoryPostgres) WithTx(ctx context.Context, fn func(repo WalletRepository) error) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	repo := NewWalletRepositoryPostgres(r.db)
	repo.injectDBTX(tx)

	err = fn(&repo)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf(err.Error(), rbErr.Error())
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
