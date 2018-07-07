package main

import (
	"bufio"
	"os"

	medium "github.com/medium/medium-sdk-go"
	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

type entry struct {
	metadata metadata
	content  string
}

type metadata struct {
	Title string   `yaml:"title"`
	Date  string   `yaml:"date"`
	Tags  []string `yaml:"tags"`
}

func main() {
	app := cli.NewApp()
	app.Name = "md2medium"
	app.Usage = "publish your markdown to medium"
	app.Version = "1.0.1"
	app.Flags = flags()
	app.Action = run
	app.Run(os.Args)
}

// flags cli flags
func flags() []cli.Flag {
	flags := []cli.Flag{
		cli.BoolFlag{
			Name:  "draft",
			Usage: "if this flag is true, files will be published as draft.",
		},
		cli.StringFlag{
			Name:  "file",
			Usage: "specify file path",
		},
	}
	return flags
}

// run entry point
func run(c *cli.Context) error {
	title := c.String("title")
	file := c.String("file")
	isDraft := c.Bool("draft")
	return publish(title, file, isDraft)
}

// publish
func publish(title string, file string, isDraft bool) error {
	m := medium.NewClientWithAccessToken(os.Getenv("MEDIUM_ACCESS_TOKEN"))

	// empty means current user
	u, err := m.GetUser("")
	if err != nil {
		return err
	}

	var status medium.PublishStatus = medium.PublishStatusPublic
	if isDraft {
		status = medium.PublishStatusDraft
	}
	entry, err := readFile(file)
	if err != nil {
		return err
	}

	_, err = m.CreatePost(medium.CreatePostOptions{
		UserID:        u.ID,
		Title:         entry.metadata.Title,
		Content:       entry.content,
		ContentFormat: medium.ContentFormatMarkdown,
		PublishStatus: status,
		Tags:          entry.metadata.Tags,
	})
	return err
}

// readFile
// parse given file and return entry
func readFile(path string) (entry, error) {
	var e entry
	var m metadata
	file, err := os.Open(path)
	if err != nil {
		return e, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var metayml string
	index := 1
	for scanner.Scan() {
		// skip first line
		if index == 1 {
			index++
			continue
		}
		if scanner.Text() == "---" {
			break
		}
		metayml += scanner.Text() + "\n"
		index++
	}

	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	m, err = parseYaml(metayml)
	if err != nil {
		return e, err
	}
	e.content = content
	e.metadata = m

	return e, nil
}

// parse header metadata yaml
func parseYaml(metayml string) (metadata, error) {
	m := metadata{}
	err := yaml.Unmarshal([]byte(metayml), &m)
	if err != nil {
		return m, err
	}
	return m, nil
}
