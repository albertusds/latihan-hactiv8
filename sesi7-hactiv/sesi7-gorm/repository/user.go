package repository

import (
	"fmt"
	"sesi7-gorm/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(*models.User) error
	GetAllUsers() (*[]models.User, error)
	GetUserById(id uint) (*models.User, error)
	DeleteUserById(id uint) error
	UpdateUserById(id uint, email string) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(request *models.User) error {
	err := r.db.Create(request).Error
	return err
}

func (r *userRepo) GetAllUsers() (*[]models.User, error) {
	var users []models.User

	err := r.db.Find(&users).Error
	return &users, err
}

func (r *userRepo) GetUserById(id uint) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, "id=?", id).Error
	return &user, err
}

func (r *userRepo) DeleteUserById(id uint) error {
	var user models.User

	err := r.db.Delete(&user, "id=?", id).Error
	return err
}

func (r *userRepo) UpdateUserById(id uint, email string) error {
	var user models.User

	err := r.db.Model(&user).Where("id=?", id).Updates(models.User{Email: email}).Error
	if err != nil {
		fmt.Println("Error updating data.")
		return err
	}

	fmt.Println("Success update data with email :", email)
	return err
}
