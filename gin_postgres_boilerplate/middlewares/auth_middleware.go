package middlewares

import (
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/config"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/constants"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/repositories"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/utils"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			utils.UnauthorizedResponse(c, constants.ErrorCodes["EA004"])
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 || tokenString[0] != "Bearer" {
			utils.UnauthorizedResponse(c, constants.ErrorCodes["EA005"])
			c.Abort()
			return
		}

		signedToken := tokenString[1]
		token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			utils.UnauthorizedResponse(c, constants.ErrorCodes["EA006"])
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.UnauthorizedResponse(c, constants.ErrorCodes["EA007"])
			c.Abort()
			return
		}

		db := config.DB
		user, err := repositories.NewUserRepository(db).GetUserByEmail(claims["email"].(string))
		if err != nil {
			utils.UnauthorizedResponse(c, constants.ErrorCodes["EU003"])
			c.Abort()
			return
		}

		c.Set("authUser", user)
		c.Next()
	}
}
