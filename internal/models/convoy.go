package models

import (
	"time"
)

type Convoy struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"type:varchar(255);not null" json:"title"`
	StartTime time.Time  `gorm:"not null" json:"start_time"`
	EndTime   *time.Time `gorm:"default:null" json:"end_time"` // Nullable
	IsPrivate bool       `gorm:"default:false" json:"is_private"`
	CreatedBy uint       `gorm:"index" json:"created_by"` // Foreign key to users(id)
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
}
