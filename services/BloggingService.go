package services

import "github.com/Phogheus/GoBlog/models"

type bloggingService struct {
	BlogRepository models.IBlogRepository
}

func NewBloggingService(blogRepo models.IBlogRepository) models.IBloggingService {
	return &bloggingService{
		BlogRepository: blogRepo,
	}
}

func (repo *bloggingService) CreateNewBlogPost(request models.CreateBlogPostRequest) models.BlogPostCreatedResponse {
	var response models.BlogPostCreatedResponse
	return response
}

func (repo *bloggingService) GetBlogPosts(request models.GetBlogPostsRequest) models.BlogPostRetrievedResponse {
	var response models.BlogPostRetrievedResponse
	return response
}

func (repo *bloggingService) UpdateBlogPost(request models.UpdateBlogPostRequest) models.BlogPostUpdatedResponse {
	var response models.BlogPostUpdatedResponse
	return response
}

func (repo *bloggingService) DeleteBlogPosts(request models.DeleteBlogPostsRequest) models.BlogPostDeletedResponse {
	var response models.BlogPostDeletedResponse
	return response
}
