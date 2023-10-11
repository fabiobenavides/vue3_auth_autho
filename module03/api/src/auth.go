package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var signingKey = []byte("sample-secret")

func generateToken(userID int) (token string) {
	expiryTime := time.Now().Add(2 * time.Hour)

	regClaims := jwt.RegisteredClaims{
		ID:        strconv.Itoa(userID),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expiryTime),
	}

	userJwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, regClaims)
	token, _ = userJwtToken.SignedString(signingKey)

	return token
}

type Auth struct {
	FullName string `json:"fullName"`
	Token    string `json:"token"`
}

func authenticateUser(w http.ResponseWriter, r *http.Request) {
	userName, password, ok := r.BasicAuth()
	if ok {
		for _, user := range users {
			if (user.UserName == userName) && (user.Password == password) {
				token := generateToken(user.ID)

				auth := Auth{
					FullName: user.FullName,
					Token:    token,
				}

				respondJSON(w, http.StatusOK, auth)
				return
			}
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return
}

func refreshUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.Header.Get("UserId"))
	if err == nil {
		for _, user := range users {
			if user.ID == userID {
				token := generateToken(user.ID)

				auth := Auth{
					FullName: user.FullName,
					Token:    token,
				}

				respondJSON(w, http.StatusOK, auth)
				return
			}
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return
}

func validateAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Token")

		if header == "" {
			respondError(w, http.StatusUnauthorized, "No Token Found")
			return
		}

		var regClaims jwt.RegisteredClaims

		token, err := jwt.ParseWithClaims(header, &regClaims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				msg := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				return nil, msg
			}

			return signingKey, nil
		})

		if err == nil {
			if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
				r.Header.Set("UserId", claims.ID)
				next.ServeHTTP(w, r)
				return
			}
		}

		respondError(w, http.StatusUnauthorized, "Invalid Token")
		return
	})
}
