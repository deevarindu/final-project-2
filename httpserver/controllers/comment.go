package controllers

import (
	"net/http"
	"strconv"

	"github.com/deevarindu/final-project-2/httpserver/controllers/params"
	"github.com/deevarindu/final-project-2/httpserver/services"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	svc *services.CommentSvc
}

func NewCommentController(svc *services.CommentSvc) *CommentController {
	return &CommentController{
		svc: svc,
	}
}

func (c *CommentController) GetComments(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	response := c.svc.GetComments()
	WriteJsonResponse(ctx, response)
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	var req params.CommentCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := c.svc.CreateComment(&req)
	WriteJsonResponse(ctx, response)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	var req params.CommentUpdateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentID, _ := strconv.Atoi(ctx.Param("id"))
	response := c.svc.UpdateComment(commentID, &req)
	WriteJsonResponse(ctx, response)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	response := c.svc.DeleteComment(ctx.Param("id"))
	WriteJsonResponse(ctx, response)
}
