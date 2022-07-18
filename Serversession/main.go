package main

import (
	"io"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Session struct {
	UserId int
	Lang   string
}

var Sessions map[string]*Session //map[token]

func NewSession() (token string) {
	token = RandStringRunes(16)
	for {
		if _, found := Sessions[token]; !found {
			break
		}
	}
	Sessions[token] = &Session{
		UserId: 0,
		Lang:   "hu",
	}
	return token
}

func ProcessSession(w http.ResponseWriter, r *http.Request) *Session {
	token := r.Header.Get("Token")
	if token == "" {
		token = NewSession()
	} else {
		if _, found := Sessions[token]; !found {
			token = NewSession()
		}
	}
	w.Header().Set("Token", token)
	return Sessions[token]

}

func Handler(w http.ResponseWriter, r *http.Request) {

	session := ProcessSession(w, r)

	switch session.Lang {
	case "hu":
		w.Write([]byte("Heló vagy valami"))
	case "en":
		w.Write([]byte("Hello or something"))
	default:
		w.Write([]byte("iNVALID LANGUAGE"))
	}
}

func HandlerChange(w http.ResponseWriter, r *http.Request) {
	session := ProcessSession(w, r)

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Can't read body("+err.Error()+")", http.StatusInternalServerError)
		return
	}

	bodyStr := string(body)
	switch bodyStr {
	case "hu", "en":
		session.Lang = bodyStr
	default:
		http.Error(w, "Invalid language code", http.StatusBadRequest)

	}

}

func main() {
	Sessions = make(map[string]*Session)

	http.HandleFunc("/", Handler)
	http.HandleFunc("/change", HandlerChange)

	http.ListenAndServe(":7777", nil)
}
