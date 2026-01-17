package main

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	readingfiles "go-playground/17_reading_files"
	"testing"
)

// func TestExample(t *testing.T) {
// 	t.Fatal("not implemented")
// }

func TestRender(t *testing.T) {
	var (
		aPost = readingfiles.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	// TestRender.it_converts_a_single_post_into_HTML.approved.txt contains expected html
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	/*
			t.Run("it converts a single post into HTML", func(t *testing.T) {
				buf := bytes.Buffer{}
				err := Render(&buf, aPost)

				if err != nil {
					t.Fatal(err)
				}

				got := buf.String()
				want := `<h1>hello world</h1>
		<p>This is a description</p>
		Tags: <ul><li>go</li><li>tdd</li></ul>`

				if got != want {
					t.Errorf("got '%s' want '%s'", got, want)
				}
			})
	*/
}
