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

func (app *App) search(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	var securities []Security
	result := app.DB.Where("name LIKE ?", "%"+query+"%").Find(&securities)
	if result.Error != nil {
		log.Println(result.Error)
	}

	for _, security := range securities {
		fmt.Fprintf(w, `
			<li class="group flex cursor-default select-none items-center rounded-md px-3 py-2">
				<svg class="h-6 w-6 flex-none text-gray-500" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
					stroke="currentColor" aria-hidden="true">
					<path stroke-linecap="round" stroke-linejoin="round"
						d="M2.25 12.75V12A2.25 2.25 0 014.5 9.75h15A2.25 2.25 0 0121.75 12v.75m-8.69-6.44l-2.12-2.12a1.5 1.5 0 00-1.061-.44H4.5A2.25 2.25 0 002.25 6v12a2.25 2.25 0 002.25 2.25h15A2.25 2.25 0 0021.75 18V9a2.25 2.25 0 00-2.25-2.25h-5.379a1.5 1.5 0 01-1.06-.44z" />
				</svg>
				<span class="ml-3 flex-auto truncate">%s</span>
			</li>
		`, security.Name)
	}
	if len(securities) == 0 {
		fmt.Fprintf(w, `<div class="px-6 py-14 text-center sm:px-14"><svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true" class="mx-auto h-6 w-6 text-gray-900 text-opacity-40"><path stroke-linecap="round" stroke-linejoin="round" d="M2.25 12.75V12A2.25 2.25 0 014.5 9.75h15A2.25 2.25 0 0121.75 12v.75m-8.69-6.44l-2.12-2.12a1.5 1.5 0 00-1.061-.44H4.5A2.25 2.25 0 002.25 6v12a2.25 2.25 0 002.25 2.25h15A2.25 2.25 0 0021.75 18V9a2.25 2.25 0 00-2.25-2.25h-5.379a1.5 1.5 0 01-1.06-.44z"></path></svg><p class="mt-4 text-sm text-gray-900">We couldn't find any results that matched your search.</p></div>`)
	}
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
	r.HandleFunc("/search", app.search)
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
