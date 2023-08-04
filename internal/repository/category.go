package repository

import (
	"context"
	"gin-mall/internal/gen/model"
	"gin-mall/internal/gen/query"
	"gin-mall/internal/params"
)

type CategoryRepository interface {
	GetById(ctx context.Context, id int32) (*model.Category, error)
	SaveData(ctx context.Context, param *params.CategoryEditParam) (int32, error)
	Page(ctx context.Context, param *params.CategoryListParam) (result []*model.Category, count int64, err error)
	GetAll(ctx context.Context, param *params.CategoryListParam) (result []*model.Category, err error)
}

type categoryRepository struct {
	*Repository
}

func GetCategoryRepo() CategoryRepository {
	return categoryRepo
}

// GetById 获取单条数据
func (r *categoryRepository) GetById(ctx context.Context, id int32) (*model.Category, error) {
	t := query.Use(r.db).Category
	return t.WithContext(ctx).Where(t.ID.Eq(id)).Take()
}

// Page 分页数据
func (r *categoryRepository) Page(ctx context.Context, param *params.CategoryListParam) (result []*model.Category, count int64, err error) {
	return r.getQuery(ctx, param).FindByPage(param.OffsetLimit())
}

// GetAll 获取所有数据不分页
func (r *categoryRepository) GetAll(ctx context.Context, param *params.CategoryListParam) (result []*model.Category, err error) {
	q := r.getQuery(ctx, param)
	return q.Find()
}

// SaveData 指定字段保存数据，ID 0 新增，ID大于0更新
func (r *categoryRepository) SaveData(ctx context.Context, param *params.CategoryEditParam) (int32, error) {
	cate := new(model.Category)
	if param.Id > 0 {
		c, err := r.GetById(ctx, param.Id)
		if err != nil {
			return 0, err
		}
		cate = c
	}
	t := query.Use(r.db).Category
	if param.IsParent != nil {
		cate.IsParent = *param.IsParent == 1
	}

	if param.Name != "" {
		cate.Name = param.Name
	}

	if param.Sort != nil {
		cate.Sort = *param.Sort
	}

	if param.ParentId != nil {
		cate.ParentID = *param.ParentId
	}

	err := t.WithContext(ctx).Save(cate)
	r.recordError("商品分类保存数据", err)
	return cate.ID, err
}

func (r *categoryRepository) getQuery(ctx context.Context, param *params.CategoryListParam) query.ICategoryDo {
	t := query.Use(r.db).Category
	q := t.WithContext(ctx)
	if param.ParentId != nil {
		q.Where(t.ParentID.Eq(*param.ParentId))
	}
	return q.Order(t.Sort)
}
