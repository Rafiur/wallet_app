package router

import (
	"github.com/Rafiur/wallet_app/internal/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

//func SetupRouter(h *handler.Handler) *echo.Echo {
//	e := echo.New()
//
//	// Health check
//	e.GET("/health", func(c echo.Context) error {
//		return c.JSON(http.StatusOK, map[string]string{
//			"status":  "ok",
//			"message": "Wallet app running - clean architecture preserved",
//		})
//	})
//
//	// Example route using your handler (add more later)
//	// e.POST("/api/v1/users/signup", h.UserSignupHandler) // when you implement it
//
//	return e
//}

func Route(e *echo.Echo, authenticate echo.MiddlewareFunc, mainHandler *handler.Handler) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "ok",
			"message": "Wallet App API is running",
			"version": "0.1.0", // you can change this or make it dynamic later
		})
	})
}
