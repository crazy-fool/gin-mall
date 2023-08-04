package repository

import (
	"context"
	"gin-mall/internal/gen/model"
	"github.com/pkg/errors"
)

type SpuRepository interface {
	Create(ctx context.Context, user *model.Spu) error
	Update(ctx context.Context, user *model.Spu) error
	GetByID(ctx context.Context, id string) (*model.Spu, error)
}

type spuRepository struct {
	*Repository
}

func GetSpuRepo() SpuRepository {
	return spuRepo
}

func (r *spuRepository) Create(ctx context.Context, user *model.Spu) error {
	if err := r.db.Create(user).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (r *spuRepository) Update(ctx context.Context, user *model.Spu) error {
	if err := r.db.Save(user).Error; err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func (r *spuRepository) GetByID(ctx context.Context, userId string) (*model.Spu, error) {
	var spu model.Spu
	if err := r.db.Where("user_id = ?", userId).First(&spu).Error; err != nil {

		return nil, errors.Wrap(err, "failed to get user by ID")
	}

	return &spu, nil
}
