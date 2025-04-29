package users

import (
	"errors"
	"github.com/Imnarka/user-service/internal/logger"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	Update(user *User) error
	Delete(id uint) error
	List() ([]User, error)
}

type UserRepositoryImpl struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewUserRepository(db *gorm.DB, logger *logger.Logger) UserRepository {
	return &UserRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepositoryImpl) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) GetByID(id uint) (*User, error) {
	r.logger.WithField("id", id).Debug("getting user by id")
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		r.logger.WithError(err).Error("could not get user by id")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Update(user *User) error {
	r.logger.WithField("id", user.ID).Debug("updating user")
	return r.db.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	r.logger.WithField("id", id).Debug("deleting user")
	result := r.db.Delete(&User{}, id)
	if result.RowsAffected == 0 {
		return ErrUserNotFound
	}
	return result.Error
}

func (r *UserRepositoryImpl) List() ([]User, error) {
	r.logger.WithField("db", r.db).Debug("get all users")
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
