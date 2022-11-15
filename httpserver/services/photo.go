package services

import (
	"database/sql"

	"github.com/deevarindu/final-project-2/httpserver/controllers/params"
	"github.com/deevarindu/final-project-2/httpserver/repositories"
	"github.com/deevarindu/final-project-2/httpserver/repositories/models"
	"github.com/deevarindu/final-project-2/httpserver/views"
)

type PhotoSvc struct {
	repo repositories.PhotoRepository
}

func NewPhotoSvc(repo repositories.PhotoRepository) *PhotoSvc {
	return &PhotoSvc{
		repo: repo,
	}
}

func (p *PhotoSvc) GetPhotos() *views.Response {
	photos, err := p.repo.GetPhotos()
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFoundResponse(err)
		}
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessGetResponse(parseModelToGetPhotos(photos), "Success get all photos")
}

func parseModelToGetPhotos(mod *[]models.Photo) *[]views.GetPhotos {
	var p []views.GetPhotos
	for _, v := range *mod {
		p = append(p, views.GetPhotos{
			Id:       *v.Id,
			Title:    v.Title,
			Caption:  v.Caption,
			PhotoUrl: v.PhotoUrl,
			UserId:   v.UserId,
		})
	}
	return &p
}

func (p *PhotoSvc) UploadPhoto(req *params.PhotoUploadRequest) *views.Response {
	photo := parseRequestToModelUploadPhoto(req)

	id := 1
	if photos, err := p.repo.GetPhotos(); err == nil {
		id = len(*photos) + 1
	}
	photo.Id = &id

	if err := p.repo.UploadPhoto(photo); err != nil {
		return views.InternalServerErrorResponse(err)
	}
	return views.SuccessCreateResponse(photo, "Success upload photo")
}

func parseRequestToModelUploadPhoto(req *params.PhotoUploadRequest) *models.Photo {
	return &models.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
		UserId:   req.UserId,
	}
}

func (p *PhotoSvc) UpdatePhoto(id int, req *params.PhotoUpdateRequest) *views.Response {
	photo := parseRequestToModelUpdatePhoto(req)
	photo.Id = &id

	if err := p.repo.UpdatePhoto(photo); err != nil {
		return views.InternalServerErrorResponse(err)
	}
	return views.SuccessUpdateResponse(photo, "Success update photo")
}

func parseRequestToModelUpdatePhoto(req *params.PhotoUpdateRequest) *models.Photo {
	return &models.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
	}
}

func (p *PhotoSvc) DeletePhoto(id string) *views.Response {
	if err := p.repo.DeletePhoto(id); err != nil {
		return views.InternalServerErrorResponse(err)
	}
	return views.SuccessDeleteResponse("Your photo has been successfuly deleted")
}
