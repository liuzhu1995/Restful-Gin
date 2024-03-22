package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/* secretKey 是服务端私有密钥，通常用于签名和验证JWT
* 签名（Signing）：当创建一个新的JWT时，你会使用这个密钥和一个签名算法（如HS256）来生成一个签名。这个签名附加在JWT的末尾，作为令牌的一部分
* 验证（Verification）：当接收到一个JWT时，你可以使用同样的secretKey和签名算法来验证签名。如果签名有效，那么你知道令牌在传输过程中没有被篡改，并且它是由拥有该secretKey的实体生成的
* 安全性（Security）：secretKey 的保密性对于JWT的安全性至关重要。如果攻击者知道了secretKey，他们就可以伪造JWT，这可能导致未授权访问你的应用程序
 */
const secretKey = "your_secret_key"
// 根据给定的email和userId生成JWT令牌
func GenerateToken(email string, userId int64) (string, error) {
	// 创建新的token
	/*
		jwt.NewWithClaims 用来创建一个新的 JWT 实例
		第一个参数是 JWT 的签名算法（比如 jwt.SigningMethodHS256，jwt.SigningMethodRS256 等）。这个算法决定了如何对 JWT 的负载（payload）进行签名，以确保其完整性和发送者的身份
		第二个参数是 JWT 的负载（payload），通常是一个 jwt.MapClaims 或 jwt.StandardClaims 实例，包含了关于 JWT 的元数据和自定义声明
	*/
	 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(), // token生命周期2个小时后失效
	})
	//SignedString(): 对JWT进行签名，并返回签名后的 JWT 字符串。需要提供一个密钥secretKey（通常是一个字节数组 []byte）
	return token.SignedString([]byte(secretKey))
}
var parsedToken *jwt.Token
func VerifyToken(tokenString  string) (int64, error) {
	var err error
	// 解析JWT
	parsedToken, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		/* token.Method.(*jwt.SigningMethodHMAC) 这个类型断言的作用是检查 token.Method 是否可以被断言为 *jwt.SigningMethodHMAC 类型
				在JWT（JSON Web Tokens）的上下文中，token.Method 通常是一个接口，表示JWT的签名算法。JWT库提供了多种签名算法的实现，
			包括HMAC（基于哈希的消息认证码）和RSA等。*jwt.SigningMethodHMAC 是 jwt.SigningMethod 接口的一个具体实现，用于HMAC签名算法

				这种类型断言通常用于在处理JWT时进行条件判断，以确保使用正确的签名算法和密钥进行验证。例如，可以使用这种类型断言来检查JWT是否
			使用HMAC签名算法，并据此选择相应的密钥进行验证
		*/
		 _, ok := token.Method.(*jwt.SigningMethodHMAC)
		 if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, fmt.Errorf("unable to parse token%v", err)
	}

	tokenValid := parsedToken.Valid

	if !tokenValid  {
		// 如果 token 无效，返回错误
		return 0, errors.New("invalid token")
	}
	// 将Claims字段断言为jwt.MapClaims类型
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
 
	// 从负载claims中提取email和userId
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	// .(string) 检查是否是字符串
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
