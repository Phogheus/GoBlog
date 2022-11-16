package services

import "github.com/Phogheus/GoBlog/models"

type userService struct {
	UserRepository models.IUserRepository
}

func NewUserService(userRepo models.IUserRepository) models.IUserService {
	return &userService{
		UserRepository: userRepo,
	}
}

func (svc *userService) CreateUser(request models.CreateBlogUserRequest) models.BlogUserCreatedResponse {
	var response models.BlogUserCreatedResponse

	// Create user

	return response
}

func (svc *userService) GetUser(request models.GetBlogUserRequest) models.BlogUserRetrievedResponse {
	var response models.BlogUserRetrievedResponse

	// Get user

	return response
}

func (svc *userService) UpdateUser(request models.UpdateBlogUserRequest) models.BlogUserUpdatedResponse {
	var response models.BlogUserUpdatedResponse

	// Update user

	return response
}

func (svc *userService) SetUserActiveStatus(request models.SetBlogUserActiveStatusRequest) models.BlogUserSetActiveStatusResponse {
	var response models.BlogUserSetActiveStatusResponse

	// Set Active Status for user

	return response
}
