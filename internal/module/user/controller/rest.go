package controller

import (
	"backend/internal/module/user/entity"
	"backend/internal/module/user/interfaces"
	"backend/internal/module/user/repository"
	"backend/internal/module/user/service"
	"backend/pkg/handler"
	"backend/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userController struct {
	service interfaces.UserService
}

func NewUserController(db *gorm.DB) *userController {
	var (
		controller = new(userController)
		repo       = repository.NewUserRepository(db)
		service    = service.NewUserService(repo)
	)
	controller.service = service

	return controller
}

func (c *userController) Register(router *gin.Engine) {
	v1 := router.Group("/v1")

	v1.GET("/users/:uuid", func(ctx *gin.Context) {
		var req entity.UserReqByUUID
		if !validator.BindUri(ctx, &req) {
			return
		}

		res, err := c.service.GetUser(&req)
		if err != nil {
			handler.Error(ctx, http.StatusNotFound, err.Error())
			return
		}

		handler.Success(ctx, http.StatusOK, "Success getting users", &res)
	})
}
