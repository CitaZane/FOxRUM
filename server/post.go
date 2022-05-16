package main

type Post struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Author   string `json:"author"`
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

/* --------------------------- register  new empty message -------------------------- */
func postTemplate() *Post {
	return &Post{
		Content:  "",
		Title:    "",
		Category: "",
		Author:   "",
	}
}
/* -------------------- check if all fields are filled in ------------------- */
func (post *Post) valid() bool{
	if len(post.GetAuthor()) == 0 {return false} 
	if len(post.GetTitle()) == 0 {return false} 
	if len(post.GetContent()) == 0 {return false} 
	if len(post.GetCategory()) == 0 {return false} 
	return true
}
