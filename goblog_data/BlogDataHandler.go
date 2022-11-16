package goblog_data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

const (
	TIME_FORMAT              string = "02-01-2006 15:04:05"
	CONNECTION_STRING_FORMAT string = "%v:%v@tcp(%v:%d)/%v"

	FAILED_TO_LOAD_CONFIG_ERROR          string = "Failed to load configuration."
	FAILED_TO_CONNECT_TO_DB_ERROR_FORMAT string = "Failed to connect to database with error: %v"

	INSERT_BLOG_POST_QUERY         string = "call InsertNewBlogPost(?, ?, ?, ?)"
	INSERT_BLOG_POST_FAILED_FORMAT string = "Call to InsertNewBlogPost failed with error: %v"

	SELECT_BLOG_POST_QUERY         string = "SELECT * FROM BlogPosts WHERE Id = ?"
	SELECT_BLOG_POST_FAILED_FORMAT string = "Call to GetBlogPostById failed with error: %v"

	UPDATE_BLOG_POST_QUERY         string = "call UpdateBlogPost(?, ?, ?, ?)"
	UPDATE_BLOG_POST_FAILED_FORMAT string = "Call to UpdateBlogPost failed with error: %v"

	DELETE_BLOG_POST_QUERY         string = "DELETE FROM BlogPosts WHERE Id = ?"
	DELETE_BLOG_POST_FAILED_FORMAT string = "Call to DeleteBlogPostById failed with error: %v"
)

var connectionString string

func CreateNewBlogPost(post BlogPost) (bool, int) {
	var insertSuccessful bool
	nextId := -1

	if post.Author == "" || post.Title == "" || post.Body == "" {
		return insertSuccessful, nextId
	}

	db, dbConnected := getDbConnection()
	defer db.Close()

	if !dbConnected {
		return insertSuccessful, nextId
	}

	result, err := db.Exec(INSERT_BLOG_POST_QUERY, post.Author, time.Now().UTC().Format(TIME_FORMAT), post.Title, post.Body)

	if err != nil {
		log.Printf(INSERT_BLOG_POST_FAILED_FORMAT, err)
		return insertSuccessful, nextId
	}

	lastInsertId, _ := result.LastInsertId()

	if lastInsertId >= 0 {
		insertSuccessful = true
		nextId = int(lastInsertId)
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

	rows, err := db.Query(SELECT_BLOG_POST_QUERY, id)
	defer rows.Close()

	if err != nil {
		log.Printf(SELECT_BLOG_POST_FAILED_FORMAT, err)
		return post
	}

	var datePostTimeString string
	var dateLastUpdatedString sql.NullString

	if rows.Next() {
		err := rows.Scan(&post.Id, &post.Author, &datePostTimeString, &dateLastUpdatedString, &post.Title, &post.Body)

		if err != nil {
			log.Printf(SELECT_BLOG_POST_FAILED_FORMAT, err)
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

	result, err := db.Exec(UPDATE_BLOG_POST_QUERY, post.Id, time.Now().UTC().Format(TIME_FORMAT), post.Title, post.Body)
	count, err := result.RowsAffected()

	if err != nil {
		log.Printf(UPDATE_BLOG_POST_FAILED_FORMAT, err)
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

	result, err := db.Exec(DELETE_BLOG_POST_QUERY, id)
	count, err := result.RowsAffected()

	if err != nil {
		log.Printf(DELETE_BLOG_POST_FAILED_FORMAT, err)
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
		log.Fatal(FAILED_TO_LOAD_CONFIG_ERROR)
	}

	user := viper.GetString("db_user")
	pass := viper.GetString("db_pass")
	host := viper.GetString("db_host")
	port := viper.GetInt("db_port")
	entry := viper.GetString("entry_database")
	connectionString = fmt.Sprintf(CONNECTION_STRING_FORMAT, user, pass, host, port, entry)
}

func getDbConnection() (*sql.DB, bool) {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Printf(FAILED_TO_CONNECT_TO_DB_ERROR_FORMAT, err)
	}

	return db, err == nil
}
