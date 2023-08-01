package service

import (
	"gin-mall/pkg/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

func (s *service) GetPage(query *gorm.DB, page int, pageSize int) map[string]any {
	list := make([]map[string]any, 0)
	var total int64
	if err := query.Count(&total).Error; err != nil {
		log.GetLog().Info("[分页搜索]查询数据条目出错：", zap.Error(err))
	}
	if total > 0 {
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		if err := query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&list).Error; err != nil {
			log.GetLog().Info("[分页搜索]查询数据出错：", zap.Error(err))
		}
	}

	return map[string]any{
		"total": total,
		"list":  list,
	}
}
