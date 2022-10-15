package models

type Weather struct {
	ID    int64 `gorm:"column:id;primaryKey;autoIncrement:true"`
	Water int64 `gorm:"column:water;not null;type:int" json:"water"`
	Wind  int64 `gorm:"column:wind;not null;type:int" json:"wind"`
}
