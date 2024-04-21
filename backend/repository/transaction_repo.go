package repository

import (
	"context"
	"database/sql"
	"e-wallet/entity"
	"fmt"
)

type TransactionRepository interface {
	GetTransactions(ctx context.Context, walletNumber string, param entity.PaginationParameter) (*entity.TransactionList, error)
}

type transactionRepositoryPostgres struct {
	db *sql.DB
}

func NewTransactionRepositoryPostgres(db *sql.DB) transactionRepositoryPostgres {
	return transactionRepositoryPostgres{
		db: db,
	}
}

func (r *transactionRepositoryPostgres) GetTransactions(ctx context.Context, walletNumber string, param entity.PaginationParameter) (*entity.TransactionList, error) {
	query := fmt.Sprintf(
		"SELECT COUNT(*) OVER(), transaction_id, sender_wallet_number, u1.user_name, recepient_wallet_number, u2.user_name, amount, source_of_fund, description, t.created_at "+
			"FROM transactions t "+
			"JOIN wallets w1 ON t.sender_wallet_number = w1.wallet_number "+
			"JOIN users u1 ON w1.user_id = u1.user_id "+
			"JOIN wallets w2 ON t.recepient_wallet_number = w2.wallet_number "+
			"JOIN users u2 ON w2.user_id = u2.user_id "+
			"WHERE description %s ILIKE '%%%s%%' "+
			"AND (sender_wallet_number = '%s' "+
			"OR recepient_wallet_number = '%s') "+
			"AND t.created_at "+
			"BETWEEN '%s' "+
			"AND '%s 23:59' "+
			"ORDER BY %s "+
			"%s "+
			"LIMIT %s "+
			"OFFSET %s", param.Not, param.Keyword, walletNumber, walletNumber, param.From, param.Until, param.SortBy, param.Sort, param.Limit, param.Offset)

	var transactionList entity.TransactionList
	var transactions []entity.Transaction

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction entity.Transaction

		err := rows.Scan(&transactionList.TotalItem, &transaction.Id, &transaction.SenderWalletNum, &transaction.SenderName, &transaction.RecepientWalletNum, &transaction.RecipientName,
			&transaction.Amount, &transaction.SourceOfFunds, &transaction.Description, &transaction.Date)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	transactionList.Transactions = transactions

	return &transactionList, nil
}
