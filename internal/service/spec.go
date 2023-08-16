package service

import (
	"context"
	"gin-mall/internal/params"
)

type SpecService interface {
	EditSpec(ctx context.Context, param *params.SpecEditParam)
	EditGroup(ctx context.Context, param *params.GroupEditParam)
	GetGroupList(ctx context.Context, param *params.GroupListParam)
	GetSpecList(ctx context.Context, param *params.SpecListParam)
}
