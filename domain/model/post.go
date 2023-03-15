package model

import (
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID          uint          `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	UserId      uint          `json:"userId"`
	Description string        `json:"description"`
	Likers      pq.Int64Array `json:"likers" gorm:"type:integer[]"`
	Images      pq.Int64Array `json:"images" gorm:"type:integer[]"`
	Comments    []Comment     `json:"comments"`
	// Comments    pq.Int64Array `json:"comments" gorm:"type:integer[]"`
}
