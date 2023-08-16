package service

import (
	"context"
	"gin-mall/internal/params"
	"gin-mall/internal/repository"
	"gin-mall/pkg/log"
	"go.uber.org/zap"
)

type SpecService interface {
	EditSpec(ctx context.Context, param *params.SpecEditParam) error
	EditGroup(ctx context.Context, param *params.GroupEditParam) error
	GetGroupList(ctx context.Context, param *params.GroupListParam) map[string]any
	GetSpecList(ctx context.Context, param *params.SpecListParam) map[string]any
}

type specService struct {
	*service
}

func (s specService) EditSpec(ctx context.Context, param *params.SpecEditParam) error {
	specRepo := repository.GetSpecRepo()
	_, err := specRepo.SaveSpecData(ctx, param)
	return err
}

func (s specService) EditGroup(ctx context.Context, param *params.GroupEditParam) error {
	specRepo := repository.GetSpecRepo()
	_, err := specRepo.SaveGroupData(ctx, param)
	return err
}

func (s specService) GetGroupList(ctx context.Context, param *params.GroupListParam) map[string]any {
	specRepo := repository.GetSpecRepo()
	result, total, err := specRepo.GetGroupPage(ctx, param)
	if err != nil {
		log.GetLog().Info("[数据库查询异常]", zap.Error(err))
	}
	return s.GetPage(total, result)
}

func (s specService) GetSpecList(ctx context.Context, param *params.SpecListParam) map[string]any {
	specRepo := repository.GetSpecRepo()
	result, total, err := specRepo.GetSpecPage(ctx, param)
	if err != nil {
		log.GetLog().Info("[数据库查询异常]", zap.Error(err))
	}
	return s.GetPage(total, result)
}

func GetSpecService() SpecService {
	return specSvc
}
