package goblog_data

import "time"

var dataStore = make(map[int]BlogPost)
var currentId int

func CreateNewBlogPost(post BlogPost) (bool, int) {
	// Error handling concerns:
	//	Bad Author
	//	Bad Title
	//	Bad Body

	nextId := getNextId()
	post.Id = nextId

	dataStore[nextId] = post

	return true, nextId
}

func GetBlogPostById(id int) BlogPost {
	existingPost, postExists := dataStore[id]

	if !postExists {
		existingPost.Id = -1
	}

	return existingPost
}

func UpdateBlogPost(post BlogPost) bool {
	existingPost, postExists := dataStore[post.Id]

	if postExists {
		existingPost.Title = post.Title
		existingPost.Body = post.Body
		existingPost.DateLastUpdated = time.Now()
		dataStore[post.Id] = existingPost
	}

	return postExists
}

func DeleteBlogPostById(id int) bool {
	_, postExists := dataStore[id]

	if postExists {
		delete(dataStore, id)
	}

	return postExists
}

func getNextId() int {
	currentId = currentId + 1
	return currentId
}
