package service

import "gin-mall/pkg/log"

func init() {
	log.GetLog().Info("===========初始化UserService=================")
	userSvc = &userService{}
}

var userSvc *userService
