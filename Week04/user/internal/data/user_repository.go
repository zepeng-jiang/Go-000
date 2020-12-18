package data

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/zepeng-jiang/Go-000/Week04/user/internal/entity"
)

type UserRepository struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewUserRepository(db *gorm.DB, cache *redis.Client) *UserRepository {
	return &UserRepository{
		db:    db,
		cache: cache,
	}
}

func (d *UserRepository) SaveUser(user *entity.User) error {
	if err := d.db.Model(user).Create(user).Error; err != nil {
		return errors.Wrapf(err, "[data] save user failed! error: %s\n", err.Error())
	}
	return nil
}
