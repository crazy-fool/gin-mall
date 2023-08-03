package service

import (
	"context"
	"gin-mall/internal/gen/model"
	"gin-mall/internal/params"
	"gin-mall/internal/repository"
	"gin-mall/pkg/log"
	"go.uber.org/zap"
)

type CategoryService interface {
	Edit(ctx context.Context, param *params.CategoryEditParam) error
	GetPageList(ctx context.Context, param *params.CategoryListParam) map[string]any
	GetList(ctx context.Context, param *params.CategoryListParam) []*model.Category
}

type categoryService struct{ *service }

func GetCategoryService() CategoryService {
	return categorySvc
}

func (s *categoryService) Edit(ctx context.Context, param *params.CategoryEditParam) error {
	categoryRepo := repository.GetCategoryRepo()
	_, err := categoryRepo.SaveData(ctx, param)
	return err
}

func (s *categoryService) GetPageList(ctx context.Context, param *params.CategoryListParam) map[string]any {
	categoryRepo := repository.GetCategoryRepo()
	result, total, err := categoryRepo.Page(ctx, param)
	if err != nil {
		log.GetLog().Info("[数据库查询异常]", zap.Error(err))
	}
	return s.GetPage(total, result)
}

func (s *categoryService) GetList(ctx context.Context, param *params.CategoryListParam) []*model.Category {
	categoryRepo := repository.GetCategoryRepo()
	if param.ParentId == nil {
		param.ParentId = new(int32)
	}
	list, err := categoryRepo.GetAll(ctx, param)
	if err != nil {
		log.GetLog().Info("[数据库查询异常]", zap.Error(err))
	}
	return list
}
