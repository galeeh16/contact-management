package entity

import "time"

type Contact struct {
	ID        uint64     `gorm:"primaryKey;column:id;autoIncrement"`
	FirstName string     `gorm:"column:first_name"`
	LastName  string     `gorm:"column:last_name"`
	Phone     string     `gorm:"column:phone"`
	UserID    uint64     `gorm:"column:user_id;"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Contact) TableName() string {
	return "contacts"
}
