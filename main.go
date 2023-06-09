package main

import (
	"fmt"
	"icl-posts/domain/model"
	"icl-posts/infrastructure/datastore"
	"icl-posts/infrastructure/httperror"
	"icl-posts/infrastructure/router"
	"icl-posts/infrastructure/validator"
	"icl-posts/registry"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	db := datastore.NewDB()

	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.Comment{})
	log.Println("Database Migration Completed!")

	r := registry.NewRegistry(db)

	e := echo.New()

	e.Validator = validator.New()
	e = router.NewRouter(e, r.NewAppController())
	e.HTTPErrorHandler = httperror.Handler

	fmt.Println("Server listen at http://localhost:80")
	if err := e.Start(":80"); err != nil {
		log.Fatalln(err)
	}
}
