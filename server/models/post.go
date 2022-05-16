package models

type Post interface {
	GetAuthor() string
	GetTitle() string
	GetContent() string
	GetCategory() string
}

type PostHandlers interface {
	AddPost(post Post)
	AddCategory(category string)
	SelectCategory(category string) string
	SelectAllCategories() []string
	GetPosts(offset int) []Post
	GetCategoryPosts(category string, offset int) []Post
	GetUserPosts(offset int, username string) []Post
	GetPost(postId int) Post

	GetUserPostCount(username string) int
}
