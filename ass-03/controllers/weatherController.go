package controllers

import (
	"ass-03/models"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Database struct {
		Connect *gorm.DB
	}

	Weather struct {
		Water int64 `json:"water"`
		Wind  int64 `json: "wind"`
	}
)

func (db *Database) CreateData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")

	water := c.PostForm("water")
	wind := c.PostForm("wind")

	intWater, _ := strconv.ParseInt(water, 0, 64)
	intWind, _ := strconv.ParseInt(wind, 0, 64)

	dataWeather := Weather{
		Water: intWater,
		Wind:  intWind,
	}

	err := db.Connect.Debug().Create(&dataWeather).Error

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"result": "Created data failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": dataWeather,
		})
	}
	return
}

func (db *Database) GetLatestData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

	var (
		data []models.Weather
	)

	db.Connect.Debug().Find(&data).Take(&data).Last(&data)
	fmt.Println("::", data)

	if len(data) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
	} else {
		var waterStatus, windStatus string

		dataWeather := Weather{
			Water: data[0].Water,
			Wind:  data[0].Wind,
		}

		if dataWeather.Water <= 5 {
			waterStatus = "Aman"
		} else if dataWeather.Water >= 6 && dataWeather.Water <= 8 {
			waterStatus = "Siaga"
		} else if dataWeather.Water > 8 {
			waterStatus = "Bahaya"
		}

		if dataWeather.Wind <= 6 {
			windStatus = "Aman"
		} else if dataWeather.Wind >= 7 && dataWeather.Wind <= 15 {
			windStatus = "Siaga"
		} else if dataWeather.Wind > 15 {
			windStatus = "Bahaya"
		}

		c.JSON(http.StatusOK, gin.H{
			"status":       dataWeather,
			"waterMessage": waterStatus,
			"windMessage":  windStatus,
		})

	}
}

func (db *Database) Update(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")

	water := rand.Intn(100)
	wind := rand.Intn(100)

	dataWeather := Weather{
		Water: int64(water),
		Wind:  int64(wind),
	}

	err := db.Connect.Debug().Create(&dataWeather).Error

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"result": "Created data failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": dataWeather,
		})
	}
	return
}
