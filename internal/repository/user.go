package repository

import (
	"context"
	"gin-mall/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.Customer) error
	Update(ctx context.Context, user *model.Customer) error
	GetByID(ctx context.Context, id string) (*model.Customer, error)
	GetByAccount(ctx context.Context, account string) (*model.Customer, error)
}

type userRepository struct {
	*Repository
}

func GetUserRepo() UserRepository {
	return userRepo
}

func (r *userRepository) Create(ctx context.Context, user *model.Customer) error {
	if err := r.db.Create(user).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.Customer) error {
	if err := r.db.Save(user).Error; err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func (r *userRepository) GetByID(ctx context.Context, userId string) (*model.Customer, error) {
	var user model.Customer
	if err := r.db.Where("user_id = ?", userId).First(&user).Error; err != nil {

		return nil, errors.Wrap(err, "failed to get user by ID")
	}

	return &user, nil
}

func (r *userRepository) GetByAccount(ctx context.Context, account string) (*model.Customer, error) {
	var customer model.Customer
	if err := r.db.Where("account = ?", account).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get user by username")
	}

	return &customer, nil
}
