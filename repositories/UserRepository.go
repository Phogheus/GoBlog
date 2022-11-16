package repositories

import (
	"database/sql"

	"github.com/Phogheus/GoBlog/models"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) models.IUserRepository {
	return &userRepository{
		DB: db,
	}
}

func (repo *userRepository) CreateNewUser(request models.CreateBlogUserRequest) (bool, error) {
	return false, nil
}

func (repo *userRepository) GetUserById(id int) (*models.BlogUser, error) {
	var user models.BlogUser

	return &user, nil
}

func (repo *userRepository) GetUserByUsername(username string) (*models.BlogUser, error) {
	var user models.BlogUser

	return &user, nil
}

func (repo *userRepository) GetUserByEmail(email string) (*models.BlogUser, error) {
	var user models.BlogUser

	return &user, nil
}

func (repo *userRepository) UpdateUser(request models.UpdateBlogUserRequest) (bool, error) {
	return false, nil
}

func (repo *userRepository) DeleteUserById(id int) (bool, error) {
	return false, nil
}

func (repo *userRepository) LockUserById(id int) (bool, error) {
	return false, nil
}

func (repo *userRepository) UnlockUserById(id int) (bool, error) {
	return false, nil
}

func (repo *userRepository) DisableUserById(id int) (bool, error) {
	return false, nil
}

func (repo *userRepository) EnableUserById(id int) (bool, error) {
	return false, nil
}
