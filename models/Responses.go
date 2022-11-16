package models

type GenericResponse struct {
	ResponseMessage string
}

type ErrorResponse struct {
	ErrorCode int
	GenericResponse
}

// Blog User Responses

type BlogUserCreatedResponse struct {
	UserId int
	GenericResponse
}

type BlogUserRetrievedResponse struct {
	User BlogUserExternal
	GenericResponse
}

type BlogUserUpdatedResponse struct {
	GenericResponse
}

type BlogUserSetActiveStatusResponse struct {
	GenericResponse
}

// Blog Post Responses

type BlogPostCreatedResponse struct {
	PostId int
	GenericResponse
}

type BlogPostRetrievedResponse struct {
	Post BlogPost
	GenericResponse
}

type BlogPostUpdatedResponse struct {
	GenericResponse
}

type BlogPostDeletedResponse struct {
	GenericResponse
}
