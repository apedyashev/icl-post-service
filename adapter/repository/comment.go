package repository

import (
	"icl-posts/domain/model"
	"icl-posts/usecase/repository"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &commentRepository{db}
}

func (cr *commentRepository) CommentsForPost(pid uint, limit uint) ([]*model.Comment, error) {
	return nil, nil
}

func (cr *commentRepository) Update(id uint, c *model.Comment) error {
	return nil
}

func (cr *commentRepository) Delete(id uint) error {
	return nil
}

func (cr *commentRepository) Create(c *model.Comment) (*model.Comment, error) {
	if err := cr.db.Create(c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
