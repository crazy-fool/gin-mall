package repository

import (
	"context"
	"fmt"
	"gin-mall/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetById(ctx context.Context, id uint) (*model.Category, error)
	SaveData(ctx context.Context, category *model.Category) error
	GetQuery(ctx context.Context, condition map[string]any) *gorm.DB
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

func (r *categoryRepository) SaveData(ctx context.Context, category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) GetQuery(ctx context.Context, condition map[string]any) *gorm.DB {
	query := r.db.Model(model.Category{})
	for k, v := range condition {

		if k == "name" {
			name := fmt.Sprintf("%v", v)
			if name != "" {
				query.Where("name like ?", "%"+name+"%")
			}
			continue
		} else if k == "parent_id" {
			query.Where("parent_id = ?", v)
		}
	}
	return query.Order("sort desc")
}
