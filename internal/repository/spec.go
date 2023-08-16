package repository

import (
	"context"
	"gin-mall/internal/gen/model"
	"gin-mall/internal/gen/query"
	"gin-mall/internal/params"
)

type SpecRepository interface {
	GetSpecById(ctx context.Context, id int32) (*model.Spec, error)
	GetGroupById(ctx context.Context, id int32) (*model.SpecGroup, error)
	SaveSpecData(ctx context.Context, param *params.SpecEditParam) (int32, error)
	SaveGroupData(ctx context.Context, param *params.GroupEditParam) (int32, error)
	GetSpecPage(ctx context.Context, param *params.SpecListParam) (result []*model.Spec, count int64, err error)
	GetGroupPage(ctx context.Context, param *params.GroupListParam) (result []*model.SpecGroup, count int64, err error)
	GetSpecAll(ctx context.Context, param *params.SpecListParam) (result []*model.Spec, err error)
	GetGroupAll(ctx context.Context, param *params.GroupListParam) (result []*model.SpecGroup, err error)
}

type specRepository struct {
	*Repository
}

func (s *specRepository) GetSpecById(ctx context.Context, id int32) (*model.Spec, error) {
	t := query.Use(s.db).Spec
	return t.WithContext(ctx).Where(t.ID.Eq(id)).Take()
}

func (s *specRepository) GetGroupById(ctx context.Context, id int32) (*model.SpecGroup, error) {
	t := query.Use(s.db).SpecGroup
	return t.WithContext(ctx).Where(t.ID.Eq(id)).Take()
}

func (s *specRepository) SaveSpecData(ctx context.Context, param *params.SpecEditParam) (int32, error) {
	spec := new(model.Spec)
	if param.Id > 0 {
		c, err := s.GetSpecById(ctx, param.Id)
		if err != nil {
			return 0, err
		}
		spec = c
	}
	t := query.Use(s.db).Spec
	if param.Name != "" {
		spec.Name = &param.Name
	}

	if param.CategoryId != 0 {
		spec.CategoryID = &param.CategoryId
	}

	if param.GroupId != 0 {
		spec.GroupID = &param.GroupId
	}
	if param.IsGenerate != nil {
		isGenerate := *param.IsGenerate == 1
		spec.IsGenerate = &isGenerate
	}

	if param.Searching != nil {
		searching := *param.Searching == 1
		spec.Searching = &searching
	}

	err := t.WithContext(ctx).Save(spec)
	s.recordError("商品属性保存", err)
	return spec.ID, err
}

func (s *specRepository) SaveGroupData(ctx context.Context, param *params.GroupEditParam) (int32, error) {
	spec := new(model.SpecGroup)
	if param.Id > 0 {
		c, err := s.GetGroupById(ctx, param.Id)
		if err != nil {
			return 0, err
		}
		spec = c
	}
	t := query.Use(s.db).SpecGroup
	if param.Name != "" {
		spec.Name = &param.Name
	}

	if param.CategoryId != 0 {
		spec.CategoryID = param.CategoryId
	}
	err := t.WithContext(ctx).Save(spec)
	s.recordError("商品属性分组保存", err)
	return spec.ID, err
}

func (s *specRepository) GetSpecPage(ctx context.Context, param *params.SpecListParam) (result []*model.Spec, count int64, err error) {
	return s.getSpecQuery(ctx, param).FindByPage(param.OffsetLimit())
}

func (s *specRepository) GetGroupPage(ctx context.Context, param *params.GroupListParam) (result []*model.SpecGroup, count int64, err error) {
	return s.getGroupQuery(ctx, param).FindByPage(param.OffsetLimit())
}

func (s *specRepository) GetSpecAll(ctx context.Context, param *params.SpecListParam) (result []*model.Spec, err error) {
	return s.getSpecQuery(ctx, param).Find()
}

func (s *specRepository) GetGroupAll(ctx context.Context, param *params.GroupListParam) (result []*model.SpecGroup, err error) {
	return s.getGroupQuery(ctx, param).Find()
}

func (s *specRepository) getGroupQuery(ctx context.Context, param *params.GroupListParam) query.ISpecGroupDo {
	t := query.Use(s.db).SpecGroup
	q := t.WithContext(ctx)
	return q.Order(t.ID)
}

func (s *specRepository) getSpecQuery(ctx context.Context, param *params.SpecListParam) query.ISpecDo {
	t := query.Use(s.db).Spec
	q := t.WithContext(ctx)
	return q.Order(t.ID)
}

func GetSpecRepo() SpecRepository {
	return specRepo
}
