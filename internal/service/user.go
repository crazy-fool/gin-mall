package service

import (
	"context"
	"gin-mall/internal/params"
	"gin-mall/internal/repository"
	"gin-mall/pkg/auth"
	"gin-mall/pkg/constant"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	Register(ctx context.Context, req *params.RegisterParam) error
	Login(ctx context.Context, req *params.LoginParam) (string, error)
}

type userService struct{}

func GetUserService() UserService {
	return userSvc
}

func (s *userService) Register(ctx context.Context, param *params.RegisterParam) error {
	// 检查用户名是否已存在
	if user, err := repository.GetUserRepo().GetByAccount(ctx, param.Account); err == nil && user != nil {
		return errors.New("用户已存在")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*param.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "请重试")
	}
	*param.Password = string(hashedPassword)

	if _, err := repository.GetUserRepo().SaveData(ctx, param); err != nil {
		return errors.Wrap(err, "请重试！")
	}
	return nil
}

func (s *userService) Login(ctx context.Context, req *params.LoginParam) (string, error) {
	user, err := repository.GetUserRepo().GetByAccount(ctx, req.Account)
	if err != nil || user == nil {
		return "", errors.New("failed to get user by account")
	}
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.Wrap(err, "failed to hash password")
	}
	token, err := auth.GetJwt().GenToken(user.Code, time.Now().Add(constant.JwtTtl))
	if err != nil || token == "" {
		return "", errors.Wrap(err, "failed to generate JWT token")
	}
	return token, nil
}
