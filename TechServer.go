package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type userHandlers struct {
	sync.Mutex
	store map[string]User
}

func newUserHandler() *userHandlers {
	return &userHandlers{
		store: map[string]User{},
	}
}

type Post struct {
	Id              string `json:"id"`
	Caption         string `json:"caption"`
	ImageURL        string `json:"img-url"`
	PostedTimestamp string `json:"time-stamp"`
}
type postHandlers struct {
	sync.Mutex
	store map[string]Post
}

func newPostHandler() *postHandlers {
	return &postHandlers{
		store: map[string]Post{},
	}
}

func (h *userHandlers) addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Write([]byte("post - adduser"))
	}
}
func (h *userHandlers) getUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("get - get user by id"))

		parts := strings.Split(r.URL.String(), "/")
		if len(parts) != 3 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		h.Lock()
		user, ok := h.store[parts[2]]
		h.Unlock()
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		jsonBytes, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)

	}
}
func (h *postHandlers) addPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Write([]byte("post - add post"))
	}
}
func (h *postHandlers) getPostById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("get - get post by id"))

		parts := strings.Split(r.URL.String(), "/")
		if len(parts) != 3 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		h.Lock()
		post, ok := h.store[parts[2]]
		h.Unlock()
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		jsonBytes, err := json.Marshal(post)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}
func (h *postHandlers) getPostByUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("get - get post by User"))

	}
}
func (h *userHandlers) userManager(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("get is called"))
		h.getUserById(w, r)
		return
	case "POST":
		w.Write([]byte("post is called"))
		h.addUser(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func main() {
	postHandlers := newPostHandler()
	userHandlers := newUserHandler()

	// Create an User POST request  ‘/users'
	http.HandleFunc("/users", userHandlers.userManager)

	// Get a user using id GET request ‘/users/<id here>’
	http.HandleFunc("/users/{id}", userHandlers.userManager)

	// Create a Post POST request ‘/posts'
	http.HandleFunc("/posts", postHandlers.addPost)

	// Get a post using id GET request ‘/posts/<id here>’
	http.HandleFunc("/posts/{id}", postHandlers.getPostById)

	// List all posts of a user GET request ‘/posts/users/<Id here>'
	http.HandleFunc("/posts/users/{id}", postHandlers.getPostByUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

/*









 */
