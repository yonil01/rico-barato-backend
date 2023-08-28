package middleware

import (
	"backend-ccff/internal/env"
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
	"crypto/rsa"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"io/ioutil"
)

var (
	verifyKey *rsa.PublicKey
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	Expires int64
	Ip      string
}

// JWT personzalizado
type jwtCustomClaims struct {
	User      *models.User `json:"user"`
	IPAddress string       `json:"ip_address"`
	jwt.StandardClaims
}

// init lee los archivos de firma y validación RSA
func init() {
	c := env.NewConfiguration()

	verifyBytes, err := ioutil.ReadFile(c.App.RSAPublicKey)
	if err != nil {
		logger.Error.Printf("leyendo el archivo público de confirmación: %s", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		logger.Error.Printf("realizando el parse en jwt RSA public: %s", err)
	}
}

func JWTProtected() fiber.Handler {
	// Create config for JWT authentication middleware.
	config := jwtware.Config{
		ErrorHandler:  jwtError,
		SigningKey:    verifyKey,
		SigningMethod: "RS256",
	}
	return jwtware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func GetUser(c *fiber.Ctx) (*models.User, error) {
	bearer := c.Get("Authorization")
	tkn := bearer[7:]

	var u *models.User
	verifyFunction := func(tkn *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	}

	token, err := jwt.ParseWithClaims(tkn, &jwtCustomClaims{}, verifyFunction)
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				logger.Warning.Printf("token expirado: %v", err)
				return u, err
			default:
				logger.Warning.Printf("Error de validacion del token: %v", err)
				return u, err
			}
		default:
			logger.Warning.Printf("Error al procesar el token: %v", err)
			return u, err
		}
	}
	u = token.Claims.(*jwtCustomClaims).User
	if !token.Valid {
		logger.Warning.Printf("Token no Valido: %v", err)
		return u, fmt.Errorf("Token no Valido")
	}
	/*if c.IP() != u.RealIP {
		logger.Warning.Printf("Token creado en un origen diferente : %v", err)
		return u, fmt.Errorf("Token creado en un origen diferente")
	}*/
	return u, nil
}
