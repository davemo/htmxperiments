package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/flowchartsman/handlebars/v3"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:embed templates
var templates embed.FS

type User struct {
	gorm.Model
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Security struct {
	gorm.Model
	Name      string
	UserID    uint
	Amount    int
	Type      string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ExerciseRequest struct {
	gorm.Model
	SecurityID   uint
	Amount       int
	ExerciseDate time.Time
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type App struct {
	DB *gorm.DB
}

func loadTemplateFile(filename string) string {
	bytes, err := templates.ReadFile("templates/" + filename)
	if err != nil {
		panic(err.Error())
	}
	return string(bytes)
}

func pluralize(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

func currency(v float64) string {
	return fmt.Sprintf("$%.2f", v)
}

func formatDateRelative(date time.Time) string {
	now := time.Now()
	diff := now.Sub(date)

	if diff < time.Minute {
		return "just now"
	} else if diff < time.Hour {
		minutes := int(diff.Minutes())
		return fmt.Sprintf("%d minute%s ago", minutes, pluralize(minutes))
	} else if diff < time.Hour*24 {
		hours := int(diff.Hours())
		return fmt.Sprintf("%d hour%s ago", hours, pluralize(hours))
	} else if diff < time.Hour*48 {
		return "yesterday"
	} else {
		return date.Format("01/02/2006")
	}
}

func (app *App) index(w http.ResponseWriter, r *http.Request) {
	var securities []Security
	result := app.DB.Find(&securities)
	if result.Error != nil {
		log.Println(result.Error)
	}
	ctx := make(map[string]interface{})
	debug, err := json.MarshalIndent(securities, "", " ")
	if err != nil {
		panic(err.Error())
	}
	ctx["securities"] = securities
	ctx["debug"] = string(debug)

	content, err := handlebars.Render(loadTemplateFile("index.html"), ctx)
	if err != nil {
		panic(err.Error())
	}

	w.Write([]byte(content))
}

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	r := mux.NewRouter()
	app := &App{DB: db}

	handlebars.RegisterHelper("pluralize", pluralize)
	handlebars.RegisterHelper("formatDateRelative", formatDateRelative)
	handlebars.RegisterHelper("currency", currency)

	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("dist/"))))
	r.HandleFunc("/", app.index)
	r.Use(loggingMiddleware)

	log.Println("Server started on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
