package controller

import (
	"errors"
	"fmt"
	"icl-posts/domain/model"
	"icl-posts/usecase/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type PostController interface {
	UserPosts(ctx echo.Context) error
	PostById(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	AddComment(ctx echo.Context) error
}

type postController struct {
	postUsecase    usecase.Post
	commentUsecase usecase.CommentUsecase
}

func NewPostController(postUsecase usecase.Post, commentUsecase usecase.CommentUsecase) PostController {
	return &postController{
		postUsecase,
		commentUsecase,
	}
}

func (pc *postController) Create(ctx echo.Context) error {
	log.Println("Create")
	var postData model.Post
	if err := ctx.Bind(&postData); err != nil {
		log.Println("post bind", err)
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	post, err := pc.postUsecase.Create(&postData)
	if err != nil {
		log.Println("post create error", err)
		return ctx.JSON(http.StatusBadRequest, nil)
	}
	return ctx.JSON(http.StatusCreated, post)
}

func (pc *postController) UserPosts(ctx echo.Context) error {
	fmt.Println("1")
	userIdStr := ctx.QueryParam("user_id")
	userId64, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, errors.New("user_id is malformed"))
	}

	fmt.Println("2")
	limitStr := ctx.QueryParam("limit")
	var limit64 int64
	limit64 = -1
	if limitStr != "" {
		limit64, err = strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, errors.New("limit is malformed"))
		}
	}

	fmt.Println("3")
	posts, err := pc.postUsecase.UserPosts(uint(userId64), int(limit64))
	if err != nil {
		fmt.Println("error in user posts controller", err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	fmt.Printf("post %+v", posts)
	return ctx.JSON(http.StatusOK, posts)
}

func (pc *postController) PostById(ctx echo.Context) error {
	postIdStr := ctx.Param("id")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	post, err := pc.postUsecase.PostById(uint(postId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, nil)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, post)
}

func (pc *postController) AddComment(ctx echo.Context) error {
	postIdStr := ctx.Param("id")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	var commentData model.Comment
	if err := ctx.Bind(&commentData); err != nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	commentData.PostId = uint(postId)

	comment, err := pc.postUsecase.AddComment(uint(postId), &commentData)
	// comment, err := pc.commentUsecase.Create(&commentData)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "cannot create comment")
	}
	return ctx.JSON(http.StatusOK, comment)
}

func (pc *postController) Update(ctx echo.Context) error {
	postIdStr := ctx.Param("id")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	var postData model.Post
	if err := ctx.Bind(&postData); err != nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	post, err := pc.postUsecase.Update(uint(postId), &postData)
	return ctx.JSON(http.StatusOK, post)
}
