package services

import (
	"context"
	"errors"
	"net/http"
	// "strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mrestuf/kanban-board/common"
	"github.com/mrestuf/kanban-board/config"
	"github.com/mrestuf/kanban-board/httpserver/controllers/params"
	"github.com/mrestuf/kanban-board/httpserver/controllers/views"
	"github.com/mrestuf/kanban-board/httpserver/repositories"
	"github.com/mrestuf/kanban-board/httpserver/repositories/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userSvc struct {
	repo repositories.UserRepo
}

func NewUserSvc(repo repositories.UserRepo) UserSvc {
	return &userSvc{
		repo: repo,
	}
}

func (s *userSvc) Register(ctx context.Context, user *params.Register) *views.Response {
	_, err := s.repo.FindUserByEmail(ctx, user.Email)
	if err == nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_EMAIL_ALREADY_USED, errors.New("email already used"))
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	//req
	input := models.User{
		FullName: user.FullName,
		Email: user.Email,
		Password: string(hashedPassword),
	}
	err = s.repo.CreateUser(ctx, &input)

	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.Register{
		Id:           input.Id,
		FullName:     input.FullName,
		Email:        input.Email,
		CreatedAt:    input.CreatedAt,
	})
}

func (s *userSvc) Login(ctx context.Context, user *params.Login) *views.Response {
	model, err := s.repo.FindUserByEmail(ctx, user.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_INVALID_CREDENTIALS, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(user.Password))
	if err != nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_INVALID_CREDENTIALS, err)
	}

	claims := &common.CustomClaims{
		Id: model.Id,
	}
	claims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(config.GetJwtExpiredTime())).Unix()
	claims.Subject = model.Email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.GetJwtSignature())
	
	return views.SuccessResponse(http.StatusOK, views.M_OK, views.Login{
		Token: ss,
	})
}