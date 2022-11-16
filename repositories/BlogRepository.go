package repositories

import (
	"database/sql"

	"github.com/Phogheus/GoBlog/models"
)

type blogRepository struct {
	DB *sql.DB
}

func NewBlogRepository(db *sql.DB) models.IBlogRepository {
	return &blogRepository{
		DB: db,
	}
}

func (repo *blogRepository) InsertNewBlogPost(request models.CreateBlogPostRequest) (int, error) {
	return 0, nil
}

func (repo *blogRepository) GetBlogPostById(id int) (*models.BlogPost, error) {
	var post models.BlogPost
	return &post, nil
}

func (repo *blogRepository) GetBlogPostsByIds(id []int) (*[]models.BlogPost, error) {
	var posts []models.BlogPost
	return &posts, nil
}

func (repo *blogRepository) GetBlogPostsByAuthorId(id int) (*[]models.BlogPost, error) {
	var posts []models.BlogPost
	return &posts, nil
}

func (repo *blogRepository) UpdateBlogPost(post models.BlogPost) (bool, error) {
	return false, nil
}

func (repo *blogRepository) DeleteBlogPostById(id int) (bool, error) {
	return false, nil
}

func (repo *blogRepository) DeleteBlogPostsByIds(id []int) (bool, error) {
	return false, nil
}

func (repo *blogRepository) DeleteBlogPostsByAuthorId(id int) (bool, error) {
	return false, nil
}
