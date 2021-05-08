package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HML status to 200
		http.StatusOK,
		// use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}
