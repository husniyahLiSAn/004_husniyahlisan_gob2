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
	dataUser struct {
		ID       uint64 `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}

	dataPhotoComment struct {
		ID       uint64 `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoUrl string `json:"photo_url"`
		UserID   uint64 `json:"user_id"`
	}

	dataComment struct {
		ID        uint64     `json:"id"`
		Message   string     `json:"message"`
		PhotoID   uint64     `json:"photo_id"`
		UserID    uint64     `json:"user_id"`
		CreatedAt *time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
		User      *dataUser
		Photo     *dataPhoto
	}
)

func (db *Database) AddComment(c *gin.Context) {
	var Comment models.Comment
	contentType := helper.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	if contentType == appJSON {
		c.ShouldBindHeader(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	fmt.Println(Comment)
	Comment.UserID = uint64(userData["id"].(float64))
	Comment.CreatedAt = &timeNow

	err := db.Connect.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func (db *Database) GetAllComments(c *gin.Context) {
	var (
		Comments []models.Comment
		result   []dataComment
	)

	err := db.Connect.Model(Comments).
		Preload("Photo").
		Preload("User").
		Find(&Comments).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for _, i := range Comments {
		user := new(dataUser)
		user.ID = i.User.ID
		user.Email = i.User.Email
		user.Username = i.User.Username

		photo := new(dataPhoto)
		photo.ID = i.Photo.ID
		photo.Title = i.Photo.Title
		photo.Caption = i.Photo.Caption
		photo.PhotoUrl = i.Photo.PhotoUrl
		photo.UserID = i.Photo.UserID

		var data dataComment
		data.ID = i.ID
		data.Message = i.Message
		data.PhotoID = i.PhotoID
		data.UserID = i.UserID
		data.CreatedAt = i.CreatedAt
		data.UpdatedAt = i.UpdatedAt
		data.User = user
		data.Photo = photo

		result = append(result, data)
	}
	if len(Comments) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":    nil,
			"count":   0,
			"message": "Data not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"count":  len(result),
			"data":   result,
		})
	}
	return
}

func (db *Database) UpdateComment(c *gin.Context) {
	contentType := helper.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.ParseUint(c.Param("commentId"), 10, 64)

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	Comment.ID = commentId
	Comment.UpdatedAt = &timeNow

	err := db.Connect.Model(&Comment).Preload("Photo").Preload("User").Where("id = ?", commentId).
		Updates(Comment).
		Error
	err = db.Connect.Model(&Comment).
		Preload("Photo").
		Preload("User").
		Find(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"title":      Comment.Photo.Title,
		"message":    Comment.Message,
		"user_id":    Comment.UserID,
		"photo_id":   Comment.PhotoID,
		"updated_at": Comment.UpdatedAt,
	})
}

func (db *Database) DeleteComment(c *gin.Context) {
	contentType := helper.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.ParseUint(c.Param("commentId"), 10, 64)

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	Comment.ID = commentId
	Comment.UpdatedAt = &timeNow

	err := db.Connect.Where("id = ?", commentId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been sucessfully deleted",
	})
}
