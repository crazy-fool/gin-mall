package service

import (
	"context"
	"gin-mall/internal/model"
	"gin-mall/internal/params"
	"gin-mall/internal/repository"
	"gin-mall/pkg/auth"
	"gin-mall/pkg/constant"
	"gin-mall/pkg/helper/sid"
	"gin-mall/pkg/log"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	Register(ctx context.Context, req *params.RegisterRequest) error
	Login(ctx context.Context, req *params.LoginRequest) (string, error)
	GetProfile(ctx context.Context, userId string) (*model.Customer, error)
	UpdateProfile(ctx context.Context, userId string, req *params.UpdateProfileRequest) error
}

type userService struct{}

func GetUserService() UserService {
	return userSvc
}

func (s *userService) Register(ctx context.Context, req *params.RegisterRequest) error {
	// 检查用户名是否已存在
	if user, err := repository.GetUserRepo().GetByAccount(ctx, req.Account); err == nil && user != nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}
	// Generate user ID
	code, err := sid.GetSid().GenString()
	if err != nil {
		return errors.Wrap(err, "failed to generate user ID")
	}
	// Create a user
	user := &model.Customer{
		Code:     code,
		Name:     req.Name,
		Password: string(hashedPassword),
		Mobile:   req.Mobile,
		Account:  req.Account,
		//LastLoginAt: nil,
	}
	if err = repository.GetUserRepo().Create(ctx, user); err != nil {
		log.GetLog().Info("操作失败:" + err.Error())
		return errors.Wrap(err, "failed to create user")
	}

	return nil
}

func (s *userService) Login(ctx context.Context, req *params.LoginRequest) (string, error) {
	user, err := repository.GetUserRepo().GetByAccount(ctx, req.Account)
	if err != nil || user == nil {
		return "", errors.New("failed to get user by account")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.Wrap(err, "failed to hash password")
	}
	token, err := auth.GetJwt().GenToken(user.Code, time.Now().Add(constant.JwtTtl))
	if err != nil || token == "" {
		return "", errors.Wrap(err, "failed to generate JWT token")
	}
	return token, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*model.Customer, error) {
	user, err := repository.GetUserRepo().GetByID(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by ID")
	}

	return user, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *params.UpdateProfileRequest) error {
	return nil
}
