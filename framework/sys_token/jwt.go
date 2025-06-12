package sys_token

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	CtxUserKey  = "ctx.user.key"
	CtxTokenKey = "ctx.token.key"
)

// User 结构体定义
type User struct {
	Uid  int64  `json:"uid"`
	Name string `json:"name"`
}

// GenerateToken 用于生成 JWT token
func GenerateToken(user *User, secretKey string) (string, error) {
	// 将用户对象转换为JSON字符串
	data, _ := json.Marshal(user)
	createTime := time.Now().Unix()
	// 创建JWT token并设置声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":   string(data),
		"create": createTime, // token创建时间
	})

	// 使用密钥对token进行签名
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

// ParseToken 用于解析 JWT token
func ParseToken(tokenString string, secretKey string) (*User, error) {
	// 解析并验证token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法是否为HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回密钥
		return []byte(secretKey), nil
	})

	// 检查解析错误
	if err != nil {
		return nil, err
	}

	// 解析成功，提取用户信息
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userData, _ := claims["user"].(string)
		var user User
		// 将 JSON 数据解码为用户结构体
		if err := json.Unmarshal([]byte(userData), &user); err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// GetUserFromCtx get user from ctx
func GetUserFromCtx(ctx context.Context) *User {
	user := ctx.Value(CtxUserKey)
	if user == nil {
		return new(User)
	}
	u, ok := user.(*User)
	if !ok {
		return new(User)
	}
	return u
}

// GetTokenFromCtx get token from ctx
func GetTokenFromCtx(ctx context.Context) string {
	tokenStr := ctx.Value(CtxTokenKey)
	if tokenStr == nil {
		return ""
	}
	str, ok := tokenStr.(string)
	if !ok {
		return ""
	}
	return str
}
