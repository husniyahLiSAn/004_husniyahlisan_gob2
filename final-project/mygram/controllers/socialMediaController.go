package controllers

import (
	helper "mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type (
	dataSocialMedia struct {
		ID             uint64     `json:"id"`
		Name           string     `json:"email"`
		SocialMediaUrl string     `json:"social_media_url"`
		UserID         uint64     `json:"user_id"`
		CreatedAt      *time.Time `json:"created_at"`
		UpdatedAt      *time.Time `json:"updated_at"`
		User           *dataUser
	}
)

func (db *Database) CreateSosmed(c *gin.Context) {
	contentType := helper.GetContentType(c)
	SocialMedia := models.SocialMedia{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	if contentType == appJSON {
		c.ShouldBindHeader(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = uint64(userData["id"].(float64))
	SocialMedia.CreatedAt = &timeNow
	err := db.Connect.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.Url,
		"user_id":          SocialMedia.UserID,
		"created_at":       SocialMedia.CreatedAt,
	})
}

func (db *Database) GetSosmed(c *gin.Context) {
	var (
		SocialMedias []models.SocialMedia
		result       []dataSocialMedia
	)

	err := db.Connect.Model(SocialMedias).
		Preload("User").
		Find(&SocialMedias).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if len(SocialMedias) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":   nil,
			"count":  0,
			"result": "Data not found",
		})
	} else {
		for _, i := range SocialMedias {
			user := new(dataUser)
			user.ID = i.User.ID
			user.Email = i.User.Email
			user.Username = i.User.Username

			var data dataSocialMedia
			data.ID = i.ID
			data.Name = i.Name
			data.SocialMediaUrl = i.Url
			data.UserID = i.UserID
			data.CreatedAt = i.CreatedAt
			data.UpdatedAt = i.UpdatedAt
			data.User = user

			result = append(result, data)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"count":  len(SocialMedias),
			"data":   result,
		})
	}
	return
}

func (db *Database) UpdateSosmed(c *gin.Context) {
	contentType := helper.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.ParseUint(c.Param("socialMediaId"), 10, 64)

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}
	SocialMedia.ID = socialMediaId
	SocialMedia.UpdatedAt = &timeNow

	err := db.Connect.Model(&SocialMedia).Where("id = ?", socialMediaId).
		Updates(SocialMedia).
		Error
	err = db.Connect.Model(&SocialMedia).
		Preload("User").
		Find(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.Url,
		"user_id":          SocialMedia.UserID,
		"updated_at":       SocialMedia.UpdatedAt,
	})
}

func (db *Database) DeleteSosmed(c *gin.Context) {
	contentType := helper.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.ParseUint(c.Param("socialMediaId"), 10, 64)

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}
	SocialMedia.ID = socialMediaId

	err := db.Connect.Where("id = ?", socialMediaId).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been sucessfully deleted",
	})
}
