package security

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// newJTI returns a random token identifier so two tokens issued for the same
// user within the same second (e.g. rapid double-login) never collide -
// JWTs are otherwise deterministic given identical claims and secret.
func newJTI() string {
	buf := make([]byte, 16)
	_, _ = rand.Read(buf)
	return hex.EncodeToString(buf)
}

type JwtClaim struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTService struct {
	secretKey  string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewJWTService(secretKey string, accessTTL, refreshTTL time.Duration) *JWTService {
	if accessTTL <= 0 {
		accessTTL = 15 * time.Minute
	}
	if refreshTTL <= 0 {
		refreshTTL = 7 * 24 * time.Hour
	}
	return &JWTService{secretKey: secretKey, accessTTL: accessTTL, refreshTTL: refreshTTL}
}

func (s *JWTService) GenerateAccessToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(s.accessTTL).Unix(),
		"iat":     time.Now().Unix(),
		"jti":     newJTI(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *JWTService) RefreshTTL() time.Duration {
	return s.refreshTTL
}

func (s *JWTService) GenerateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(s.refreshTTL).Unix(),
		"iat":     time.Now().Unix(),
		"jti":     newJTI(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *JWTService) ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
