package routes

import (
    "log"
    "net/http"
    "../actions"

    "github.com/gorilla/mux"
)

func Run () {

    r := mux.NewRouter()

    r.HandleFunc("/", actions.Index).Methods("GET")
    r.HandleFunc("/about", actions.About)
    r.HandleFunc("/create", actions.Create)
    r.HandleFunc("/articles", actions.Articles).Methods("GET")
    r.HandleFunc("/post/{id:[0-9]+}", actions.PostDetail)
    r.HandleFunc("/save_article", actions.SaveArticle).Methods("POST")
    r.HandleFunc("/method", actions.Method)
    r.HandleFunc("/contacts", actions.Contacts)
    r.HandleFunc("/api", actions.Api)
    http.Handle("/", r)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

    log.Fatal(http.ListenAndServe(":9999", nil))
}