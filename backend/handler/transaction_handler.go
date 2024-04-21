package handler

import (
	"context"
	"net/http"

	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/dto"
	"e-wallet/entity"
	"e-wallet/service"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) TransactionHandler {
	return TransactionHandler{
		transactionService: transactionService,
	}
}

func (h *TransactionHandler) GetTransactionList(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	keyword := ctx.Query(constants.TransactionListKeyword)
	sortBy := ctx.Query(constants.TransactionListSortBy)
	sort := ctx.Query(constants.TransactionListSort)
	limit := ctx.Query(constants.TransactionListLimit)
	page := ctx.Query(constants.TransactionListPage)
	from := ctx.Query(constants.TransactionListFrom)
	until := ctx.Query(constants.TransactionListUntil)

	id, isExist := ctx.Get(string(constants.UserId))
	if !isExist {
		ctx.Error(apperror.StatusUnauthorized())
		return
	}

	c := context.WithValue(ctx, constants.UserId, id)

	transactionList, err := h.transactionService.GetTransactionList(c,
		entity.PaginationParameter{
			Keyword: keyword,
			SortBy:  sortBy,
			Sort:    sort,
			Limit:   limit,
			Page:    page,
			From:    from,
			Until:   until,
		})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.MessageResponse{
		Message: constants.MsgResOK,
		Data:    dto.ToTransactionListDTO(*transactionList),
	})
}
