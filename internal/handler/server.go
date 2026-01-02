package handler

import (
	"github.com/Rafiur/wallet_app/internal/security"
	"github.com/Rafiur/wallet_app/internal/service"
)

type Handler struct {
	accountService              *service.AccountService
	accountCurrenciesService    *service.AccountCurrenciesService
	bankService                 *service.BankService
	budgetService               *service.BudgetService
	cashFlowSummaryService      *service.CashFlowSummaryService
	currencyService             *service.CurrencyService
	expenseCategoryService      *service.ExpenseCategoryService
	investmentService           *service.InvestmentService
	recurringTransactionService *service.RecurringTransactionService
	sessionService              *service.SessionService
	transactionService          *service.TransactionService
	transferService             *service.TransferService
	userService                 *service.UserService
	jwtService                  *security.JWTService
	passwordService             *security.PasswordService
}

func NewHandler(
	accountService *service.AccountService,
	accountCurrenciesService *service.AccountCurrenciesService,
	bankService *service.BankService,
	budgetService *service.BudgetService,
	cashFlowSummaryService *service.CashFlowSummaryService,
	currencyService *service.CurrencyService,
	expenseCategoryService *service.ExpenseCategoryService,
	investmentService *service.InvestmentService,
	recurringTransactionService *service.RecurringTransactionService,
	sessionService *service.SessionService,
	transactionService *service.TransactionService,
	transferService *service.TransferService,
	userService *service.UserService,
	jwtService *security.JWTService,
	passwordService *security.PasswordService,
) *Handler {
	return &Handler{
		accountService:              accountService,
		accountCurrenciesService:    accountCurrenciesService,
		bankService:                 bankService,
		budgetService:               budgetService,
		cashFlowSummaryService:      cashFlowSummaryService,
		currencyService:             currencyService,
		expenseCategoryService:      expenseCategoryService,
		investmentService:           investmentService,
		recurringTransactionService: recurringTransactionService,
		sessionService:              sessionService,
		transactionService:          transactionService,
		transferService:             transferService,
		userService:                 userService,
		jwtService:                  jwtService,
		passwordService:             passwordService,
	}
}
