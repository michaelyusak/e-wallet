package constants

type contextKey string

var (
	UserId    = contextKey("user-id")
	RequestId = contextKey("request-id")

	TransactionListKeyword = "search"
	TransactionListSortBy  = "sortBy"
	TransactionListSort    = "sort"
	TransactionListLimit   = "limit"
	TransactionListPage    = "page"
	TransactionListFrom    = "from"
	TransactionListUntil   = "until"

	GameGachaSelection = "selection"
)
