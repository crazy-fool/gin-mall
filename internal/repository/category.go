package repository

import (
	"context"
	"gin-mall/internal/model"
	"gin-mall/internal/params"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetById(ctx context.Context, id uint) (*model.Category, error)
	SaveData(ctx context.Context, param *params.CategoryEditParam) (uint, error)
	GetQuery(ctx context.Context, param *params.CategoryListParam) *gorm.DB
}

type categoryRepository struct {
	*Repository
}

func GetCategoryRepo() CategoryRepository {
	return categoryRepo
}

func (r *categoryRepository) GetById(ctx context.Context, id uint) (*model.Category, error) {
	var category model.Category
	if err := r.db.Find(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get user by username")
	}

	return &category, nil
}

func (r *categoryRepository) SaveData(ctx context.Context, param *params.CategoryEditParam) (uint, error) {
	cate := model.Category{
		Name:     param.Name,
		ParentId: param.ParentId,
		IsParent: param.IsParent,
		Sort:     param.Sort,
		ID:       param.Id,
	}
	return cate.ID, r.db.Save(&cate).Error
}

func (r *categoryRepository) GetQuery(ctx context.Context, param *params.CategoryListParam) *gorm.DB {
	query := r.db.Model(model.Category{})
	if param.ParentId != nil {
		query.Where("parent_id = ?", *param.ParentId)
	}
	return query.Order("sort desc")
}
