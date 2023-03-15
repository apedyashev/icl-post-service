package repository

import "icl-posts/domain/model"

type CommentRepository interface {
	CommentsForPost(pid uint, limit uint) ([]*model.Comment, error)
	Update(id uint, c *model.Comment) error
	Delete(id uint) error
	Create(c *model.Comment) (*model.Comment, error)
}
