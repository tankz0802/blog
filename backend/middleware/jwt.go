package middleware

import (
	"blog/config"
	error_msg "blog/error"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func JwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": error_msg.ERROR_TOKEN_NOT_EXIST,
				"msg":    error_msg.GetErrorMsg(error_msg.ERROR_TOKEN_NOT_EXIST),
			})
			c.Abort()
			return
		}
		claims, code := ParseToken(token)
		if code != error_msg.SUCCESS {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": code,
				"msg":    error_msg.GetErrorMsg(code),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	Username  string `json:"username"`
	jwt.StandardClaims
}

// CreateToken 生成一个token
func CreateToken(username string) (string, int) {
	claims := CustomClaims{
		username,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3660), // 过期时间 3600s
			Issuer:    config.TokenIssuer,                   //签名的发行者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(config.TokenSignKey))
	if err != nil {
		log.Println(err.Error())
		return "", error_msg.ERROR
	}
	return signedString, error_msg.SUCCESS
}

// 解析Token
func ParseToken(tokenString string) (*CustomClaims, int) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenSignKey), nil
	})
	if err != nil {
		log.Println(err.Error())
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, error_msg.ERROR_TOKEN_FORMAT_ERROR
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, error_msg.ERROR_TOKEN_EXPIRED
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, error_msg.ERROR_TOKEN_NOT_AUTH
			} else {
				return nil, error_msg.ERROR_TOKEN_INVAILD
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, error_msg.SUCCESS
	}
	return nil, error_msg.ERROR
}

// 更新token
func RefreshToken(tokenString string) (string, int) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.TokenSignKey, nil
	})
	if err != nil {
		return "", error_msg.ERROR
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(12) * time.Hour).Unix()
		return CreateToken(claims.Username)
	}
	return "", error_msg.ERROR
}
