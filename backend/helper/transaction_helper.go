package helper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/entity"
)

var (
	columns = []string{"date", "amount", "to"}
)

func isSortByValid(sortBy string) bool {
	for _, column := range columns {
		if strings.ToLower(sortBy) == column {
			return true
		}
	}

	return false
}

func swapSortBy(sortBy *string) {
	switch strings.ToLower(*sortBy) {
	case columns[0]:
		*sortBy = constants.SortByDate
	case columns[1]:
		*sortBy = constants.SortByAmout
	case columns[2]:
		*sortBy = constants.SortByTo
	default:
		*sortBy = constants.SortByDate
	}
}

func checkDate(from, until *string) error {
	fromTime, err := time.Parse(constants.TimeStr, *from)
	if err != nil {
		if *from != "" {
			return apperror.BadRequest("date format should be YYYY-MM-DD")
		}
		*from = constants.DefaultFrom
	}

	untilTime, err := time.Parse(constants.TimeStr, *until)
	if err != nil {
		if *until != "" {
			return apperror.BadRequest("date format should be YYYY-MM-DD")
		}
		*until = time.Now().Format(constants.TimeStr)
	}

	if untilTime.Before(fromTime) {
		return apperror.BadRequest("date in field from must be before date in field until")
	}

	return nil
}

func CheckPaginationParams(params *entity.PaginationParameter) error {
	if strings.Contains(params.Keyword, "not") {
		params.Keyword = strings.Join(strings.Split(params.Keyword, " ")[1:], " ")
		params.Not = "NOT"
	}

	if !isSortByValid(params.SortBy) {
		if params.SortBy != "" {
			return apperror.BadRequest(fmt.Sprintf("parameter sort by is invalid, use available keywords "+
				"or leave it empty to apply default value, available keywords: %s",
				strings.Join(columns, ", ")))
		}
	}

	swapSortBy(&params.SortBy)

	if params.Sort != "asc" && params.Sort != "desc" {
		if params.Sort != "" {
			return apperror.BadRequest("parameter sort must be either asc (ascending) or desc (descending)")
		}

		params.Sort = constants.DefaultSort
	}

	limitInt, err := strconv.Atoi(params.Limit)
	if err != nil {
		if params.Limit != "" {
			return apperror.BadRequest("parameter limit must be a number")
		}

		params.Limit = constants.DefaultLimit
	}

	if params.Page == "" {
		params.Page = constants.DefaultPage
	}

	pageInt, err := strconv.Atoi(params.Page)
	if err != nil {
		return apperror.BadRequest("parameter page must be a number")
	}

	params.Offset = fmt.Sprintf("%v", limitInt*(pageInt-1))

	err = checkDate(&params.From, &params.Until)
	if err != nil {
		return err
	}

	return nil
}

func FormatTransactionList(transactionList *entity.TransactionList, walletNumber string, page string) {
	transactions := &transactionList.Transactions
	listPage := &transactionList.Page

	var formattedTransactions []entity.Transaction

	for _, transaction := range *transactions {
		if transaction.SenderWalletNum == walletNumber {
			transaction.Type = "expense"

		} else if transaction.RecepientWalletNum == walletNumber {
			transaction.Type = "income"
		}

		formattedTransactions = append(formattedTransactions, transaction)
	}

	*transactions = formattedTransactions
	*listPage = page
}
