package services

import (
	"database/sql"

	"github.com/deevarindu/final-project-2/httpserver/controllers/params"
	"github.com/deevarindu/final-project-2/httpserver/repositories"
	"github.com/deevarindu/final-project-2/httpserver/repositories/models"
	"github.com/deevarindu/final-project-2/httpserver/views"
)

type SocialMediaSvc struct {
	repo repositories.SocialMediaRepository
}

func NewSocialMediaSvc(repo repositories.SocialMediaRepository) *SocialMediaSvc {
	return &SocialMediaSvc{
		repo: repo,
	}
}

func (s *SocialMediaSvc) GetSocialMedias() *views.Response {
	socialMedias, err := s.repo.GetSocialMedias()
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFoundResponse(err)
		}
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessGetResponse(parseModelToGetSocialMedias(socialMedias), "Success get all social medias")
}

func parseModelToGetSocialMedias(mod *[]models.SocialMedia) *[]views.GetSocialMedias {
	var s []views.GetSocialMedias
	for _, v := range *mod {
		s = append(s, views.GetSocialMedias{
			Id:             *v.Id,
			Name:           v.Name,
			SocialMediaUrl: v.SocialMediaUrl,
			UserId:         v.UserId,
		})
	}
	return &s
}

func (s *SocialMediaSvc) AddSocialMedia(req *params.SocialMediaAddRequest) *views.Response {
	socialMedia := parseRequestToModelAddSocialMedia(req)

	id := 1
	if socialMedias, err := s.repo.GetSocialMedias(); err == nil {
		id = len(*socialMedias) + 1
	}
	socialMedia.Id = &id

	if err := s.repo.AddSocialMedia(socialMedia); err != nil {
		return views.InternalServerErrorResponse(err)
	}
	return views.SuccessCreateResponse(socialMedia, "Success add social media")
}

func parseRequestToModelAddSocialMedia(req *params.SocialMediaAddRequest) *models.SocialMedia {
	return &models.SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
	}
}

func (s *SocialMediaSvc) UpdateSocialMedia(id int, req *params.SocialMediaUpdateRequest) *views.Response {
	socialMedia := parseRequestToModelUpdateSocialMedia(req)
	socialMedia.Id = &id

	if err := s.repo.UpdateSocialMedia(socialMedia); err != nil {
		return views.InternalServerErrorResponse(err)
	}
	return views.SuccessUpdateResponse(socialMedia, "Success update social media")
}

func parseRequestToModelUpdateSocialMedia(req *params.SocialMediaUpdateRequest) *models.SocialMedia {
	return &models.SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
	}
}

func (s *SocialMediaSvc) DeleteSocialMedia(id string) *views.Response {
	if err := s.repo.DeleteSocialMedia(id); err != nil {
		return views.InternalServerErrorResponse(err)
	}
	return views.SuccessDeleteResponse("Your social media has been deleted")
}
