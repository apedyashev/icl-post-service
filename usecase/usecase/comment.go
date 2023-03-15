package usecase

import (
	"icl-posts/domain/model"
	"icl-posts/usecase/repository"
)

type CommentUsecase interface {
	Create(c *model.Comment) (*model.Comment, error)
}

type commentUsecase struct {
	commentRepo repository.CommentRepository
}

func NewCommentUsecase(commentRepo repository.CommentRepository) CommentUsecase {
	return &commentUsecase{
		commentRepo,
	}
}

func (cu *commentUsecase) Create(c *model.Comment) (*model.Comment, error) {
	return nil, nil
}
