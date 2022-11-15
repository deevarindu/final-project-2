package controllers

import (
	"net/http"
	"strconv"

	"github.com/deevarindu/final-project-2/httpserver/controllers/params"
	"github.com/deevarindu/final-project-2/httpserver/services"
	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	svc *services.PhotoSvc
}

func NewPhotoController(svc *services.PhotoSvc) *PhotoController {
	return &PhotoController{
		svc: svc,
	}
}

func (p *PhotoController) GetPhotos(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	response := p.svc.GetPhotos()
	WriteJsonResponse(ctx, response)
}

func (p *PhotoController) UploadPhoto(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	var req params.PhotoUploadRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := p.svc.UploadPhoto(&req)
	WriteJsonResponse(ctx, response)
}

func (p *PhotoController) UpdatePhoto(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	var req params.PhotoUpdateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photoID, _ := strconv.Atoi(ctx.Param("id"))
	response := p.svc.UpdatePhoto(photoID, &req)
	WriteJsonResponse(ctx, response)
}

func (p *PhotoController) DeletePhoto(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	response := p.svc.DeletePhoto(ctx.Param("id"))
	WriteJsonResponse(ctx, response)
}
