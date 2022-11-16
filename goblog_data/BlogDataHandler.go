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
	SELECT_BLOG_POST_QUERY         string = "SELECT * FROM BlogPosts WHERE Id = ?"
	UPDATE_BLOG_POST_QUERY         string = "call UpdateBlogPost(?, ?, ?, ?)"
	DELETE_BLOG_POST_QUERY         string = "DELETE FROM BlogPosts WHERE Id = ?"
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

	rows, _ := db.Query(INSERT_BLOG_POST_QUERY, post.Author, time.Now().UTC().Format(TIME_FORMAT), post.Title, post.Body)

	if rows.Next() {
		err := rows.Scan(&nextId)

		if err == nil {
			insertSuccessful = true
		}
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

	rows, _ := db.Query(SELECT_BLOG_POST_QUERY, id)
	defer rows.Close()

	if rows.Next() {
		var datePostTimeString string
		var dateLastUpdatedString sql.NullString

		err := rows.Scan(&post.Id, &post.Author, &datePostTimeString, &dateLastUpdatedString, &post.Title, &post.Body)

		if err == nil {
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
	return doExecWithRowCheck(UPDATE_BLOG_POST_QUERY, post.Id, time.Now().UTC().Format(TIME_FORMAT), post.Title, post.Body)
}

func DeleteBlogPostById(id int) bool {
	return doExecWithRowCheck(DELETE_BLOG_POST_QUERY, id)
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
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
	db, _ := sql.Open("mysql", connectionString)
	err := db.Ping()

	if err != nil {
		log.Printf(FAILED_TO_CONNECT_TO_DB_ERROR_FORMAT, err)
	}

	return db, err == nil
}

func doExecWithRowCheck(cmd string, args ...any) bool {
	db, dbConnected := getDbConnection()
	defer db.Close()

	if dbConnected {
		result, err := db.Exec(cmd, args...)

		if err == nil {
			count, _ := result.RowsAffected()
			return count > 0
		}
	}

	return false
}
