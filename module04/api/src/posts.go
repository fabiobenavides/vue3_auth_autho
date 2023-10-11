package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Post struct {
	ID        int    `json:"id"`
	Timestamp int    `json:"timestamp"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
}

var posts = []*Post{}

func getPosts(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, posts)
}

func submitPost(w http.ResponseWriter, r *http.Request) {
	var post *Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post.Author = "N/A"
	userId, _ := strconv.Atoi(r.Header.Get("UserId"))
	for _, user := range users {
		if user.ID == userId {
			post.Author = user.FullName
			break
		}
	}

	if post.ID > 0 {
		for _, item := range posts {
			if item.ID == post.ID {
				item.Timestamp = int(time.Now().Unix())
				item.Title = post.Title
				item.Content = post.Content
				item.Author = post.Author
				break
			}
		}
	} else {
		post.ID = 1
		for _, item := range posts {
			if item.ID >= post.ID {
				post.ID = item.ID + 1
			}
		}

		posts = append(posts, post)
	}

	savePosts()

	respondJSON(w, http.StatusOK, post)
}

func loadPosts() (err error) {
	content, err := ioutil.ReadFile("posts.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(content, &posts)
}

func savePosts() (err error) {
	content, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("posts.json", content, 0o644)
}
