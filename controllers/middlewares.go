package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthRequired middleware restricts access for authenticated users only
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		user := models.User{}
		bearer := c.GetHeader("Authorization")

		if strings.Contains(bearer, "Bearer ") {
			ss := strings.SplitAfter(bearer, " ")
			if len(ss) == 2 {
				tokenString := ss[1]
				token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						abortWithError(c, http.StatusUnauthorized, fmt.Errorf("Invalid token signing method, please, login again"))
					}
					return []byte(config.Settings.JWTSecret), nil
				})
				if err != nil {
					abortWithError(c, http.StatusUnauthorized, fmt.Errorf("Invalid authentication token, please, login again"))
				}

				if claims, ok := token.Claims.(*models.JWTClaims); ok && token.Valid {
					user, _ = models.UsersDB.GetByEmail(claims.Subject)
				}
			}
		}
		if user.ID != 0 && user.Status == models.ACTIVE {
			c.Set("user", user)
			c.Next()
		} else {
			abortWithError(c, http.StatusUnauthorized, fmt.Errorf("Please login to make this request"))
		}
	}
}

//AdminRequired middleware restricts access for authenticated admin users only
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		user := models.User{}
		if u, exists := c.Get("user"); exists {
			user = u.(models.User)
		}
		if user.ID != 0 && user.IsAdmin() {
			c.Next()
		} else {
			abortWithError(c, http.StatusUnauthorized, fmt.Errorf("Admin user is required to proceed"))
		}
	}
}

//LogErrors middleware logs all application errors to logs/%env%.log file
func LogErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			log.Printf("Error: %s, URL: %s, Agent: %s\n", err, c.Request.URL.Path, c.Request.UserAgent())
		}
	}
}
