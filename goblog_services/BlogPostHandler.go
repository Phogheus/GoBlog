package goblog_services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Phogheus/GoBlog/goblog_data"
)

func HandleBlogRequest(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		head, _ := ShiftPath(request.URL.Path) // Ignore any trailing route info

		postId, err1 := strconv.Atoi(head)

		if err1 != nil {
			errMsg := ErrorResponse{err1.Error()}
			json.NewEncoder(writer).Encode(errMsg)
		}

		post := goblog_data.GetBlogPostById(postId)

		if post.Id == -1 {
			errMsg := ErrorResponse{"Post with id " + head + " not found."}
			json.NewEncoder(writer).Encode(errMsg)
		} else {
			json.NewEncoder(writer).Encode(post)
		}
	} else if request.Method == "POST" {
		var newPost goblog_data.BlogPost
		err := json.NewDecoder(request.Body).Decode(&newPost)

		if err != nil {
			errMsg := ErrorResponse{err.Error()}
			json.NewEncoder(writer).Encode(errMsg)
		} else {
			createSuccessful, newPostId := goblog_data.CreateNewBlogPost(newPost)

			if !createSuccessful || newPostId == -1 {
				errMsg := ErrorResponse{"Failed to create new post."}
				json.NewEncoder(writer).Encode(errMsg)
			}

			writer.WriteHeader(http.StatusCreated)
		}
	} else {
		json.NewEncoder(writer).Encode(BadMethodError)
	}
}