package middlewares

import (
	"net/http"
	"strconv"
	"web-app/database"
	"web-app/helpers"
	"web-app/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"error":    "Bad Request",
				"messsage": "Invalid Parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		UserID := uint(userData["id"].(float64))
		Product := models.Product{}

		err = db.Select("user_id").First(&Product, uint(productId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doest'n exist",
			})
			return
		}

		if Product.UserID != UserID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
		}
		c.Next()
	}
}
