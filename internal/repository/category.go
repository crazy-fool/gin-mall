package repository

import (
	"context"
	"gin-mall/internal/gen/model"
	"gin-mall/internal/gen/query"
	"gin-mall/internal/params"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetById(ctx context.Context, id int32) (*model.Category, error)
	SaveData(ctx context.Context, param *params.CategoryEditParam) (int32, error)
	GetQuery(ctx context.Context, param *params.CategoryListParam) *gorm.DB
}

type categoryRepository struct {
	*Repository
}

func GetCategoryRepo() CategoryRepository {
	return categoryRepo
}

func (r *categoryRepository) GetById(ctx context.Context, id int32) (*model.Category, error) {
	var category model.Category
	if err := r.db.Find(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get user by username")
	}

	return &category, nil
}

func (r *categoryRepository) SaveData(ctx context.Context, param *params.CategoryEditParam) (int32, error) {
	cate := &model.Category{
		ID:   param.Id,
		Name: param.Name,
		Sort: *param.Sort,
	}
	m := query.Use(r.db).Category
	if param.ParentId != nil {
		cate.ParentID = *param.ParentId
	}
	if param.IsParent != nil {
		cate.IsParent = *param.IsParent == 1
	}

	err := m.WithContext(ctx).Save(cate)
	return cate.ID, err
}

func (r *categoryRepository) GetQuery(ctx context.Context, param *params.CategoryListParam) *gorm.DB {
	query := r.db.Model(model.Category{})
	if param.ParentId != nil {
		query.Where("parent_id = ?", *param.ParentId)
	}
	return query.Order("sort desc")
}
