package registry

import (
	"icl-posts/adapter/controller"

	"gorm.io/gorm"
)

type ControllerRegistry interface {
	NewAppController() controller.AppController
}

type controllerRegistry struct {
	db *gorm.DB
}

func NewRegistry(db *gorm.DB) ControllerRegistry {
	return &controllerRegistry{db}
}

func (cr *controllerRegistry) NewAppController() controller.AppController {
	return controller.AppController{
		Post: cr.NewPostController(),
	}
}
