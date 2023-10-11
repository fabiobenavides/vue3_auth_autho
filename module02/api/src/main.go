package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func corsHandler() func(http.Handler) http.Handler {
	credentials := handlers.AllowCredentials()
	headers := handlers.AllowedHeaders([]string{"Authorization", "X-Requested-With", "Content-Type", "Access-Control-Allow-Credentials"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"})
	exposed := handlers.ExposedHeaders([]string{"Authorization"})

	return handlers.CORS(credentials, headers, origins, methods, exposed)
}

func main() {
	var hostAddress string

	flag.StringVar(&hostAddress, "address", ":8080", "host address")

	flag.Parse()

	rtr := mux.NewRouter()
	rtr.Use(corsHandler())

	rtr.HandleFunc("/posts", getPosts).Methods("GET")

	fmt.Printf("Blog server running at %s\n", hostAddress)

	log.Fatal(http.ListenAndServe(hostAddress, rtr))
}

type Post struct {
	ID        int    `json:"id"`
	Timestamp int    `json:"timestamp"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	posts := []*Post{
		{
			ID:        1,
			Timestamp: 1692376226,
			Title:     "Demo Post 01",
			Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			Author:    "",
		},
		{
			ID:        2,
			Timestamp: 1692376226,
			Title:     "Demo Post 02",
			Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			Author:    "",
		},
		{
			ID:        3,
			Timestamp: 1692376226,
			Title:     "Demo Post 03",
			Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			Author:    "",
		},
	}

	respondJSON(w, http.StatusOK, posts)
}

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
