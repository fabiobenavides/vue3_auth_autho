package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gorilla/mux"
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

func createPost(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	claims := token.CustomClaims.(*CustomClaims)
	if !claims.HasScope("create:post") {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message":"Insufficient scope."}`))
		return
	}

	var post *Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post.Author = token.RegisteredClaims.Subject
	post.ID = 1
	for _, item := range posts {
		if item.ID >= post.ID {
			post.ID = item.ID + 1
		}
	}

	posts = append(posts, post)

	savePosts()

	respondJSON(w, http.StatusOK, post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	claims := token.CustomClaims.(*CustomClaims)
	if !claims.HasScope("update:post") {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message":"Insufficient scope."}`))
		return
	}

	vars := mux.Vars(r)
	postID, _ := strconv.Atoi(vars["id"])

	var post *Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post.Author = token.RegisteredClaims.Subject
	for _, item := range posts {
		if item.ID == postID {
			item.Timestamp = int(time.Now().Unix())
			item.Title = post.Title
			item.Content = post.Content
			item.Author = post.Author
			break
		}
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
