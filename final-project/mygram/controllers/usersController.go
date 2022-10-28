package controllers

import (
	helper "mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Database struct {
		Connect *gorm.DB
	}
)

var (
	timeNow = time.Now()
	appJSON = "application/json"
)

func (db *Database) UserRegister(c *gin.Context) {
	var user models.User
	contentType := helper.GetContentType(c)

	if contentType == appJSON {
		c.ShouldBindHeader(&user)
	} else {
		c.ShouldBind(&user)
	}

	user.CreatedAt = &timeNow

	err := db.Connect.Debug().Create(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"age":      user.Age,
	})
}

func (db *Database) UserLogin(c *gin.Context) {
	contentType := helper.GetContentType(c)
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindHeader(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Connect.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	comparePass := helper.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}
	token := helper.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (db *Database) UpdateUser(c *gin.Context) {
	contentType := helper.GetContentType(c)
	User := models.User{}

	userId, _ := strconv.ParseUint(c.Param("userId"), 10, 64)

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	User.ID = userId
	User.UpdatedAt = &timeNow

	err := db.Connect.Model(&User).Where("id = ?", userId).
		Updates(models.User{Email: User.Email, Username: User.Username}).
		Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"age":        User.Age,
		"updated_at": User.UpdatedAt,
	})
}

func (db *Database) DeleteUser(c *gin.Context) {
	contentType := helper.GetContentType(c)
	User := models.User{}

	userId, _ := strconv.ParseUint(c.Param("userId"), 10, 64)

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	User.ID = userId
	User.UpdatedAt = &timeNow

	err := db.Connect.Where("id = ?", userId).Delete(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been sucessfully deleted",
	})
}
