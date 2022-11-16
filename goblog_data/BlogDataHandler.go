package goblog_data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

const TIME_FORMAT string = "02-01-2006 15:04:05"

var dataStore = make(map[int]BlogPost)
var connectionString string

func CreateNewBlogPost(post BlogPost) (bool, int) {
	var insertSuccessful bool
	nextId := -1

	if !validateBlogPost(post) {
		return insertSuccessful, nextId
	}

	db, dbConnected := getDbConnection()
	defer db.Close()

    if !dbConnected {
		return insertSuccessful, nextId
    }

	rows, err := db.Query("call InsertNewBlogPost(?, ?, ?, ?)", post.Author, time.Now().UTC().Format(TIME_FORMAT), post.Title, post.Body)
	defer rows.Close();

	if err != nil {
		log.Printf("Call to InsertNewBlogPost failed with error: %v", err)
		return insertSuccessful, nextId
	}

	if rows.Next() {
		err := rows.Scan(&nextId)

		if err != nil {
			log.Printf("Scan of returned value failed with error: %v", err)
			nextId = -1
		} else {
			insertSuccessful = true
		}
	} else {
		log.Println("Failed to insert new blog post (no id returned).")
	}

	return insertSuccessful, nextId
}

func GetBlogPostById(id int) BlogPost {
	var post BlogPost
	post.Id = -1

	db, dbConnected := getDbConnection()
	defer db.Close()

    if !dbConnected {
		return post
    }

	rows, err := db.Query("SELECT * FROM BlogPosts WHERE Id = ?", id)
	defer rows.Close();

	if err != nil {
		log.Printf("Call to GetBlogPostById failed with error: %v", err)
		return post
	}

	var datePostTimeString string
	var dateLastUpdatedString sql.NullString

	if rows.Next() {
		err := rows.Scan(&post.Id, &post.Author, &datePostTimeString, &dateLastUpdatedString, &post.Title, &post.Body)

		if err != nil {
			log.Printf("Scan of returned value failed with error: %v", err)
		} else {
			datePostTime, err := time.Parse(TIME_FORMAT, datePostTimeString)

			if err == nil {
				post.DatePosted = datePostTime
			}

			if dateLastUpdatedString.Valid {
				dateLastUpdated, err := time.Parse(TIME_FORMAT, dateLastUpdatedString.String)

				if err == nil {
					post.DateLastUpdated = dateLastUpdated
				}
			}
		}
	}

	return post
}

func UpdateBlogPost(post BlogPost) bool {
	var updateSuccessful bool

	db, dbConnected := getDbConnection()
	defer db.Close()

    if !dbConnected {
		return updateSuccessful
    }

	result, err := db.Exec("call UpdateBlogPost(?, ?, ?, ?)", post.Id, time.Now().UTC().Format(TIME_FORMAT), post.Title, post.Body)
	count, err := result.RowsAffected()

	if err != nil {
		log.Printf("Call to UpdateBlogPost failed with error: %v", err)
	} else {
		updateSuccessful = count > 0
	}

	return updateSuccessful
}

func DeleteBlogPostById(id int) bool {
	var deleteSuccessful bool

	db, dbConnected := getDbConnection()
	defer db.Close()

    if !dbConnected {
		return deleteSuccessful
    }

	result, err := db.Exec("DELETE FROM BlogPosts WHERE Id = ?", id)
	count, err := result.RowsAffected()

	if err != nil {
		log.Printf("Call to DeleteBlogPostById failed with error: %v", err)
	} else {
		deleteSuccessful = count > 0
	}

	return deleteSuccessful
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Failed to load configuration.")
    }

	user := viper.GetString("db_user")
	pass := viper.GetString("db_pass")
	host := viper.GetString("db_host")
	port := viper.GetInt("db_port")
	entry := viper.GetString("entry_database")
	connectionString = fmt.Sprintf("%v:%v@tcp(%v:%d)/%v", user, pass, host, port, entry)
}

func validateBlogPost(post BlogPost) bool {
	isValid := true

	if post.Author == "" || post.Title == "" || post.Body == "" {
		isValid = false
	}

	if !isValid {
		log.Print("Attempted to post invalid BlogPost")
	}

	return isValid
}

func getDbConnection() (*sql.DB, bool) {
	db, err := sql.Open("mysql", connectionString)
	
    if err != nil {
		log.Printf("Failed to connect to database with error: %v", err)
    }

	return db, err == nil
}
