package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"pendaftaran-sidang/internal/model/entity"
	"strings"
)

type AuthConfig struct {
	Filter       func(*fiber.Ctx) error
	Unauthorized fiber.Handler
}

var secretKey = []byte("secret")

func UserAuthentication(c AuthConfig) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//db := config.OpenConnection()

		//check header
		header := ctx.Get("Authorization")
		if header == "" {
			return c.Unauthorized(ctx)
		}

		tokenString := strings.Replace(header, "Bearer ", "", 1)

		// check user token is valid
		userToken := entity.UserToken{}

		validateJWT, err := ValidateJWT(tokenString)
		if err != nil {
			fmt.Println(err)
		}

		//fmt.Println("type:", reflect.TypeOf(validateJWT["id"]))

		userToken.UserId = validateJWT["id"].(float64)
		userToken.Nama = validateJWT["nama"].(string)
		userToken.Role = validateJWT["role"].(string)
		userToken.Username = validateJWT["username"].(string)

		fmt.Println(userToken.Role)

		ctx.Locals("user", userToken.UserId)
		ctx.Locals("role", userToken.Role)

		return ctx.Next()
	}
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	return checkJWT(tokenString, "secret")
}

func checkJWT(tokenString string, secret string) (jwt.MapClaims, error) {
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	var JWT_SIGNATURE_KEY = []byte(secret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}