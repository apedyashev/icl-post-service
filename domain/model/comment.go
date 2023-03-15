package model

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	PostId    uint      `json:"postId"`
	AuthorId  uint      `json:"authorId"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
