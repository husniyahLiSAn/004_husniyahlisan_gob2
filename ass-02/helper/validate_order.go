package helper

import (
	"ass-02/models"
	"errors"
	"strings"
)

func ValidateOrder(action string, order models.Order) error {
	switch strings.ToLower(action) {
	case "create":
		if len(order.CustomerName) == 0 {
			return errors.New("Required Customer Name")
		}
		if len(order.Items) == 0 {
			return errors.New("Required Items")
		} else {
			for _, i := range order.Items {
				if len(i.ItemCode) == 0 {
					return errors.New("Required Customer Name")
				}
				if i.Quantity < 0 {
					return errors.New("Quantity item minimal 1")
				}
			}
		}
		return nil
	case "update":
		if len(order.CustomerName) == 0 {
			return errors.New("Required Customer Name")
		}
		return nil
	}
	return nil
}
