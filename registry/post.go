package registry

import (
	"icl-posts/adapter/controller"
	"icl-posts/adapter/repository"
	"icl-posts/usecase/usecase"
)

func (cr *controllerRegistry) NewPostController() controller.PostController {
	commentRepo := repository.NewCommentRepository(cr.db)
	pu := usecase.NewPostUsecase(
		repository.NewPostRepository(cr.db),
		commentRepo,
	)
	c := usecase.NewCommentUsecase(
		commentRepo,
	)

	return controller.NewPostController(pu, c)
}
