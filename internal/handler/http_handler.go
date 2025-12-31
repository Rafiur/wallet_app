package handler

import (
	"net/http"

	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
	"github.com/labstack/echo/v4"
)

// Account handlers
func (h *Handler) CreateAccount(c echo.Context) error {
	var req schema.Account
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	ctx := c.Request().Context()
	created, err := h.accountService.Create(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, created)
}

func (h *Handler) GetAccount(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	data, err := h.accountService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListAccounts(c echo.Context) error {
	filter := &entity.FilterAccountListRequest{
		ID:     c.QueryParam("id"),
		Name:   c.QueryParam("name"),
		UserID: c.QueryParam("user_id"),
	}
	ctx := c.Request().Context()
	data, err := h.accountService.List(ctx, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateAccount(c echo.Context) error {
	id := c.Param("id")
	var req schema.Account
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if id != "" {
		req.ID = id
	}
	ctx := c.Request().Context()
	updated, err := h.accountService.Update(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteAccount(c echo.Context) error {
	id := c.Param("id")
	var req entity.CommonDeleteReq
	if id != "" {
		req.ID = id
	} else {
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	}
	ctx := c.Request().Context()
	if err := h.accountService.Delete(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Users
func (h *Handler) CreateUser(c echo.Context) error {
	var req schema.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	ctx := c.Request().Context()
	created, err := h.userService.Create(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, created)
}

func (h *Handler) GetUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	data, err := h.userService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListUsers(c echo.Context) error {
	filter := &entity.FilterUserListRequest{
		ID:       c.QueryParam("id"),
		FullName: c.QueryParam("full_name"),
	}
	ctx := c.Request().Context()
	data, err := h.userService.List(ctx, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var req schema.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if id != "" {
		req.ID = id
	}
	ctx := c.Request().Context()
	updated, err := h.userService.Update(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	var req entity.CommonDeleteReq
	if id != "" {
		req.ID = id
	} else {
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	}
	ctx := c.Request().Context()
	if err := h.userService.Delete(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Transactions
func (h *Handler) CreateTransaction(c echo.Context) error {
	var req schema.Transaction
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	ctx := c.Request().Context()
	created, err := h.transactionService.Create(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, created)
}

func (h *Handler) GetTransaction(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	data, err := h.transactionService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListTransactions(c echo.Context) error {
	filter := &entity.FilterTransactionListRequest{
		AccountID:         c.QueryParam("account_id"),
		UserID:            c.QueryParam("user_id"),
		ExpenseCategoryID: c.QueryParam("expense_category_id"),
		TransactionType:   c.QueryParam("transaction_type"),
	}
	ctx := c.Request().Context()
	data, err := h.transactionService.List(ctx, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateTransaction(c echo.Context) error {
	id := c.Param("id")
	var req schema.Transaction
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if id != "" {
		req.ID = id
	}
	ctx := c.Request().Context()
	updated, err := h.transactionService.Update(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteTransaction(c echo.Context) error {
	id := c.Param("id")
	var req entity.CommonDeleteReq
	if id != "" {
		req.ID = id
	} else {
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	}
	ctx := c.Request().Context()
	if err := h.transactionService.Delete(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Transfers
func (h *Handler) CreateTransfer(c echo.Context) error {
	var req schema.Transfer
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	ctx := c.Request().Context()
	created, err := h.transferService.Create(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, created)
}

func (h *Handler) GetTransfer(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	data, err := h.transferService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListTransfers(c echo.Context) error {
	filter := &entity.FilterTransferListRequest{
		FromAccountID: c.QueryParam("from_account_id"),
		ToAccountID:   c.QueryParam("to_account_id"),
		Currency:      c.QueryParam("currency"),
		Status:        c.QueryParam("status"),
	}
	ctx := c.Request().Context()
	data, err := h.transferService.List(ctx, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateTransfer(c echo.Context) error {
	id := c.Param("id")
	var req schema.Transfer
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if id != "" {
		req.ID = id
	}
	ctx := c.Request().Context()
	updated, err := h.transferService.Update(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteTransfer(c echo.Context) error {
	id := c.Param("id")
	var req entity.CommonDeleteReq
	if id != "" {
		req.ID = id
	} else {
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	}
	ctx := c.Request().Context()
	if err := h.transferService.Delete(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Sessions
func (h *Handler) CreateSession(c echo.Context) error {
	var req schema.Session
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	ctx := c.Request().Context()
	created, err := h.sessionService.Create(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, created)
}

func (h *Handler) GetSession(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	data, err := h.sessionService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListSessions(c echo.Context) error {
	refreshToken := c.QueryParam("refresh_token")
	if refreshToken == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "refresh_token or user-based listing not supported via this endpoint"})
	}
	ctx := c.Request().Context()
	s, err := h.sessionService.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, s)
}

func (h *Handler) DeleteSession(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	if err := h.sessionService.Delete(ctx, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
