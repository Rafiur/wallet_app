package handler

import (
	"context"
	"net/http"
	"time"

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
	req.UserID = authUserID(c)
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
	if data.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "account not found"})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListAccounts(c echo.Context) error {
	filter := &entity.FilterAccountListRequest{
		ID:     c.QueryParam("id"),
		Name:   c.QueryParam("name"),
		UserID: authUserID(c),
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
	ctx := c.Request().Context()
	existing, err := h.accountService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "account not found"})
	}
	if existing.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "account not found"})
	}
	var req schema.Account
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.ID = id
	req.UserID = existing.UserID
	updated, err := h.accountService.Update(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteAccount(c echo.Context) error {
	ctx := c.Request().Context()
	userID := authUserID(c)
	id := c.Param("id")
	var req entity.CommonDeleteReq
	if id != "" {
		req.ID = id
	} else {
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	}
	ids := req.IDs
	if req.ID != "" {
		ids = append(ids, req.ID)
	}
	for _, accountID := range ids {
		account, err := h.accountService.GetByID(ctx, accountID)
		if err != nil || account.UserID != userID {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "account not found"})
		}
	}
	if err := h.accountService.Delete(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Users
func (h *Handler) CreateUser(c echo.Context) error {
	// schema.User.Password is tagged json:"-" (so password hashes never leak in
	// responses), which also means it can't be bound directly from the request body.
	var req struct {
		FullName string `json:"full_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	ctx := c.Request().Context()
	created, err := h.userService.Create(ctx, &schema.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	})
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
	if id != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	ctx := c.Request().Context()
	data, err := h.userService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

// ListUsers has no admin role in this app, so it only ever returns the caller's own record.
func (h *Handler) ListUsers(c echo.Context) error {
	filter := &entity.FilterUserListRequest{ID: authUserID(c)}
	ctx := c.Request().Context()
	data, err := h.userService.List(ctx, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	if id != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	// schema.User.Password is tagged json:"-" (so password hashes never leak in
	// responses), which also means it can't be bound directly from the request body.
	var req struct {
		FullName string `json:"full_name"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	ctx := c.Request().Context()
	updated, err := h.userService.Update(ctx, &schema.User{
		ID:       id,
		FullName: req.FullName,
		Password: req.Password,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	if id != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	ctx := c.Request().Context()
	if err := h.userService.Delete(ctx, &entity.CommonDeleteReq{ID: id}); err != nil {
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
	req.UserID = authUserID(c)
	ctx := c.Request().Context()
	account, err := h.accountService.GetByID(ctx, req.AccountID)
	if err != nil || account.UserID != req.UserID {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "account not found"})
	}
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
	if data.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "transaction not found"})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListTransactions(c echo.Context) error {
	filter := &entity.FilterTransactionListRequest{
		AccountID:         c.QueryParam("account_id"),
		UserID:            authUserID(c),
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
	ctx := c.Request().Context()
	existing, err := h.transactionService.GetByID(ctx, id)
	if err != nil || existing.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "transaction not found"})
	}
	var req schema.Transaction
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.ID = id
	req.UserID = existing.UserID
	updated, err := h.transactionService.Update(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteTransaction(c echo.Context) error {
	ctx := c.Request().Context()
	userID := authUserID(c)
	id := c.Param("id")
	var req entity.CommonDeleteReq
	if id != "" {
		req.ID = id
	} else {
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	}
	ids := req.IDs
	if req.ID != "" {
		ids = append(ids, req.ID)
	}
	for _, txID := range ids {
		tx, err := h.transactionService.GetByID(ctx, txID)
		if err != nil || tx.UserID != userID {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "transaction not found"})
		}
	}
	if err := h.transactionService.Delete(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Transfers
// ownAccountIDs returns the set of account ids belonging to the authenticated user.
func (h *Handler) ownAccountIDs(ctx context.Context, c echo.Context) (map[string]bool, error) {
	accounts, err := h.accountService.List(ctx, &entity.FilterAccountListRequest{UserID: authUserID(c)})
	if err != nil {
		return nil, err
	}
	ids := make(map[string]bool, len(accounts))
	for _, a := range accounts {
		ids[a.ID] = true
	}
	return ids, nil
}

func (h *Handler) CreateTransfer(c echo.Context) error {
	var req schema.Transfer
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	ctx := c.Request().Context()
	owned, err := h.ownAccountIDs(ctx, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if !owned[req.FromAccountID] || !owned[req.ToAccountID] {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "account not found"})
	}
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
	owned, err := h.ownAccountIDs(ctx, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if !owned[data.FromAccountID] && !owned[data.ToAccountID] {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "transfer not found"})
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
	owned, err := h.ownAccountIDs(ctx, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	data, err := h.transferService.List(ctx, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	filtered := make([]*schema.Transfer, 0, len(data))
	for _, t := range data {
		if owned[t.FromAccountID] || owned[t.ToAccountID] {
			filtered = append(filtered, t)
		}
	}
	return c.JSON(http.StatusOK, filtered)
}

func (h *Handler) UpdateTransfer(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	existing, err := h.transferService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "transfer not found"})
	}
	owned, err := h.ownAccountIDs(ctx, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if !owned[existing.FromAccountID] && !owned[existing.ToAccountID] {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "transfer not found"})
	}
	var req schema.Transfer
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.ID = id
	if req.FromAccountID != "" && !owned[req.FromAccountID] {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "account not found"})
	}
	if req.ToAccountID != "" && !owned[req.ToAccountID] {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "account not found"})
	}
	updated, err := h.transferService.Update(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteTransfer(c echo.Context) error {
	ctx := c.Request().Context()
	owned, err := h.ownAccountIDs(ctx, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	id := c.Param("id")
	var req entity.CommonDeleteReq
	if id != "" {
		req.ID = id
	} else {
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	}
	ids := req.IDs
	if req.ID != "" {
		ids = append(ids, req.ID)
	}
	for _, transferID := range ids {
		t, err := h.transferService.GetByID(ctx, transferID)
		if err != nil || (!owned[t.FromAccountID] && !owned[t.ToAccountID]) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "transfer not found"})
		}
	}
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
	req.UserID = authUserID(c)
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
	if data.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "session not found"})
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
	if err != nil || s.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "session not found"})
	}
	return c.JSON(http.StatusOK, s)
}

func (h *Handler) DeleteSession(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	existing, err := h.sessionService.GetByID(ctx, id)
	if err != nil || existing.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "session not found"})
	}
	if err := h.sessionService.Delete(ctx, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Budgets
func (h *Handler) CreateBudget(c echo.Context) error {
	var req schema.Budget
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.UserID = authUserID(c)
	ctx := c.Request().Context()
	if err := h.budgetService.Create(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, req)
}

func (h *Handler) GetBudget(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	data, err := h.budgetService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	if data.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "budget not found"})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListBudgets(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := h.budgetService.GetByUserID(ctx, authUserID(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateBudget(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	existing, err := h.budgetService.GetByID(ctx, id)
	if err != nil || existing.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "budget not found"})
	}
	var req schema.Budget
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.ID = id
	req.UserID = existing.UserID
	if err := h.budgetService.Update(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, req)
}

func (h *Handler) DeleteBudget(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	existing, err := h.budgetService.GetByID(ctx, id)
	if err != nil || existing.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "budget not found"})
	}
	if err := h.budgetService.Delete(ctx, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Currencies
func (h *Handler) CreateCurrency(c echo.Context) error {
	var req schema.Currency
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	ctx := c.Request().Context()
	if err := h.currencyService.Create(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, req)
}

func (h *Handler) GetCurrency(c echo.Context) error {
	code := c.Param("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "code is required"})
	}
	ctx := c.Request().Context()
	data, err := h.currencyService.GetByCode(ctx, code)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListCurrencies(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := h.currencyService.List(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateCurrency(c echo.Context) error {
	code := c.Param("code")
	var req schema.Currency
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if code != "" {
		req.Code = code
	}
	ctx := c.Request().Context()
	if err := h.currencyService.Update(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, req)
}

func (h *Handler) DeleteCurrency(c echo.Context) error {
	code := c.Param("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "code is required"})
	}
	ctx := c.Request().Context()
	if err := h.currencyService.Delete(ctx, code); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Expense Categories
// categoryOwnedByCaller reports whether a category is private to (created by) the caller.
// Global categories (UserID == nil) are visible to everyone but not owned by anyone.
func categoryOwnedByCaller(cat *schema.ExpenseCategory, userID string) bool {
	return cat.UserID != nil && *cat.UserID == userID
}

func (h *Handler) CreateExpenseCategory(c echo.Context) error {
	var req schema.ExpenseCategory
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	userID := authUserID(c)
	req.UserID = &userID
	ctx := c.Request().Context()
	created, err := h.expenseCategoryService.Create(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, created)
}

func (h *Handler) GetExpenseCategory(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	data, err := h.expenseCategoryService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	if data.UserID != nil && !categoryOwnedByCaller(data, authUserID(c)) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "expense category not found"})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListExpenseCategories(c echo.Context) error {
	filter := &entity.FilterExpenseCategoryListRequest{
		ID:               c.QueryParam("id"),
		Name:             c.QueryParam("name"),
		ParentCategoryID: c.QueryParam("parent_category_id"),
		Type:             c.QueryParam("type"),
		UserID:           authUserID(c),
	}
	ctx := c.Request().Context()
	data, err := h.expenseCategoryService.List(ctx, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateExpenseCategory(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	existing, err := h.expenseCategoryService.GetByID(ctx, id)
	if err != nil || !categoryOwnedByCaller(existing, authUserID(c)) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "expense category not found"})
	}
	var req schema.ExpenseCategory
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.ID = id
	req.UserID = existing.UserID
	updated, err := h.expenseCategoryService.Update(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteExpenseCategory(c echo.Context) error {
	ctx := c.Request().Context()
	userID := authUserID(c)
	id := c.Param("id")
	var req entity.CommonDeleteReq
	if id != "" {
		req.ID = id
	} else {
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	}
	ids := req.IDs
	if req.ID != "" {
		ids = append(ids, req.ID)
	}
	for _, catID := range ids {
		cat, err := h.expenseCategoryService.GetByID(ctx, catID)
		if err != nil || !categoryOwnedByCaller(cat, userID) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "expense category not found"})
		}
	}
	if err := h.expenseCategoryService.Delete(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Banks
func (h *Handler) CreateBank(c echo.Context) error {
	var req schema.Bank
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.UserID = authUserID(c)
	ctx := c.Request().Context()
	if err := h.bankService.Create(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, req)
}

func (h *Handler) GetBank(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	data, err := h.bankService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	if data.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "bank not found"})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListBanks(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := h.bankService.GetByUserID(ctx, authUserID(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateBank(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	existing, err := h.bankService.GetByID(ctx, id)
	if err != nil || existing.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "bank not found"})
	}
	var req schema.Bank
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.ID = id
	req.UserID = existing.UserID
	if err := h.bankService.Update(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, req)
}

func (h *Handler) DeleteBank(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	existing, err := h.bankService.GetByID(ctx, id)
	if err != nil || existing.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "bank not found"})
	}
	if err := h.bankService.Delete(ctx, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Investments
func (h *Handler) CreateInvestment(c echo.Context) error {
	var req schema.Investment
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.UserID = authUserID(c)
	ctx := c.Request().Context()
	if err := h.investmentService.Create(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, req)
}

func (h *Handler) GetInvestment(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	data, err := h.investmentService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	if data.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "investment not found"})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListInvestments(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := h.investmentService.GetByUserID(ctx, authUserID(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateInvestment(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	existing, err := h.investmentService.GetByID(ctx, id)
	if err != nil || existing.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "investment not found"})
	}
	var req schema.Investment
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.ID = id
	req.UserID = existing.UserID
	if err := h.investmentService.Update(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, req)
}

func (h *Handler) DeleteInvestment(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	existing, err := h.investmentService.GetByID(ctx, id)
	if err != nil || existing.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "investment not found"})
	}
	if err := h.investmentService.Delete(ctx, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Recurring Transactions
func (h *Handler) CreateRecurringTransaction(c echo.Context) error {
	var req schema.RecurringTransaction
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.UserID = authUserID(c)
	ctx := c.Request().Context()
	if err := h.recurringTransactionService.Create(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, req)
}

func (h *Handler) GetRecurringTransaction(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	data, err := h.recurringTransactionService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	if data.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "recurring transaction not found"})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListRecurringTransactions(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := h.recurringTransactionService.GetByUserID(ctx, authUserID(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateRecurringTransaction(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	existing, err := h.recurringTransactionService.GetByID(ctx, id)
	if err != nil || existing.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "recurring transaction not found"})
	}
	var req schema.RecurringTransaction
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.ID = id
	req.UserID = existing.UserID
	if err := h.recurringTransactionService.Update(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, req)
}

func (h *Handler) DeleteRecurringTransaction(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	existing, err := h.recurringTransactionService.GetByID(ctx, id)
	if err != nil || existing.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "recurring transaction not found"})
	}
	if err := h.recurringTransactionService.Delete(ctx, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Account Currencies
func (h *Handler) CreateAccountCurrency(c echo.Context) error {
	var req schema.AccountCurrencies
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	ctx := c.Request().Context()
	owned, err := h.ownAccountIDs(ctx, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if !owned[req.AccountID] {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "account not found"})
	}
	if err := h.accountCurrenciesService.Create(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, req)
}

func (h *Handler) GetAccountCurrency(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	data, err := h.accountCurrenciesService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	owned, err := h.ownAccountIDs(ctx, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if !owned[data.AccountID] {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "account currency not found"})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListAccountCurrencies(c echo.Context) error {
	accountID := c.QueryParam("account_id")
	if accountID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "account_id is required"})
	}
	ctx := c.Request().Context()
	owned, err := h.ownAccountIDs(ctx, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if !owned[accountID] {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "account not found"})
	}
	data, err := h.accountCurrenciesService.GetByAccountID(ctx, accountID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateAccountCurrency(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	existing, err := h.accountCurrenciesService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "account currency not found"})
	}
	owned, err := h.ownAccountIDs(ctx, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if !owned[existing.AccountID] {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "account currency not found"})
	}
	var req schema.AccountCurrencies
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.ID = id
	if req.AccountID != "" && !owned[req.AccountID] {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "account not found"})
	}
	if err := h.accountCurrenciesService.Update(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, req)
}

func (h *Handler) DeleteAccountCurrency(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	existing, err := h.accountCurrenciesService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "account currency not found"})
	}
	owned, err := h.ownAccountIDs(ctx, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if !owned[existing.AccountID] {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "account currency not found"})
	}
	if err := h.accountCurrenciesService.Delete(ctx, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Cash Flow Summary
func (h *Handler) CreateCashFlowSummary(c echo.Context) error {
	var req schema.CashFlowSummary
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.UserID = authUserID(c)
	ctx := c.Request().Context()
	if err := h.cashFlowSummaryService.Create(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, req)
}

func (h *Handler) GetCashFlowSummary(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	data, err := h.cashFlowSummaryService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	if data.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "cash flow summary not found"})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) ListCashFlowSummaries(c echo.Context) error {
	period := c.QueryParam("period")
	if period == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "period is required"})
	}
	ctx := c.Request().Context()
	data, err := h.cashFlowSummaryService.GetByUserIDAndPeriod(ctx, authUserID(c), period, time.Now()) // assuming startDate as now, adjust if needed
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateCashFlowSummary(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	existing, err := h.cashFlowSummaryService.GetByID(ctx, id)
	if err != nil || existing.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "cash flow summary not found"})
	}
	var req schema.CashFlowSummary
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.ID = id
	req.UserID = existing.UserID
	if err := h.cashFlowSummaryService.Update(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, req)
}

func (h *Handler) DeleteCashFlowSummary(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	ctx := c.Request().Context()
	existing, err := h.cashFlowSummaryService.GetByID(ctx, id)
	if err != nil || existing.UserID != authUserID(c) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "cash flow summary not found"})
	}
	if err := h.cashFlowSummaryService.Delete(ctx, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Auth handlers
func (h *Handler) Login(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	ctx := c.Request().Context()
	user, err := h.userService.GetByEmail(ctx, req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}
	if !h.passwordService.CheckPassword(req.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}
	// Generate tokens
	accessToken, err := h.jwtService.GenerateAccessToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to generate access token"})
	}
	refreshToken, err := h.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to generate refresh token"})
	}
	// Store session
	session := &schema.Session{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(h.jwtService.RefreshTTL()),
	}
	if _, err := h.sessionService.Create(ctx, session); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create session"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": map[string]string{
			"id":        user.ID,
			"full_name": user.FullName,
			"email":     user.Email,
		},
	})
}

func (h *Handler) RefreshToken(c echo.Context) error {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	ctx := c.Request().Context()
	session, err := h.sessionService.GetByRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid refresh token"})
	}
	if session.ExpiresAt.Before(time.Now()) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "refresh token expired"})
	}
	// Generate new access token
	accessToken, err := h.jwtService.GenerateAccessToken(session.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to generate access token"})
	}
	// Rotate refresh token so a used/replayed token stops working
	newRefreshToken, err := h.jwtService.GenerateRefreshToken(session.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to generate refresh token"})
	}
	session.RefreshToken = newRefreshToken
	session.ExpiresAt = time.Now().Add(h.jwtService.RefreshTTL())
	if _, err := h.sessionService.Update(ctx, session); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to rotate session"})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": newRefreshToken,
	})
}

func (h *Handler) Logout(c echo.Context) error {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if req.RefreshToken == "" {
		return c.NoContent(http.StatusNoContent)
	}
	ctx := c.Request().Context()
	session, err := h.sessionService.GetByRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return c.NoContent(http.StatusNoContent)
	}
	_ = h.sessionService.Delete(ctx, session.ID)
	return c.NoContent(http.StatusNoContent)
}
