package dataFetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func convertJSONResToStruct[T Post | Comment](url string) ([]T, error) {
	var contents []T

	res, err := http.Get(url) // no timeout
	if err != nil {
		return nil, err
	}

	defer res.Body.Close() // do not check for error, as it would cause no side effects

	body, err := io.ReadAll(res.Body)

	err = json.Unmarshal(body, &contents)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func getEndpointWithQueryParameters(posts []Post) string {
	endpoint := "https://jsonplaceholder.typicode.com/comments?"

	for _, post := range posts {
		// every comment has a postId, which is the id of the original post, under which the comment was written
		endpoint += fmt.Sprintf("&postId=%d", post.ID)
	}

	return endpoint
}

/*
GetPosts gets all posts for a specific UserID from the JSONPlaceholder API and returns it as an array of Post's.
*/
func GetPosts(userID int) ([]Post, error) {
	endpoint := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", userID)

	return convertJSONResToStruct[Post](endpoint)
}

/*
GetComments gets all comments, which are related to the passed Post's from the JSONPlaceholder API and returns it as an array of Comment's.
*/
func GetComments(posts []Post) ([]Comment, error) {
	endpoint := getEndpointWithQueryParameters(posts)
	return convertJSONResToStruct[Comment](endpoint)
}
