package controllers

import (
	"fmt"
	helper "mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type (
	dataUserPhoto struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}

	dataPhoto struct {
		ID        uint64     `json:"id"`
		Title     string     `json:"title"`
		Caption   string     `json:"caption"`
		PhotoUrl  string     `json:"photo_url"`
		UserID    uint64     `json:"user_id"`
		CreatedAt *time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
		User      *dataUserPhoto
	}
)

func (db *Database) AddPhoto(c *gin.Context) {
	contentType := helper.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}

	if contentType == appJSON {
		c.ShouldBindHeader(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = uint64(userData["id"].(float64))
	Photo.CreatedAt = &timeNow
	err := db.Connect.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
	return
}

func (db *Database) GetAllPhotos(c *gin.Context) {
	var (
		photos []models.Photo
		result []dataPhoto
	)

	err := db.Connect.Model(photos).Preload("User").
		Find(&photos).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if len(photos) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":     nil,
			"count":    0,
			"meessage": "Data not found",
		})
	} else {
		for _, i := range photos {
			user := new(dataUserPhoto)
			user.Email = i.User.Email
			user.Username = i.User.Username

			var data dataPhoto
			data.ID = i.ID
			data.Title = i.Title
			data.Caption = i.Caption
			data.PhotoUrl = i.PhotoUrl
			data.UserID = i.UserID
			data.CreatedAt = i.CreatedAt
			data.UpdatedAt = i.UpdatedAt
			data.User = user

			result = append(result, data)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"count":  len(photos),
			"data":   result,
		})
	}
	return
}

func (db *Database) UpdatePhoto(c *gin.Context) {
	fmt.Println("hi")
	contentType := helper.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.ParseUint(c.Param("photoId"), 10, 64)

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UpdatedAt = &timeNow

	err := db.Connect.Model(&Photo).Preload("User").Where("id = ?", photoId).
		Updates(Photo).
		Error

	err = db.Connect.Model(&Photo).Preload("User").
		Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"updated_at": Photo.UpdatedAt,
	})
	return
}

func (db *Database) DeletePhoto(c *gin.Context) {
	contentType := helper.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.ParseUint(c.Param("photoId"), 10, 64)

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	Photo.ID = photoId
	Photo.UpdatedAt = &timeNow

	err := db.Connect.Where("id = ?", photoId).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been sucessfully deleted",
	})
	return
}
