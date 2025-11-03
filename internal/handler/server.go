package handler

import (
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
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
}

func NewHandler(
	accountRepo repository.AccountRepoInterface,
	accountCurrenciesRepo repository.AccountCurrenciesRepository,
	bankRepo repository.BankRepository,
	budgetRepo repository.BudgetRepository,
	cashFlowSummaryRepo repository.CashFlowSummaryRepository,
	currencyRepo repository.CurrencyRepository,
	expenseCategoryRepo repository.ExpenseCategoryRepoInterface,
	investmentRepo repository.InvestmentRepository,
	recurringTransactionRepo repository.RecurringTransactionRepository,
	sessionRepo repository.SessionRepoInterface,
	transactionRepo repository.TransactionRepoInterface,
	transferRepo repository.TransferRepoInterface,
	userRepo repository.UserRepoInterface,
) *Handler {
	return &Handler{
		accountService:              service.NewAccountService(accountRepo),
		accountCurrenciesService:    service.NewAccountCurrenciesService(accountCurrenciesRepo),
		bankService:                 service.NewBankService(bankRepo),
		budgetService:               service.NewBudgetService(budgetRepo),
		cashFlowSummaryService:      service.NewCashFlowSummaryService(cashFlowSummaryRepo),
		currencyService:             service.NewCurrencyService(currencyRepo),
		expenseCategoryService:      service.NewExpenseCategoryService(expenseCategoryRepo),
		investmentService:           service.NewInvestmentService(investmentRepo),
		recurringTransactionService: service.NewRecurringTransactionService(recurringTransactionRepo),
		sessionService:              service.NewSessionService(sessionRepo),
		transactionService:          service.NewTransactionService(transactionRepo, accountRepo),
		transferService:             service.NewTransferService(transferRepo, accountRepo),
		userService:                 service.NewUserService(userRepo),
	}
}
