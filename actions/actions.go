package actions

import (
    "fmt"
    "net/http"
    "html/template"
    "database/sql"
    //"../config"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)

type User struct {
    Name string
    Age uint16
    Money int16
    Avg_grades float64
    Happiness float64
    Hobbies []string
}

type Article struct {
    Id uint16
    Title, Anons, Full_text string
}

var posts = []Article{}

func Index (w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
    if err != nil { fmt.Fprintf(w, err.Error()) }
    currentUser := User{"Bob", 25, -50, 4.2, 0.8, []string {"Beer","Eat","Sleep"}}
    tmpl.ExecuteTemplate(w, "index", currentUser)
}

func Articles (w http.ResponseWriter, r *http.Request) {

    tmpl, err := template.ParseFiles("templates/articles.html", "templates/header.html", "templates/footer.html")
    if err != nil { panic(err) }

    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/rcmoney")
    if err != nil { panic(err) }
    defer db.Close()
    //fmt.Fprintf(w, "MethodPut")

    res, err := db.Query("SELECT * FROM `go_article`")
    if err != nil { panic(err) }
    defer res.Close()
    //fmt.Fprintf(w, "MethodPut")

    posts = []Article{}
    for res.Next() {
        var post Article
        err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Full_text)
        if err != nil { panic(err) }
        posts = append(posts, post)
    }
    tmpl.ExecuteTemplate(w, "articles", posts)
}

func PostDetail (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    //w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "ID: %v\n", vars["id"])
}

func Create (w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
    if err != nil { fmt.Fprintf(w, err.Error()) }
    tmpl.ExecuteTemplate(w, "create", nil)
    //fmt.Sprintf(config.Getdb())
}

func SaveArticle (w http.ResponseWriter, r *http.Request) {
    title := r.FormValue("title")
    anons := r.FormValue("anons")
    full_text := r.FormValue("full_text")

    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/rcmoney")
    if err != nil { panic(err) }
    defer db.Close()

    insert, err := db.Query(fmt.Sprintf("INSERT INTO `go_article` (`title`, `anons`, `full_text`) VALUES('%s', '%s', '%s')", title, anons, full_text))
    if err != nil { panic(err) }
    defer insert.Close()
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Method (w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:
            fmt.Fprintf(w, "MethodGet")
        case http.MethodPost:
            fmt.Fprintf(w, "MethodPost")
        case http.MethodPut:
            fmt.Fprintf(w, "MethodPut")
        case http.MethodDelete:
            fmt.Fprintf(w, "MethodDelete")
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func About (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "about")
}

func Contacts (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "contacts")
}

func Api (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "api")
}
