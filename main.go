package main

import (
	"os"

	medium "github.com/medium/medium-sdk-go"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "md2medium"
	app.Usage = "publish your markdown to medium"
	app.Version = "1.0.1"
	app.Flags = flags()
	app.Action = run
	app.Run(os.Args)
}

func flags() []cli.Flag {
	flags := []cli.Flag{
		cli.BoolFlag{
			Name:  "draft",
			Usage: "if this flag is true, files will be published as draft.",
		},
		cli.StringFlag{
			Name:  "title",
			Usage: "specify title",
		},
		cli.StringFlag{
			Name:  "file",
			Usage: "specify file path",
		},
	}
	return flags
}

func run(c *cli.Context) error {
	title := c.String("title")
	file := c.String("file")
	isDraft := c.Bool("draft")
	fmt.Println(title)
	fmt.Println(file)
	fmt.Println(isDraft)
	return publish(title, file, isDraft)
}

func publish(title string, file string, isDraft bool) error {
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
