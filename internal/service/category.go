package service

import (
	"context"
	"gin-mall/internal/params"
)

type CategoryService interface {
	Edit(ctx context.Context, param *params.EditCategoryParam) error
}

type categoryService struct{}

func GetCategoryService() UserService {
	return userSvc
}

func (s *categoryService) Edit(ctx context.Context, param *params.EditCategoryParam) {
	if param.Id > 0 {

	}
}
