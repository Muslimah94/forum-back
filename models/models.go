package models

// Users ...
type Users struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}

// Roles ...
type Roles struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Categories ...
type Categories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Posts ...
type Posts struct {
	ID           int    `json:"id"`
	AuthorID     int    `json:"author_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CreationDate string `json:"creation_date"`
	AuthorNick   string `json:"author_nick"`
}

// Comments ...
type Comments struct {
	ID       int    `json:"id"`
	AuthorID int    `json:"author_id"`
	PostID   int    `json:"post_id"`
	Content  string `json:"content"`
}

// Reactions ...
type Reactions struct {
	ID        int `json:"id"`
	AuthorID  int `json:"author_id"`
	Type      int `json:"type"`
	PostID    int `json:"post_id"`
	CommentID int `json:"comment_id"`
}

// PostsCategories ...
type PostsCategories struct {
	PostID     int `json:"post_id"`
	CategoryID int `json:"category_id"`
}