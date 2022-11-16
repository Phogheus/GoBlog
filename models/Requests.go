package models

// Blog User Requests

type CreateBlogUserRequest struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

type GetBlogUserRequest struct {
	Id       int
	Username string
	Email    string
}

type UpdateBlogUserRequest struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

// Order of operations: DeleteUser, LockUser and/or DisableUser, UnlockUser and/or EnableUser
type SetBlogUserActiveStatusRequest struct {
	DeleteUser  bool
	LockUser    bool
	UnlockUser  bool
	DisableUser bool
	EnableUser  bool
}

// Blog Post Requests

type CreateBlogPostRequest struct {
	Title string
	Body  string
}

type GetBlogPostsRequest struct {
	PostIds  []int
	AuthorId int
}

type UpdateBlogPostRequest struct {
	Id    int
	Title string
	Body  string
}

type DeleteBlogPostsRequest struct {
	PostIds  []int
	AuthorId int
}
