package controllers

import (
	"net/http"
	"strconv"

	"github.com/deevarindu/final-project-2/httpserver/controllers/params"
	"github.com/deevarindu/final-project-2/httpserver/services"
	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	svc *services.SocialMediaSvc
}

func NewSocialMediaController(svc *services.SocialMediaSvc) *SocialMediaController {
	return &SocialMediaController{
		svc: svc,
	}
}

func (s *SocialMediaController) GetSocialMedias(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	response := s.svc.GetSocialMedias()
	WriteJsonResponse(ctx, response)
}

func (s *SocialMediaController) AddSocialMedia(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	var req params.SocialMediaAddRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := s.svc.AddSocialMedia(&req)
	WriteJsonResponse(ctx, response)
}

func (s *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	var req params.SocialMediaUpdateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	socialMediaID, _ := strconv.Atoi(ctx.Param("id"))
	response := s.svc.UpdateSocialMedia(socialMediaID, &req)
	WriteJsonResponse(ctx, response)
}

func (s *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	response := s.svc.DeleteSocialMedia(ctx.Param("id"))
	WriteJsonResponse(ctx, response)
}
