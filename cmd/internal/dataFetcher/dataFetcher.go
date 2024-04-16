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

func GetPosts(UserID int) ([]Post, error) {
	endpoint := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", UserID)
	return convertJSONResToStruct[Post](endpoint)
}

func GetComments(UserID int) ([]Comment, error) {
	endpoint := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?userId=%d", UserID)
	return convertJSONResToStruct[Comment](endpoint)
}
