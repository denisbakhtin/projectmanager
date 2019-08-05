package controllers

import (
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
						c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
							"error": "Invalid token signing method, please, login again",
						})
					}
					return []byte(config.Settings.JWTSecret), nil
				})
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error": "Invalid authentication token, please, login again",
					})
				}

				if claims, ok := token.Claims.(*models.JWTClaims); ok && token.Valid {
					models.DB.Where("email = ?", claims.Subject).First(&user)
				}
			}
		}
		if user.ID != 0 && user.Status == models.ACTIVE {
			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Please login to make this request",
			})
		}
	}
}
