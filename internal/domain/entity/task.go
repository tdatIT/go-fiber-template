package entity

import "time"

type Task struct {
	ID          int        `gorm:"id;primaryKey;type:bigint;not null" json:"id"`
	Title       string     `gorm:"title;type:varchar(255);not null" json:"title"`
	Description string     `gorm:"description;type:text;not null" json:"description"`
	Status      int        `gorm:"status;type:smallint;default:1" json:"status"`
	Deadline    *time.Time `gorm:"deadline;type:timestamp;" json:"deadline"`
	CreatedAt   time.Time  `gorm:"autoUpdateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"deleted_at;type:timestamp;" json:"deleted_at"` // Soft delete
}

func (*Task) TableName() string {
	return "tasks"
}
