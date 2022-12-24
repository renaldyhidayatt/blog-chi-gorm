package repository

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
	"errors"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CheckUser(loginRequest request.LoginRequest) (entity.User, error) {
	var user entity.User

	r.db.Where("username=?", loginRequest.Username).Take(&user)

	if user.Username == "" {
		return entity.User{}, errors.New("username tidak ditemukan")
	}

	if !user.Status {
		return entity.User{}, errors.New("account tidak aktif")
	}

	return user, nil

}

func (r *authRepository) ForgetPassword(forgetPassword request.ForgetPasswordRequest, userId int64) (entity.User, error) {
	var user entity.User

	r.db.Where("id_user=?", userId).Take(&user)

	if user.ID == 0 {
		return entity.User{}, errors.New("id user tidak ditemukan")
	}

	row := r.db.Model(&user).Where("id_user=?", userId).Update("password", forgetPassword.NewPassword)

	if row.Error != nil {
		return entity.User{}, row.Error
	}

	return user, nil

}
