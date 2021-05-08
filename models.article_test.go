package main

import "testing"

// Test the function that gets all articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// check that the length of the articles returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// check that each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}
