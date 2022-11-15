package services

import (
	"database/sql"

	"github.com/deevarindu/final-project-2/httpserver/controllers/params"
	"github.com/deevarindu/final-project-2/httpserver/repositories"
	"github.com/deevarindu/final-project-2/httpserver/repositories/models"
	"github.com/deevarindu/final-project-2/httpserver/views"
)

type CommentSvc struct {
	repo repositories.CommentRepository
}

func NewCommentSvc(repo repositories.CommentRepository) *CommentSvc {
	return &CommentSvc{
		repo: repo,
	}
}

func (c *CommentSvc) GetComments() *views.Response {
	comments, err := c.repo.GetComments()
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFoundResponse(err)
		}
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessGetResponse(parseModelToGetComments(comments), "Success get all comments")
}

func parseModelToGetComments(mod *[]models.Comment) *[]views.GetComments {
	var c []views.GetComments
	for _, v := range *mod {
		c = append(c, views.GetComments{
			Id:      *v.Id,
			Message: v.Message,
			UserId:  v.UserId,
			PhotoId: v.PhotoId,
		})
	}
	return &c
}

func (c *CommentSvc) CreateComment(req *params.CommentCreateRequest) *views.Response {
	comment := parseRequestToModelCreateComment(req)

	id := 1
	if comments, err := c.repo.GetComments(); err == nil {
		id = len(*comments) + 1
	}
	comment.Id = &id

	if err := c.repo.CreateComment(comment); err != nil {
		return views.InternalServerErrorResponse(err)
	}
	return views.SuccessCreateResponse(comment, "Success create comment")
}

func parseRequestToModelCreateComment(req *params.CommentCreateRequest) *models.Comment {
	return &models.Comment{
		Message: req.Message,
		PhotoId: req.PhotoId,
	}
}

func (c *CommentSvc) UpdateComment(id int, req *params.CommentUpdateRequest) *views.Response {
	comment := parseRequestToModelUpdateComment(req)
	comment.Id = &id

	if err := c.repo.UpdateComment(comment); err != nil {
		return views.InternalServerErrorResponse(err)
	}
	return views.SuccessUpdateResponse(comment, "Success update comment")
}

func parseRequestToModelUpdateComment(req *params.CommentUpdateRequest) *models.Comment {
	return &models.Comment{
		Message: req.Message,
	}
}

func (c *CommentSvc) DeleteComment(id string) *views.Response {
	if err := c.repo.DeleteComment(id); err != nil {
		return views.InternalServerErrorResponse(err)
	}
	return views.SuccessDeleteResponse("Success delete comment")
}
