package repository

import "icl-posts/domain/model"

type PostRepository interface {
	UserPosts(uid uint, limit int) ([]model.Post, error)
	PostById(id uint) (*model.Post, error)
	Create(p *model.Post) error
	Update(id uint, p *model.Post) (*model.Post, error)
	Delete(id uint) error
	AddComment(postId uint, c *model.Comment) (*model.Comment, error)
}
