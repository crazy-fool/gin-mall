package service

import (
	"context"
	"gin-mall/internal/params"
	"gin-mall/internal/repository"
	"gin-mall/pkg/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type CategoryService interface {
	Edit(ctx context.Context, param *params.CategoryEditParam) error
	GetPageList(ctx context.Context, param *params.CategoryListParam) map[string]any
	GetList(ctx context.Context, param *params.CategoryListParam) []map[string]any
}

type categoryService struct{ *service }

func GetCategoryService() CategoryService {
	return categorySvc
}

func (s *categoryService) Edit(ctx context.Context, param *params.CategoryEditParam) error {
	categoryRepo := repository.GetCategoryRepo()
	if param.Id > 0 {
		category, err := categoryRepo.GetById(ctx, param.Id)
		if category == nil {
			return errors.Wrap(err, "未获取到分类信息")
		}
	}
	_, err := categoryRepo.SaveData(ctx, param)
	return err
}

func (s *categoryService) GetPageList(ctx context.Context, param *params.CategoryListParam) map[string]any {
	categoryRepo := repository.GetCategoryRepo()
	query := categoryRepo.GetQuery(ctx, param)
	return s.GetPage(query, param.Page, param.PageSize)
}

func (s *categoryService) GetList(ctx context.Context, param *params.CategoryListParam) []map[string]any {
	categoryRepo := repository.GetCategoryRepo()
	if param.ParentId == nil {
		param.ParentId = new(uint)
	}
	query := categoryRepo.GetQuery(ctx, param)
	list := make([]map[string]any, 0)
	if err := query.Find(&list).Error; err != nil {
		log.GetLog().Info("查询分类数据失败", zap.Error(err))
	}
	return list
}
