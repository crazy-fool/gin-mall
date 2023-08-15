package service

import (
	"context"
	"gin-mall/internal/gen/model"
)

type SpuService interface {
	EditSpu(ctx context.Context) (model.Spu, error)
}
type spuService struct{}

func GetSpuService() UserService {
	return userSvc
}

func (s *spuService) EditSpu(ctx context.Context) {

}
