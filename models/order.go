package models

import "time"

type Order struct {
	ID           uint `gorm:"primaryKey"`
	CustomerName string
	OrderedAt    time.Time
	Items        []Item
}
