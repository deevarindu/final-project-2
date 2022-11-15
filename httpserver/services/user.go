package services

import (
	"database/sql"
	"strings"

	"github.com/deevarindu/final-project-2/httpserver/controllers/params"
	"github.com/deevarindu/final-project-2/httpserver/repositories"
	"github.com/deevarindu/final-project-2/httpserver/repositories/models"
	"github.com/deevarindu/final-project-2/httpserver/views"
	"golang.org/x/crypto/bcrypt"
)

type UserSvc struct {
	repo repositories.UserRepository
}

func NewUserSvc(repo repositories.UserRepository) *UserSvc {
	return &UserSvc{
		repo: repo,
	}
}

func (u *UserSvc) FindUserByEmail(email string) *models.User {
	users, _ := u.repo.GetUsers()
	for _, user := range *users {
		if strings.EqualFold(user.Email, email) {
			return &user
		}
	}
	return nil
}

func (u *UserSvc) GetUsers() *views.Response {
	users, err := u.repo.GetUsers()
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFoundResponse(err)
		}
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessGetResponse(parseModelToGetUsers(users), "Success get all users")
}

func parseModelToGetUsers(mod *[]models.User) *[]views.GetUsers {
	var u []views.GetUsers
	for _, v := range *mod {
		u = append(u, views.GetUsers{
			Id:       *v.Id,
			Username: v.Username,
			Age:      v.Age,
		})
	}
	return &u
}

func (u *UserSvc) GetUser(id string) *views.Response {
	user, err := u.repo.GetUser(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFoundResponse(err)
		}
		return views.InternalServerErrorResponse(err)
	}
	return views.SuccessGetResponse(parseModelToGetUser(user), "Success get user")
}

func parseModelToGetUser(mod *models.User) *views.GetUser {
	return &views.GetUser{
		Id:       *mod.Id,
		Username: mod.Username,
		Email:    mod.Email,
		Password: mod.Password,
		Age:      mod.Age,
	}
}

func (u *UserSvc) CreateUser(req *params.UserCreateRequest) *views.Response {
	user := parseRequestToModelCreateUser(req)
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return views.BadRequestResponse(err)
	}
	user.Password = string(hash)

	id := 1
	if users, err := u.repo.GetUsers(); err == nil {
		id = len(*users) + 1
	}
	user.Id = &id

	err = u.repo.CreateUser(user)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return views.DataConflictResponse(err)
		}
		return views.InternalServerErrorResponse(err)
	}
	return views.SuccessCreateResponse(user, "Success create user")
}

func parseRequestToModelCreateUser(req *params.UserCreateRequest) *models.User {
	return &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
	}
}

func (u *UserSvc) UpdateUser(id int, req *params.UserUpdateRequest) *views.Response {
	user := parseRequestToModelUpdateUser(req)
	user.Id = &id
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return views.BadRequestResponse(err)
	}
	user.Password = string(hash)

	err = u.repo.UpdateUser(user)
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFoundResponse(err)
		}
		return views.InternalServerErrorResponse(err)
	}
	return views.SuccessUpdateResponse(user, "Success update user")
}

func parseRequestToModelUpdateUser(req *params.UserUpdateRequest) *models.User {
	return &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
	}
}

func (u *UserSvc) DeleteUser(id string) *views.Response {
	err := u.repo.DeleteUser(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFoundResponse(err)
		}
		return views.InternalServerErrorResponse(err)
	}
	return views.SuccessDeleteResponse("Your account has been successfuly deleted")
}
