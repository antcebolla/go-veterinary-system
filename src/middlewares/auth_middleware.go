package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//obtain session token
		sessionToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if sessionToken == "" {
			log.Println("No session token found in Authorization header.")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"access": "unauthorized", "error": "No token provided"})
			return
		}
		// verify session token
		claims, err := jwt.Verify(c.Request.Context(), &jwt.VerifyParams{
			Token: sessionToken,
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"access": "unauthorized"})
			c.Abort() // Detiene la ejecución de los handlers siguientes
			return
		}
		// fetch user
		usr, err := user.Get(c.Request.Context(), claims.Subject)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
			c.Abort()
			return
		}

		// set user and claims in context
		c.Set("user", usr)
		c.Set("claims", claims)
		// continue to next handler
		c.Next()
	}
}
