package repo

import (
	"context"
	"github.com/DalvinCodes/digital-commerce/users/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
}

type UserRepo struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	return r.Db.Debug().Create(&user).Error
}
