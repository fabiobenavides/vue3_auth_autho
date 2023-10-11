package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func corsHandler() func(http.Handler) http.Handler {
	log.Println("registering cors middleware")

	credentials := handlers.AllowCredentials()
	headers := handlers.AllowedHeaders([]string{"Authorization", "X-Requested-With", "Content-Type", "Access-Control-Allow-Credentials"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	exposed := handlers.ExposedHeaders([]string{"Token"})

	return handlers.CORS(credentials, headers, origins, methods, exposed)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, X-Requested-With, Content-Type, Access-Control-Allow-Credentials, Token")
			w.Header().Set("Access-Control-Exposed-Headers", "Token, Authorization")
			w.Header().Set("Content-Type", "application/json")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func main() {
	var hostAddress string

	flag.StringVar(&hostAddress, "address", ":8080", "host address")

	flag.Parse()

	rtr := mux.NewRouter()
	rtr.Use(loggingMiddleware)

	rtr.HandleFunc("/posts", getPosts).Methods("GET")
	rtr.HandleFunc("/register", registerUser).Methods("POST")
	rtr.HandleFunc("/authenticate", authenticateUser).Methods("POST")

	authRtr := rtr.Methods(http.MethodPost).Subrouter()
	authRtr.HandleFunc("/posts", submitPost).Methods("POST")
	authRtr.HandleFunc("/refresh", refreshUser).Methods("POST")
	authRtr.Use(validateAuth)

	fmt.Printf("Blog server running at %s\n", hostAddress)

	loadUsers()
	loadPosts()

	log.Fatal(http.ListenAndServe(hostAddress, corsMiddleware(rtr)))
}

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondError(w http.ResponseWriter, code int, message string) {
	payload := map[string]string{"error": message}

	respondJSON(w, code, payload)
}
