package controllers

import (
	"ass-02/helper"
	"ass-02/models"
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

	CreateItem struct {
		ItemCode    string `json:"itemCode"`
		Description string `json:"description"`
		Quantity    int64  `json:"quantity"`
	}

	CreateOrder struct {
		OrderedAt    string       `json:"orderedAt"`
		CustomerName string       `json:"customerName"`
		Items        []CreateItem `json:"items"`
	}

	UpdateItem struct {
		LineItemId  uint64 `json:"lineItemId"`
		ItemCode    string `json:"itemCode"`
		Description string `json:"description"`
		Quantity    int64  `json:"quantity"`
	}

	UpdateOrder struct {
		OrderedAt    string       `json:"orderedAt"`
		CustomerName string       `json:"customerName"`
		Items        []UpdateItem `json:"items"`
	}

	GetRespOrders struct {
		OrderID      uint64       `json:"orderId"`
		OrderedAt    string       `json:"orderedAt"`
		CustomerName string       `json:"customerName"`
		Items        []UpdateItem `json:"items"`
	}
)

var (
	timeNow = time.Now()
	appJSON = "application/json"
)

func initCreateDataItem(itemCode, description string, quantity int) models.Item {
	item := models.Item{
		ItemCode:    itemCode,
		Description: description,
		Quantity:    int64(quantity),
	}
	return item
}

func (db *Database) CreateOrder(c *gin.Context) {
	var (
		order models.Order
		items []CreateItem
	)
	contentType := helper.GetContentType(c)

	if contentType == appJSON {
		c.ShouldBindJSON(&order)
	} else {
		c.ShouldBind(&order)
	}

	err := helper.ValidateOrder("create", order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err,
		})
		return
	}
	order.OrderedAt = timeNow

	err = db.Connect.Create(&order).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
	} else {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": err.Error(),
			})
		} else {
			for _, i := range order.Items {
				newItem := new(CreateItem)
				newItem.ItemCode = i.ItemCode
				newItem.Description = i.Description
				newItem.Quantity = i.Quantity
				items = append(items, *newItem)
			}
			result := CreateOrder{
				OrderedAt:    order.OrderedAt.Format("2006-01-02 15:04:05"),
				CustomerName: order.CustomerName,
				Items:        items,
			}
			c.JSON(http.StatusOK, gin.H{
				"result": "Data has been created!",
				"data":   result,
			})
		}
	}
	return
}

func (db *Database) GetAllOrders(c *gin.Context) {
	var (
		orders []models.Order
		result []GetRespOrders
	)

	err := db.Connect.Model(orders).Preload("Items").
		Order("created_at desc").
		Find(&orders).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	if len(orders) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":   nil,
			"count":  0,
			"result": "Data not found",
		})
	} else {
		for _, order := range orders {
			var items []UpdateItem
			for _, i := range order.Items {
				if i.OrderId == order.ID {
					var newItem UpdateItem
					newItem.LineItemId = i.ID
					newItem.ItemCode = i.ItemCode
					newItem.Description = i.Description
					newItem.Quantity = i.Quantity
					items = append(items, newItem)
				}
			}
			data := GetRespOrders{
				OrderID:      order.ID,
				OrderedAt:    order.OrderedAt.Format("2006-01-02 15:04:05"),
				CustomerName: order.CustomerName,
				Items:        items,
			}
			result = append(result, data)
		}
		c.JSON(http.StatusOK, gin.H{
			"data":  result,
			"count": len(result),
		})
	}
	return
}

func (db *Database) UpdateOrder(c *gin.Context) {
	var (
		order       models.Order
		item        models.Item
		items       []UpdateItem
		updateOrder UpdateOrder
	)
	contentType := helper.GetContentType(c)

	orderId, _ := strconv.Atoi(c.Param("orderId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&updateOrder)
	} else {
		c.ShouldBind(&updateOrder)
	}

	err := db.Connect.Where("order_id = ?", orderId).First(&order).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "Data not found",
		})
	} else {
		err = db.Connect.Model(&order).
			UpdateColumn("customer_name", updateOrder.CustomerName).
			Error
		for _, i := range updateOrder.Items {
			item = models.Item{
				ID:          i.LineItemId,
				ItemCode:    i.ItemCode,
				Description: i.Description,
				Quantity:    i.Quantity,
			}
			err = db.Connect.Model(&item).
				UpdateColumns(item).
				Error
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": "Update Failed",
			})
		} else {
			for _, i := range updateOrder.Items {
				var newItem UpdateItem
				newItem.LineItemId = i.LineItemId
				newItem.ItemCode = i.ItemCode
				newItem.Description = i.Description
				newItem.Quantity = i.Quantity
				items = append(items, newItem)
			}
			result := UpdateOrder{
				OrderedAt:    order.OrderedAt.Format("2006-01-02 15:04:05"),
				CustomerName: order.CustomerName,
				Items:        items,
			}
			c.JSON(http.StatusOK, gin.H{
				"result": "Data has been updated!",
				"data":   result,
			})
		}
	}
	return
}

func (db *Database) DeleteOrder(c *gin.Context) {
	var (
		order models.Order
		item  models.Item
	)

	orderId := c.Param("orderId")

	err := db.Connect.Where("order_id = ?", orderId).First(&order).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "Data not found",
		})
	} else {
		err = db.Connect.Where("order_id = ?", orderId).Delete(&item).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": "Failed delete data",
			})
		} else {
			err = db.Connect.Delete(&order).Error
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"result": "Failed delete data",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"result": "Data has been deleted",
				})
			}
		}
	}
	return
}
