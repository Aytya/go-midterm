package domain

import "time"

type Todo struct {
	ID       string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" db:"id" json:"id"`
	Title    string    `gorm:"type:varchar(50);not null" db:"title" json:"title"`
	DateTime time.Time `gorm:"type:timestamp;not null" db:"datetime" json:"datetime"`
	ActiveAt time.Time `gorm:"type:timestamp;not null" db:"active_at" json:"active_at"`
	Status   bool      `gorm:"type:boolean;not null" db:"status" json:"status"`
	Priority string    `gorm:"type:varchar(50);not null" db:"priority" json:"priority"`
}
