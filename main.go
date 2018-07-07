package main

import (
	"log"
	"os"

	medium "github.com/medium/medium-sdk-go"
)

func main() {
	m := medium.NewClientWithAccessToken(os.Getenv("MEDIUM_ACCESS_TOKEN"))

	// empty means current user
	u, err := m.GetUser("")
	if err != nil {
		log.Fatal(err)
	}

	_, err = m.CreatePost(medium.CreatePostOptions{
		UserID:        u.ID,
		Title:         "Title",
		Content:       "<h2>Title</h2><p>Content</p>",
		ContentFormat: medium.ContentFormatHTML,
		PublishStatus: medium.PublishStatusPublic,
	})
	if err != nil {
		log.Fatal(err)
	}
}
