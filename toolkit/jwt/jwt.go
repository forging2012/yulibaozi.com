package jwt


import (
	"errors"
	"time"


	jwt "github.com/dgrijalva/jwt-go"
)

// type StandardClaims struct {
// 	// 接收jwt的一方
// 	Audience  string `json:"aud,omitempty"`
// 	// jwt的过期时间，这个过期时间必须要大于签发时间
// 	ExpiresAt int64  `json:"exp,omitempty"`
// 	//  jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
// 	Id        string `json:"jti,omitempty"`
// 	// jwt的签发时间
// 	IssuedAt  int64  `json:"iat,omitempty"`
// 	//jwt签发者
// 	Issuer    string `json:"iss,omitempty"`
// 	// 定义在什么时间之前，该jwt都是不可用的.
// 	NotBefore int64  `json:"nbf,omitempty"`
// 	// jwt所面向的用户
// 	Subject   string `json:"sub,omitempty"`
// }

// 创建自己的Claims
type JwtClaims struct {
	*jwt.StandardClaims
	//用户编号
	Uid int64
	Username string
}

var (
	//盐
	secret []byte = []byte("yulibaozi.com")
)

// CreateJwtToken 生成一个jwttoken
func CreateJwtToken(id int64,username string) (string, error) {
	
	expireToken := time.Now().Add(time.Hour * 24).Unix()
	
	claims := JwtClaims{
		&jwt.StandardClaims{
		
			NotBefore: time.Now().Unix(),
		
			ExpiresAt: expireToken,
			
			Issuer: "yulibaozi",
		},
		id,
		username,
	}
	// 对自定义claims加密,jwt.SigningMethodHS256是加密算法得到第二部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 给这个token盐加密 第三部分,得到一个完整的三段的加密
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// DestoryJwtToken 删除JwtToken
func DestoryJwtToken() (string, error) {
	claims := JwtClaims{
		&jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 99998),
			ExpiresAt: int64(time.Now().Unix() - 99999),
			Issuer:    "yulibaozi",
		},
		-1,
		"",
	}
	// 对自定义claims加密,jwt.SigningMethodHS256是加密算法得到第二部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 给这个token盐加密 第三部分,得到一个完整的三段的加密
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// VerifyToken 得到一个JwtToken,然后验证是否合法,防止伪造
func VerifyJwtToken(jwtToken string) bool {
	// 解析、验证并返回一个令牌。keyFunc将收到解析令牌,应该返回验证的关键。如果一切都是干净的,error是空
	_, err := jwt.Parse(jwtToken, func(*jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return false
	}
	return true
}
// ParseJwtToken 解析token得到是自己创建的Claims
func ParseJwtToken(jwtToken string) (*JwtClaims, error) {
	var jwtclaim = &JwtClaims{}
	_, err := jwt.ParseWithClaims(jwtToken, jwtclaim, func(*jwt.Token) (interface{}, error) {
		//得到盐
		return secret, nil
	})
	if err != nil {
		return nil, errors.New("解析jwtToken失败")
	}
	return jwtclaim, nil
}
