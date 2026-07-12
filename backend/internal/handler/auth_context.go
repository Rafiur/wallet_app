package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/Rafiur/wallet_app/internal/security"
)

// authUserID returns the id of the user identified by the JWT on the request.
// The auth middleware has already validated the token before the handler runs.
func authUserID(c echo.Context) string {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return ""
	}
	claims, ok := token.Claims.(*security.JwtClaim)
	if !ok {
		return ""
	}
	return claims.UserID
}
