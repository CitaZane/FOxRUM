package models

/* ------------------------ register function on room ----------------------- */
type Comment interface {
}

/* ------------- register functions that interact with database ------------- */
type CommentHandlers interface {
	GetComments(postId, offset int) []Comment
	AddComment(postId int, content, author string)
	GetLastComment(postId int) Comment

	GetUserCommentCount(username string) int
}
