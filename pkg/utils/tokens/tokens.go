package tokens

import (
	"time"

	"github.com/Ocyss/Qblog/pkg/errno"

	"github.com/Ocyss/Qblog/pkg/configs"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GetToken 生成token
func GetToken(id uint, username string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24 * 30) // 一个月过期
	SetClaims := MyClaims{
		id,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "Ocyss_04",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	return reqClaim.SignedString([]byte(configs.GetApp().Jwt))
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, error) {
	key, err := jwt.ParseWithClaims(token, &MyClaims{}, func(*jwt.Token) (any, error) {
		return []byte(configs.GetApp().Jwt), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := key.Claims.(*MyClaims); ok && key.Valid {
		return claims, nil
	} else {
		return nil, errno.TokenExpired
	}
}
