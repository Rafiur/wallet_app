package entity

type TransactionTypeMeta struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Color string `json:"color"`
	Icon  string `json:"icon"`
}

const (
	TransactionTypeIncome   = "income"
	TransactionTypeExpense  = "expense"
	TransactionTypeTransfer = "transfer"
	TransactionTypeRefund   = "refund"
	TransactionTypeCashback = "cashback"
)

var TransactionTypes = []TransactionTypeMeta{
	{
		Key:   TransactionTypeIncome,
		Label: "Income",
		Color: "#00C853", // Green
		Icon:  "💰",
	},
	{
		Key:   TransactionTypeExpense,
		Label: "Expense",
		Color: "#D50000", // Red
		Icon:  "💸",
	},
	{
		Key:   TransactionTypeTransfer,
		Label: "Transfer",
		Color: "#2962FF", // Blue
		Icon:  "🔁",
	},
	{
		Key:   TransactionTypeRefund,
		Label: "Refund",
		Color: "#FFD600", // Yellow
		Icon:  "↩️",
	},
	{
		Key:   TransactionTypeCashback,
		Label: "Cashback",
		Color: "#00ACC1", // Teal
		Icon:  "🎁",
	},
}

//🛣️ Step 3: Add the route
//internal/router/routes.go
//router.GET("/api/v1/transactions/types", handler.GetTransactionTypes)

//📦 Expose it via API (you’ve already done this for strings)
//handler/transaction.go
//func GetTransactionTypes(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"types": entity.TransactionTypes,
//	})
//}
