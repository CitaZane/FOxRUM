package handlers

import (
	"database/sql"
	"fmt"
	"real-time-forum/models"
	"strconv"
	"time"
)

type Comment struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Time    string `json:"time"`
}

type CommentHandler struct {
	Db *sql.DB
}

func (handler *CommentHandler) GetComments(postId int, offset int) []models.Comment {
	rows, err := handler.Db.Query("SELECT author, content, created_at FROM comments WHERE post_id = ? ORDER BY created_at DESC LIMIT 30 OFFSET ?", postId, offset)
	checkErr(err)
	var comments []models.Comment
	defer rows.Close()
	for rows.Next() {
		var comment Comment
		var dateTime time.Time
		rows.Scan(&comment.Author, &comment.Content, &dateTime)
		comment.Time = strconv.FormatInt(dateTime.Unix(), 10)
		comments = append(comments, &comment)
	}
	return comments
}

func (handler *CommentHandler) GetLastComment(postId int) models.Comment {
	row := handler.Db.QueryRow("SELECT author, content, created_at FROM comments WHERE post_id = ? ORDER BY created_at DESC LIMIT 1", postId)
	var comment Comment
	var dateTime time.Time
	if err := row.Scan(&comment.Author, &comment.Content, &dateTime); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
	}
	comment.Time = strconv.FormatInt(dateTime.Unix(), 10)
	return &comment
}

func (handler *CommentHandler) AddComment(postId int, content, author string) {
	stmt, err := handler.Db.Prepare("INSERT INTO comments( author, content, post_id) values(?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(author, content, postId)
	checkErr(err)
}

func (handler *CommentHandler) GetUserCommentCount(username string) int {
	row := handler.Db.QueryRow("SELECT COUNT(*) FROM comments where author = ?", username)
	var postCount int
	if err := row.Scan(&postCount); err != nil {
		if err == sql.ErrNoRows {
			return 0
		}
	}
	return postCount
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("This is a panic!!!")
		panic(err)
	}
}
