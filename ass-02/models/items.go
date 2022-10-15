package models

type Item struct {
	ID          uint64 `gorm:"column:item_id;primaryKey;autoIncrement:true" json:"itemId"`
	ItemCode    string `gorm:"column:item_code;not null;type:varchar(50)" json:"itemCode"`
	Description string `gorm:"column:description;not null;type:varchar(255)" json:"description"`
	Quantity    int64  `gorm:"column:quantity;not null;type:int" json:"quantity"`
	OrderId     uint64 `json:"order_id"`
}
