package models

type IUserRepository interface {
	CreateNewUser(request CreateBlogUserRequest) (bool, error)
	GetUserById(id int) (*BlogUser, error)
	GetUserByUsername(username string) (*BlogUser, error)
	GetUserByEmail(email string) (*BlogUser, error)
	UpdateUser(request UpdateBlogUserRequest) (bool, error)
	DeleteUserById(id int) (bool, error)
	LockUserById(id int) (bool, error)
	UnlockUserById(id int) (bool, error)
	DisableUserById(id int) (bool, error)
	EnableUserById(id int) (bool, error)
}

type IBlogRepository interface {
	InsertNewBlogPost(request CreateBlogPostRequest) (int, error)
	GetBlogPostById(id int) (*BlogPost, error)
	GetBlogPostsByIds(id []int) (*[]BlogPost, error)
	GetBlogPostsByAuthorId(id int) (*[]BlogPost, error)
	UpdateBlogPost(post BlogPost) (bool, error)
	DeleteBlogPostById(id int) (bool, error)
	DeleteBlogPostsByIds(id []int) (bool, error)
	DeleteBlogPostsByAuthorId(id int) (bool, error)
}

type IUserService interface {
	CreateUser(request CreateBlogUserRequest) BlogUserCreatedResponse
	GetUser(request GetBlogUserRequest) BlogUserRetrievedResponse
	UpdateUser(request UpdateBlogUserRequest) BlogUserUpdatedResponse
	SetUserActiveStatus(request SetBlogUserActiveStatusRequest) BlogUserSetActiveStatusResponse
}

type IBloggingService interface {
	CreateNewBlogPost(request CreateBlogPostRequest) BlogPostCreatedResponse
	GetBlogPosts(request GetBlogPostsRequest) BlogPostRetrievedResponse
	UpdateBlogPost(request UpdateBlogPostRequest) BlogPostUpdatedResponse
	DeleteBlogPosts(request DeleteBlogPostsRequest) BlogPostDeletedResponse
}
