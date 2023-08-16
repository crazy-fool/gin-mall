package service

import (
	"context"
	"gin-mall/internal/gen/model"
	"gin-mall/internal/params"
)

type SpuService interface {
	EditSpu(ctx context.Context, param *params.SpuEditParam) (model.Spu, error)
}
type spuService struct{ *service }

func GetSpuService() SpuService {
	return spuSvc
}

func (s *spuService) EditSpu(ctx context.Context, param *params.SpuEditParam) (model.Spu, error) {
	panic("1")
}
