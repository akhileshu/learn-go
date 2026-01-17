package readingfiles

import (
	// blogposts "github.com/quii/learn-go-with-tests/reading-files"
	// "testing"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"

	// "io"
	"io/fs"
	// "testing/fstest"

	// blogposts "github.com/quii/fstest-spike"
	"log"
	"os"
)

// from from package root
//
// learn-go/17_reading_files on main ❯ go run .
func main() {
	// entry point
	posts, err := NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range posts {
		log.Println("==============post==============")
		log.Println(post)
	}
}

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f)
		if err != nil {
			return nil, err //todo: needs clarification, should we totally fail if one file fails? or just ignore?
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postFile fs.File) (Post, error) {
	/*
		postData, err := io.ReadAll(postFile)
		if err != nil {
			return Post{}, err
		}

		post := Post{Title: string(postData)[7:]}
		return post, nil
	*/

	scanner := bufio.NewScanner(postFile)

	/*
		readLine := func() string {
				scanner.Scan()
				return scanner.Text()
			}
				title := readLine()[len(titleSeparator):]
			description := readLine()[len(descriptionSeparator):]

			return Post{Title: title, Description: description}, nil
	*/

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}
	/*
		readBody := func() string {
			var body strings.Builder

			for scanner.Scan() {
				if scanner.Text() == "---" {
					break // Exit the loop when the condition is met
				}
				// Process the scanned line/token here if needed
			}

			for scanner.Scan() {
				if body.String() != "" {
					body.WriteString("\n")
				}
				body.WriteString(scanner.Text())
			}
			return body.String()
		}
	*/

	readBody := func() string {
		scanner.Scan() // ignore a line
		buf := bytes.Buffer{}
		for scanner.Scan() {
			fmt.Fprintln(&buf, scanner.Text())
		}
		return strings.TrimSuffix(buf.String(), "\n")
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(),
	}, nil
}
