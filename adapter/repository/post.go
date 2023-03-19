package repository

import (
	"fmt"
	"icl-posts/domain/model"
	"icl-posts/usecase/repository"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.PostRepository {
	return &postRepository{db}
}

func (pr *postRepository) Create(p *model.Post) error {
	if err := pr.db.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) UserPosts(uid uint, limit int) ([]model.Post, error) {
	posts := []model.Post{}
	fmt.Println("limit", limit)
	// NOTE: if associations population would be needed, then Preload("Comments") could help
	err := pr.db.Where("user_id = ?", uid).
		Limit(limit).
		Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (pr *postRepository) PostById(id uint) (*model.Post, error) {
	var post model.Post
	err := pr.db.Preload("Comments").Where("id = ?", id).First(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (pr *postRepository) Update(id uint, changes *model.Post) (*model.Post, error) {
	var existingPost model.Post
	existingPost.ID = id
	err := pr.db.
		Preload("Comments").
		Where("id = ?", id).
		Find(&existingPost).Error

	if err != nil {
		return nil, err
	}

	log.Printf("Updating, changes %+v", changes)
	err = pr.db.Model(&existingPost).
		// IMPORTANT: don't allow to update assocoiations
		Omit(clause.Associations).
		Updates(changes).Error
	if err != nil {
		log.Println("Update error", err)
		return nil, err
	}

	return &existingPost, nil
}

func (pr *postRepository) Delete(id uint) error {
	return nil
}

func (pr *postRepository) AddComment(postId uint, c *model.Comment) (*model.Comment, error) {
	var post model.Post
	err := pr.db.Where("id = ?", postId).First(&post).Association("Comments").Append(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
