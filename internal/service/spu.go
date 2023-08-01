package service

import (
	"context"
	"gin-mall/internal/model"
	"gin-mall/internal/params"
)

type SpuService interface {
	EditSpu(ctx context.Context, req *params.EditSpuParam) (model.Spu, error)
}
type spuService struct{}

func GetSpuService() UserService {
	return userSvc
}

func (s *spuService) EditSpu(ctx context.Context, req *params.EditSpuParam) {

}
