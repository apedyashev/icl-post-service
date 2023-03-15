package router

import (
	"icl-posts/adapter/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, ac controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/posts", func(context echo.Context) error { return ac.Post.UserPosts(context) })
	e.GET("/posts/:id", func(context echo.Context) error { return ac.Post.PostById(context) })
	e.POST("/posts", func(context echo.Context) error { return ac.Post.Create(context) })
	e.POST("/posts/:id/comments", func(context echo.Context) error { return ac.Post.AddComment(context) })
	e.PUT("/posts/:id", func(context echo.Context) error { return ac.Post.Update(context) })

	return e
}
