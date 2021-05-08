package main

// article structure
type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// initialize an article list
// fill the list with two hardcoded articles
// For this project, I'm storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static file
var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// return a list of all articles
func getAllArticles() []article {
	return articleList
}
