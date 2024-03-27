package models

type Item struct {
	ID          uint `gorm:"primaryKey"`
	ItemCode    string
	Description string
	Quantity    int
	OrderID     uint
}
