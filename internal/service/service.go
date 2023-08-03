package service

import (
	"gin-mall/pkg/log"
)

func init() {
	log.GetLog().Info("===========初始化UserService=================")
	service := new(service)
	userSvc = &userService{}
	categorySvc = &categoryService{service}
}

type service struct {
}

var userSvc *userService
var categorySvc *categoryService

func (s *service) GetPage(total int64, result any) map[string]any {

	return map[string]any{
		"total": total,
		"list":  result,
	}
}
