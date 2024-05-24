package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var (
	templates = template.Must(template.ParseFiles("index.html", "register.html", "login.html"))
	db        *sql.DB
)

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	fmt.Println("Sunucu başlatılıyor....\nhttp://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func createTable() {
	createUsersTableSQL := `CREATE TABLE IF NOT EXISTS users (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "username" TEXT NOT NULL UNIQUE,
        "password" TEXT NOT NULL
    );`

	statement, err := db.Prepare(createUsersTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		_, err := db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, password)
		if err != nil {
			http.Error(w, "Kayıt başarısız: "+err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if err := templates.ExecuteTemplate(w, "register.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		var storedPassword string
		err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
		if err != nil {
			http.Error(w, "Giriş başarısız: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if password == storedPassword {
			fmt.Fprintf(w, "Giriş başarılı! Hoşgeldiniz, %s", username)
		} else {
			http.Error(w, "Geçersiz kullanıcı adı veya şifre", http.StatusUnauthorized)
		}
		return
	}
	if err := templates.ExecuteTemplate(w, "login.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
