package service

import (
	"context"
	"gin-mall/internal/model"
	"gin-mall/internal/params"
	"gin-mall/internal/repository"
	"github.com/pkg/errors"
)

type CategoryService interface {
	Edit(ctx context.Context, param *params.CategoryEditParam) error
	GetList(ctx context.Context, param *params.CategoryListParam) map[string]any
}

type categoryService struct{ *service }

func GetCategoryService() CategoryService {
	return categorySvc
}

func (s *categoryService) Edit(ctx context.Context, param *params.CategoryEditParam) error {
	categoryRepo := repository.GetCategoryRepo()
	if *param.Id > 0 {
		category, err := categoryRepo.GetById(ctx, *param.Id)
		if category == nil {
			return errors.Wrap(err, "未获取到分类信息")
		}
	}
	cate := &model.Category{
		Name:     param.Name,
		ParentId: *(param.ParentId),
		IsParent: *(param.IsParent),
		Sort:     *(param.Sort),
		ID:       *(param.Id),
	}
	return categoryRepo.SaveData(ctx, cate)
}

func (s *categoryService) GetList(ctx context.Context, param *params.CategoryListParam) map[string]any {
	categoryRepo := repository.GetCategoryRepo()
	condition := map[string]any{}
	query := categoryRepo.GetQuery(ctx, condition)
	return s.GetPage(query, param.Page, param.PageSize)
}
