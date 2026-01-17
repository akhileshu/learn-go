package readingfiles

import (
	// blogposts "github.com/quii/learn-go-with-tests/reading-files"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	/*
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("hi")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		posts, err := NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}

	*/

	/*
		t.Run("get one file", func(t *testing.T) {
			fs := fstest.MapFS{
				"hello world.md":  {Data: []byte("Title: Post 1")},
				"hello-world2.md": {Data: []byte("Title: Post 2")},
			}

			// rest of test code cut for brevity
			posts, _ := NewPostsFromFS(fs)
			got := posts[0]
			want := Post{Title: "Post 1"}

			assertPost(t, got , want)
		})
	*/

	t.Run("get one file", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
			secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
		)

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, _ := NewPostsFromFS(fs)

		// rest of test code cut for brevity
		assertPost(t, posts[0], Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		})

	})

}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
