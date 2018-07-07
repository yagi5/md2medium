package main

import (
	"os"

	medium "github.com/medium/medium-sdk-go"
)

func main() {
	m := medium.NewClientWithAccessToken(os.Getenv("MEDIUM_ACCESS_TOKEN"))

	// empty means current user
	u, err := m.GetUser("")
	panicOnErr(err)

	_, err = m.CreatePost(medium.CreatePostOptions{
		UserID:        u.ID,
		Title:         "Title",
		Content:       "<h2>Title</h2><p>Content</p>",
		ContentFormat: medium.ContentFormatHTML,
		PublishStatus: medium.PublishStatusPublic,
	})
	panicOnErr(err)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
