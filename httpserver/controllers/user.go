package controllers

import (
	"net/http"
	"strconv"

	"github.com/deevarindu/final-project-2/helper/jwt"
	"github.com/deevarindu/final-project-2/httpserver/controllers/params"
	"github.com/deevarindu/final-project-2/httpserver/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserController) Login(ctx *gin.Context) {
	var req params.UserLoginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := u.svc.FindUserByEmail(req.Email)
	if user == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if bcryptErr != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	token, err := jwt.GenerateToken(*user.Id, user.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (u *UserController) Register(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	var req params.UserCreateRequest
	err := ctx.ShouldBindJSON(&req)
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

func (u *UserController) GetUser(ctx *gin.Context) {
	email := ctx.GetString("email")
	response := u.svc.GetUser(email)
	WriteJsonResponse(ctx, response)
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	var req params.UserUpdateRequest
	err := ctx.ShouldBindJSON(&req)
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
