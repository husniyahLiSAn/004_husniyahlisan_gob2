package models

import "time"

type Order struct {
	ID           uint64    `gorm:"column:order_id;primaryKey;autoIncrement:true" json:"orderId"`
	CustomerName string    `gorm:"column:customer_name;not null;type:varchar(150)" json:"customerName"`
	OrderedAt    time.Time `gorm:"type:timestamp;column:created_at" json:"orderedAt"`
	Items        []Item    `json:"items"`
}
