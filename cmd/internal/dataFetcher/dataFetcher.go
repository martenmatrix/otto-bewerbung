package dataFetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	UserID   int    `json:"userId"`
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Comments []Comment
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

func getEndpointWithQueryParameters(posts []Post, domain string) string {
	endpoint := domain + "/comments?"

	for _, post := range posts {
		// every comment has a postId, which is the id of the original post, under which the comment was written
		endpoint += fmt.Sprintf("&postId=%d", post.ID)
	}

	return endpoint
}

/*
getComments gets all comments, which are related to the passed Post's from the JSONPlaceholder API and returns it as an array of Comment's.
*/
func getComments(posts []Post, domain string) ([]Comment, error) {
	endpoint := getEndpointWithQueryParameters(posts, domain)
	return convertJSONResToStruct[Comment](endpoint)
}

/*
getPosts gets all posts without comments for a specific UserID from the JSONPlaceholder API and returns it as an array of Post's.
*/
func getPosts(userID int, domain string) ([]Post, error) {
	endpoint := fmt.Sprintf(domain+"/posts?userId=%d", userID)

	return convertJSONResToStruct[Post](endpoint)
}

/*
GetPostsWComments gets all posts including its comments for a specific UserID from the JSONPlaceholder API and returns it as an array of Post's.

The userID specifies by which author the posts should be filtered.

The domain is the domain name from the API, this parameter is only used for testing purposes, otherwise just input an empty string.
*/
func GetPostsWComments(userID int, url string) ([]Post, error) {
	if url == "" {
		// not in testing environment
		url = "https://jsonplaceholder.typicode.com"
	}

	posts, err := getPosts(userID, url)
	if err != nil {
		return nil, err
	}

	comments, err := getComments(posts, url)
	if err != nil {
		return nil, err
	}

	// postID represents key in map, id of original post
	sortedComments := make(map[int][]Comment)

	for _, comment := range comments {
		sortedComments[comment.PostID] = append(sortedComments[comment.PostID], comment)
	}

	var postsWComments []Post
	for _, post := range posts {
		post.Comments = sortedComments[post.ID]
		postsWComments = append(postsWComments, post)
	}

	return postsWComments, nil
}
