package goblog_data

import (
	"fmt"
	"testing"
)

func TestBlogPostLifecycle(t *testing.T) {
	newPost := testCreateNewBlogPost(t)

	testGetBlogPostById(t, newPost.Id)

	newPost.Title = "My New Title"
	newPost.Body = "Go was here"
	newPost.Author = "Who wrote this anyway?" // Should not update

	testUpdateBlogPost(t, newPost)

	testDeleteBlogPostById(t, newPost.Id)
}

func TestInvalidCreateNewBlogPost(t *testing.T) {
	var newPost BlogPost = BlogPost{}

	newPostSuccess, _ := CreateNewBlogPost(newPost)

	if newPostSuccess {
		t.Fatal("Create was successful and should not have been.")
	}

	newPost.Author = "Me"

	newPostSuccess, _ = CreateNewBlogPost(newPost)

	if newPostSuccess {
		t.Fatal("Create was successful and should not have been.")
	}

	newPost.Title = "A title"

	newPostSuccess, _ = CreateNewBlogPost(newPost)

	if newPostSuccess {
		t.Fatal("Create was successful and should not have been.")
	}

	newPost.Body = "Some body text"

	newPostSuccess, _ = CreateNewBlogPost(newPost)

	if !newPostSuccess {
		t.Fatal("Create was unsuccessful and should have been.")
	}
}

func TestGetBlogPostByInvalidId(t *testing.T) {
	idToGet := 12345
	getPostById := GetBlogPostById(idToGet)

	if getPostById.Id != -1 {
		t.Fatalf("Got unexpected response for post with id %d", idToGet)
	}
}

func TestUpdateBlogPostByInvalidId(t *testing.T) {
	newPost := BlogPost{}
	updateSuccessful := UpdateBlogPost(newPost)

	if updateSuccessful {
		t.Fatal("Update was successful and should not have been.")
	}
}

func TestDeleteBlogPostByInvalidId(t *testing.T) {
	idToDelete := 12345
	deleteSuccessful := DeleteBlogPostById(idToDelete)

	if deleteSuccessful {
		t.Fatal("Delete was successful and should not have been.")
	}
}

func TestInvalidDbConnections(t *testing.T) {
	tempConnString := connectionString
	connectionString = fmt.Sprintf(CONNECTION_STRING_FORMAT, "garbageuser", "garbagepass", "127.0.0.1", 9999, "garbage")

	defer setConnectionString(tempConnString)

	var newPost BlogPost = BlogPost{
		Author: "Me",
		Title:  "My test post",
		Body:   "This is my post. There are many like it, but this one is mine.",
	}

	newPostSuccess, _ := CreateNewBlogPost(newPost)

	if newPostSuccess {
		t.Fatal("Create was successful and should not have been.")
	}

	getPostById := GetBlogPostById(1)

	if getPostById.Id != -1 {
		t.Fatal("Get was successful and should not have been.")
	}

	updateSuccessful := UpdateBlogPost(newPost)

	if updateSuccessful {
		t.Fatal("Update was successful and should not have been.")
	}

	deleteSuccessful := DeleteBlogPostById(1)

	if deleteSuccessful {
		t.Fatal("Delete was successful and should not have been.")
	}
}

func setConnectionString(str string) {
	connectionString = str
}

func testCreateNewBlogPost(t *testing.T) BlogPost {
	var newPost BlogPost = BlogPost{
		Id:     999, // Setting this directly tests that doing so does not affect the Id chosen during create/add
		Author: "Me",
		Title:  "My test post",
		Body:   "This is my post. There are many like it, but this one is mine.",
	}

	newPostSuccess, newPostId := CreateNewBlogPost(newPost)

	if !newPostSuccess || newPostId == -1 {
		t.Fatal("Failed to create new blog post.")
	} else if newPostId == newPost.Id {
		t.Fatal("Post was created, but has an unexpected id.")
	}

	newPost.Id = newPostId

	return newPost
}

func testGetBlogPostById(t *testing.T, id int) {
	getPostById := GetBlogPostById(id)

	if getPostById.Id == -1 || getPostById.Id != id {
		t.Fatal("Retrieved incorrect/invalid post.")
	}
}

func testUpdateBlogPost(t *testing.T, post BlogPost) {
	updateSuccessful := UpdateBlogPost(post)

	if !updateSuccessful {
		t.Fatal("Update failed.")
	}

	getPostById := GetBlogPostById(post.Id)

	// Author should remain unchanged, only Title and Body should change
	if getPostById.Author == post.Author || getPostById.Title != post.Title || getPostById.Body != post.Body {
		t.Fatal("Post to update with and result do not match.")
	} else if getPostById.DateLastUpdated.Before(getPostById.DatePosted) {
		t.Fatal("Updated date did not change.")
	}
}

func testDeleteBlogPostById(t *testing.T, id int) {
	deleteSuccessful := DeleteBlogPostById(id)

	if !deleteSuccessful {
		t.Fatal("Delete failed")
	}

	getPostById := GetBlogPostById(id)

	if getPostById.Id != -1 {
		t.Fatal("Delete failed")
	}
}
