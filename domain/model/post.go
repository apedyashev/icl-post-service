package model

import (
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	UserId      uint           `json:"userId" validate:"required,gte=1"`
	Description string         `json:"description"`
	Likers      pq.Int64Array  `json:"likers" gorm:"type:integer[]"`
	Images      pq.StringArray `json:"images" gorm:"type:varchar(64)[]"`
	Comments    []Comment      `json:"comments"`
}
