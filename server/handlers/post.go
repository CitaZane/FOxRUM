package handlers

import (
	"database/sql"
	"real-time-forum/models"
	"strconv"
	"time"
)

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Author   string `json:"author"`
	Time     string `json:"time"`
}

func (post *Post) GetAuthor() string {
	return post.Author
}
func (post *Post) GetTitle() string {
	return post.Title
}
func (post *Post) GetContent() string {
	return post.Content
}
func (post *Post) GetCategory() string {
	return post.Category
}

type PostHandler struct {
	Db *sql.DB
}

/* ----------------------------- insert into db ----------------------------- */
func (handler *PostHandler) AddPost(post models.Post) {
	stmt, err := handler.Db.Prepare("INSERT INTO posts( author, title, content, category) values(?,?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(post.GetAuthor(), post.GetTitle(), post.GetContent(), post.GetCategory())
	checkErr(err)
}

func (handler *PostHandler) AddCategory(category string) {
	stmt, err := handler.Db.Prepare("INSERT INTO categories( name) values(?)")
	checkErr(err)

	_, err = stmt.Exec(category)
	checkErr(err)
}

/* --------------------------------- select --------------------------------- */
/* ------------------------------- categories ------------------------------- */
func (handler *PostHandler) SelectCategory(category string) string {
	row := handler.Db.QueryRow("SELECT name FROM categories where name = ? LIMIT 1", category)
	var catFound string
	if err := row.Scan(&catFound); err != nil {
		if err == sql.ErrNoRows {
			return ""
		}
	}
	return catFound
}

func (handler *PostHandler) SelectAllCategories() []string {
	rows, err := handler.Db.Query("SELECT name FROM categories")
	checkErr(err)
	var categories []string
	defer rows.Close()
	for rows.Next() {
		var category string
		rows.Scan(&category)
		categories = append(categories, category)
	}
	return categories
}

/* ---------------------------------- posts --------------------------------- */

func (handler *PostHandler) GetCategoryPosts(category string, offset int) []models.Post {
	rows, err := handler.Db.Query("SELECT author, title, content, created_at, post_id FROM posts WHERE category = ? ORDER BY created_at DESC LIMIT 30 OFFSET ?", category, offset)
	checkErr(err)
	var posts []models.Post
	defer rows.Close()
	for rows.Next() {
		var post Post
		var dateTime time.Time
		rows.Scan(&post.Author, &post.Title, &post.Content, &dateTime, &post.Id)
		post.Time = strconv.FormatInt(dateTime.Unix(), 10)
		posts = append(posts, &post)
	}
	return posts
}

func (handler *PostHandler) GetPosts(offset int) []models.Post {
	rows, err := handler.Db.Query("SELECT author, title, content, created_at, post_id FROM posts ORDER BY created_at DESC LIMIT 30 OFFSET ?", offset)
	checkErr(err)
	var posts []models.Post
	defer rows.Close()
	for rows.Next() {
		var post Post
		var dateTime time.Time
		rows.Scan(&post.Author, &post.Title, &post.Content, &dateTime, &post.Id)
		post.Time = strconv.FormatInt(dateTime.Unix(), 10)
		posts = append(posts, &post)
	}
	return posts
}

func (handler *PostHandler) GetUserPosts(offset int, username string) []models.Post {
	rows, err := handler.Db.Query("SELECT author, title, content, created_at, post_id FROM posts WHERE author = ? ORDER BY created_at DESC LIMIT 30 OFFSET ?", username, offset)
	checkErr(err)
	var posts []models.Post
	defer rows.Close()
	for rows.Next() {
		var post Post
		var dateTime time.Time
		rows.Scan(&post.Author, &post.Title, &post.Content, &dateTime, &post.Id)
		post.Time = strconv.FormatInt(dateTime.Unix(), 10)
		posts = append(posts, &post)
	}
	return posts
}

func (handler *PostHandler) GetPost(postId int) models.Post {
	row := handler.Db.QueryRow("SELECT author, title, content, created_at, post_id, category FROM posts WHERE post_id = ? LIMIT 1", postId)
	var post Post
	var dateTime time.Time
	if err := row.Scan(&post.Author, &post.Title, &post.Content, &dateTime, &post.Id, &post.Category); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
	}
	post.Time = strconv.FormatInt(dateTime.Unix(), 10)
	return &post
}

func (handler *PostHandler) GetUserPostCount(username string) int {
	row := handler.Db.QueryRow("SELECT COUNT(*) FROM posts where author = ?", username)
	var postCount int
	if err := row.Scan(&postCount); err != nil {
		if err == sql.ErrNoRows {
			return 0
		}
	}
	return postCount
}
