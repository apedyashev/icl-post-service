package usecase

import (
	"icl-posts/domain/model"
	"icl-posts/usecase/repository"
)

type Post interface {
	UserPosts(uid uint, limit int) ([]model.Post, error)
	PostById(id uint) (*model.Post, error)
	Create(post *model.Post) (*model.Post, error)
	Update(id uint, p *model.Post) (*model.Post, error)
	Delete(id uint) error
	AddComment(postId uint, c *model.Comment) (*model.Comment, error)
}

type postUsecase struct {
	postRepo    repository.PostRepository
	commentRepo repository.CommentRepository
}

func NewPostUsecase(postRepo repository.PostRepository, commentRepo repository.CommentRepository) Post {
	return &postUsecase{
		postRepo,
		commentRepo,
	}
}

func (pu *postUsecase) Create(post *model.Post) (*model.Post, error) {
	err := pu.postRepo.Create(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (pu *postUsecase) UserPosts(uid uint, limit int) ([]model.Post, error) {
	pp, err := pu.postRepo.UserPosts(uid, limit)
	if err != nil {
		return nil, err
	}

	return pp, nil
}

func (pu *postUsecase) PostById(id uint) (*model.Post, error) {
	return pu.postRepo.PostById(id)
}

func (pu *postUsecase) Update(id uint, p *model.Post) (*model.Post, error) {
	return pu.postRepo.Update(id, p)
}

func (pu *postUsecase) Delete(id uint) error {
	return nil
}

func (pu *postUsecase) AddComment(postId uint, c *model.Comment) (*model.Comment, error) {
	c, err := pu.commentRepo.Create(c)
	if err != nil {
		return nil, err
	}

	return pu.postRepo.AddComment(postId, c)
}
