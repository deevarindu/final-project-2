package controllers

import (
	"net/http"
	"strconv"

	"github.com/deevarindu/final-project-2/httpserver/controllers/params"
	"github.com/deevarindu/final-project-2/httpserver/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	svc *services.UserSvc
}

func NewUserController(svc *services.UserSvc) *UserController {
	return &UserController{
		svc: svc,
	}
}

func (u *UserController) GetUsers(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	response := u.svc.GetUsers()
	WriteJsonResponse(ctx, response)
}

func (u *UserController) GetUser(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	response := u.svc.GetUser(ctx.Param("id"))
	WriteJsonResponse(ctx, response)
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	var req params.UserCreateRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := u.svc.CreateUser(&req)
	WriteJsonResponse(ctx, response)
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	var req params.UserUpdateRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := strconv.Atoi(ctx.Param("id"))
	response := u.svc.UpdateUser(userID, &req)
	WriteJsonResponse(ctx, response)
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	response := u.svc.DeleteUser(ctx.Param("id"))
	WriteJsonResponse(ctx, response)
}
