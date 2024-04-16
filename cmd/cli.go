package main

import (
	"fmt"
	"log"
	"os"
	"otto/cmd/internal/dataFetcher"
	"otto/cmd/internal/parseArgs"
)

func main() {
	args, err := parseArgs.ParseArgs(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if args.UserID == 0 {
		log.Fatal("UserID flag is missing, e.g. -userId 1")
	}

	posts, err := dataFetcher.GetPostsWComments(args.UserID, "")

	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		fmt.Printf("Title: %s \nBody: %s \nComments:\n", post.Title, post.Body)
		for _, comment := range post.Comments {
			fmt.Printf("\tMail: %s \n\tName: %s \n\tBody: %s \n\n", comment.Email, comment.Name, comment.Body)
		}
		fmt.Print("\n\n")
	}
}
