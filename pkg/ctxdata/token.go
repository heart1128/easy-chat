package ctxdata

import "github.com/golang-jwt/jwt/v4"

// 生成jwt

const Identify = "heart.com" // 个人信息标识

// GetJwtToken
//
//	生成jwt三部分：Header.Payload.Signature
//	@Description:
//	@param secretKey 秘钥
//	@param iat		当前时间
//	@param seconds	要过期的时间
//	@param uid		信息标识
//	@return string
//	@return error
func GetJwtToken(secretKey string, iat, seconds int64, uid string) (string, error) {

	// payload部分
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds // 设置过期时间
	claims["iat"] = iat
	claims[Identify] = uid // 信息标识和uid

	// signature部分
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	// 根据秘钥生成
	return token.SignedString([]byte(secretKey))
}
