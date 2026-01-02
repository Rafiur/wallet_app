package main

import (
	"fmt"
	"net/http"

	"github.com/Rafiur/wallet_app/internal/config"
	"github.com/Rafiur/wallet_app/internal/config/database/postgres"
	"github.com/Rafiur/wallet_app/internal/handler"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/repo_postgres"
	"github.com/Rafiur/wallet_app/internal/router"
	"github.com/Rafiur/wallet_app/internal/security"
	"github.com/Rafiur/wallet_app/internal/service"
	"github.com/Rafiur/wallet_app/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtClaim struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func main() {
	// Initialize configuration
	config.Init()
	dynamicConfig := config.GetDynamicConfig()
	if dynamicConfig == nil {
		fmt.Println("Failed to load dynamic config")
		return
	}

	// Initialize logger
	logger.NewApiLogger(dynamicConfig).InitLogger()

	// Initialize database
	db, err := postgres.NewPostgresDB()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}
	defer db.Close()

	// Initialize repositories
	accountRepo := repo_postgres.NewAccountRepo(db)
	accountService := service.NewAccountService(accountRepo)
	accountCurrenciesRepo := repo_postgres.NewAccountCurrenciesRepository(db)
	accountCurrenciesService := service.NewAccountCurrenciesService(accountCurrenciesRepo)
	bankRepo := repo_postgres.NewBankRepository(db)
	bankService := service.NewBankService(bankRepo)
	budgetRepo := repo_postgres.NewBudgetRepository(db)
	budgetService := service.NewBudgetService(budgetRepo)
	cashFlowSummaryRepo := repo_postgres.NewCashFlowSummaryRepository(db)
	cashFlowSummaryService := service.NewCashFlowSummaryService(cashFlowSummaryRepo)
	currencyRepo := repo_postgres.NewCurrencyRepository(db)
	currencyService := service.NewCurrencyService(currencyRepo)
	expenseCategoryRepo := repo_postgres.NewExpenseCategory(db)
	expenseCategoryService := service.NewExpenseCategoryService(expenseCategoryRepo)
	investmentRepo := repo_postgres.NewInvestmentRepository(db)
	investmentService := service.NewInvestmentService(investmentRepo)
	recurringTransactionRepo := repo_postgres.NewRecurringTransactionRepository(db)
	recurringTransactionService := service.NewRecurringTransactionService(recurringTransactionRepo)
	sessionRepo := repo_postgres.NewSessionRepo(db)
	sessionService := service.NewSessionService(sessionRepo)
	transactionRepo := repo_postgres.NewTransactionRepo(db)
	transactionService := service.NewTransactionService(transactionRepo, accountRepo)
	transferRepo := repo_postgres.NewTransferRepo(db)
	transferService := service.NewTransferService(transferRepo, accountRepo)
	userRepo := repo_postgres.NewUserRepo(db)
	passwordService := security.NewPasswordService()
	userService := service.NewUserService(userRepo, passwordService)

	// Initialize security services
	jwtService := security.NewJWTService(dynamicConfig.JwtSecret)

	// Initialize handler
	mainHandler := handler.NewHandler(
		accountService,
		accountCurrenciesService,
		bankService,
		budgetService,
		cashFlowSummaryService,
		currencyService,
		expenseCategoryService,
		investmentService,
		recurringTransactionService,
		sessionService,
		transactionService,
		transferService,
		userService,
		jwtService,
		passwordService,
	)
	// Initialize Echo
	e := echo.New()
	e.Use(
		middleware.Gzip(),
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:3000", "https://your-frontend.com"},
			AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
			AllowCredentials: true,
			MaxAge:           86400,
		}),
		middleware.Logger(),
	)

	// JWT middleware
	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtClaim)
		},
		SigningKey: []byte(dynamicConfig.JwtSecret),
		//ErrorHandler: middlewares.JwtErrorHandler,
	}
	authenticate := echojwt.WithConfig(jwtConfig)

	// Route initialization
	router.Route(e, authenticate, mainHandler)

	// Start server
	port := dynamicConfig.ServerPort
	if port == "" {
		port = "8080"
	}
	if err := e.Start(":" + port); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
