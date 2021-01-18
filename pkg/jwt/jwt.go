package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorbac/app/models"
	"gorbac/pkg/config"
	"strconv"
	"time"
)

// 定义授权保存信息
type CustomClaims struct {
	Id       uint64
	Username string
	jwt.StandardClaims
}

// 签名字符串
var secretary = config.GetString("jwt.secretary")

// 创建 GetToken
func GetToken(account models.Account) (map[string]interface{}, error) {
	// 7200秒过期
	maxAge, _ := strconv.Atoi(config.GetString("jwt.export"))
	// 获取两小时后的时间戳
	expTime := time.Now().Add(time.Duration(maxAge) * time.Second).Unix()
	// 设置授权保存信息
	customClaims := &CustomClaims{
		Id:       account.Id,
		Username: account.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,  // 过期时间，必须设置
			Issuer:    "iszmxw", // 非必须，也可以填充用户名，
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(secretary))
	if err != nil {
		return nil, err
	}
	rlt := make(map[string]interface{})
	rlt["expTime"] = expTime
	rlt["token"] = tokenString
	return rlt, nil
}

// 验证 Token
func AuthToken(tokenString string) (uint64, interface{}) {
	if tokenString == "" {
		return 0, "认证失败"
	}
	//kv := strings.Split(tokenString, " ")
	//if kv[0] != "Bearer" {
	//	return 0, "认证失败"
	//}
	//tokenString = kv[1]
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretary), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := uint64(claims["Id"].(float64))
		return id, nil
	} else {
		return 0, "认证已过期"
	}
}
