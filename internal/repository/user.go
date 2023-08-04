package repository

import (
	"context"
	"gin-mall/internal/gen/model"
	"gin-mall/internal/gen/query"
	"gin-mall/internal/params"
	"gin-mall/pkg/helper/sid"
)

type UserRepository interface {
	SaveData(ctx context.Context, param *params.RegisterParam) (int32, error)
	GetByUid(ctx context.Context, uid string) (*model.Customer, error)
	GetByID(ctx context.Context, id int32) (*model.Customer, error)
	GetByAccount(ctx context.Context, account string) (*model.Customer, error)
}

type userRepository struct {
	*Repository
}

func GetUserRepo() UserRepository {
	return userRepo
}

func (r *userRepository) GetByUid(ctx context.Context, uid string) (*model.Customer, error) {
	t := query.Use(r.db).Customer
	return t.WithContext(ctx).Where(t.Code.Eq(uid)).Take()
}

func (r *userRepository) GetByID(ctx context.Context, id int32) (*model.Customer, error) {
	t := query.Use(r.db).Customer
	return t.WithContext(ctx).Where(t.ID.Eq(id)).Take()
}

func (r *userRepository) GetByAccount(ctx context.Context, account string) (*model.Customer, error) {
	t := query.Use(r.db).Customer
	return t.WithContext(ctx).Where(t.Account.Eq(account)).Take()
}

func (r *userRepository) SaveData(ctx context.Context, param *params.RegisterParam) (int32, error) {
	customer := new(model.Customer)
	if param.Id > 0 {
		c, err := r.GetByID(ctx, param.Id)
		if err != nil {
			return 0, err
		}
		customer = c
	}
	t := query.Use(r.db).Customer
	if param.Name != "" {
		customer.Name = param.Name
	}

	if param.Password != "" {
		customer.Password = param.Password
	}

	if param.Account != "" {
		customer.Account = param.Account
	}

	if param.Mobile != "" {
		customer.Mobile = param.Mobile
	}

	if customer.Code == "" {
		// Generate user ID
		code, err := sid.GetSid().GenString()
		if err != nil {
			return 0, err
		}
		customer.Code = code
	}

	err := t.WithContext(ctx).Save(customer)
	r.recordError("customer保存数据", err)
	return customer.ID, err
}
