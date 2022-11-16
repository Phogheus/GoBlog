package goblog_services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Phogheus/GoBlog/goblog_data"
)

func HandleBlogRequest(writer http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case "GET":
		processGET(writer, request)

	case "POST":
		processPOST(writer, request)

	case "PATCH":
		processPATCH(writer, request)

	case "DELETE":
		processDELETE(writer, request)

	default:
		json.NewEncoder(writer).Encode(BadMethodError)
	}
}

func processGET(writer http.ResponseWriter, request *http.Request) {
	head, _ := ShiftPath(request.URL.Path) // Ignore any trailing route info

	if head == "" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	postId, err1 := strconv.Atoi(head)

	if err1 != nil {
		errMsg := ErrorResponse{err1.Error()}
		json.NewEncoder(writer).Encode(errMsg)
	}

	post := goblog_data.GetBlogPostById(postId)

	if post.Id == -1 {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(writer).Encode(post)
	}
}

func processPOST(writer http.ResponseWriter, request *http.Request) {
	var newPost goblog_data.BlogPost
	err := json.NewDecoder(request.Body).Decode(&newPost)

	if err != nil {
		errMsg := ErrorResponse{err.Error()}
		json.NewEncoder(writer).Encode(errMsg)
	} else {
		createSuccessful, newPostId := goblog_data.CreateNewBlogPost(newPost)

		if !createSuccessful || newPostId == -1 {
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			writer.WriteHeader(http.StatusCreated)
		}
	}
}

func processPATCH(writer http.ResponseWriter, request *http.Request) {
	var postUpdate goblog_data.BlogPost
	err := json.NewDecoder(request.Body).Decode(&postUpdate)

	if err != nil {
		errMsg := ErrorResponse{err.Error()}
		json.NewEncoder(writer).Encode(errMsg)
	} else {
		updateSuccessful := goblog_data.UpdateBlogPost(postUpdate)

		if !updateSuccessful {
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			writer.WriteHeader(http.StatusOK)
		}
	}
}

func processDELETE(writer http.ResponseWriter, request *http.Request) {
	head, _ := ShiftPath(request.URL.Path)

	postId, err1 := strconv.Atoi(head)

	if err1 != nil {
		errMsg := ErrorResponse{err1.Error()}
		json.NewEncoder(writer).Encode(errMsg)
	}

	deleteSuccessful := goblog_data.DeleteBlogPostById(postId)

	if !deleteSuccessful {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}
