package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type SiteData struct {
	StudSitePath        string // Ссылка на основной сайт
	Title               string // Заголовок вкладки
	MetaTitle           string // Мета-заголовок (Open Graph)
	Description         string // Описание сайта (Open Graph)
	VerificationName    string // Поле name для метатэга верификации в Вебмастере
	VerificationContent string // Поле content для метатэга верификации в Вебмастере
}

const StudSitePath = "stud.l9labs.ru"

var (
	VerificationName    string
	VerificationContent string
)

// Главная страница
func indexHandler(w http.ResponseWriter, _ *http.Request) {
	data := SiteData{
		StudSitePath:        StudSitePath,
		Title:               "l9_stud",
		MetaTitle:           "Учебная экосистема - l9labs",
		VerificationName:    VerificationName,
		VerificationContent: VerificationContent,
	}

	tmpl, _ := template.ParseFiles("templates/index.html")
	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

// Страница бота
func botHandler(w http.ResponseWriter, _ *http.Request) {
	data := SiteData{
		StudSitePath: StudSitePath,
		Title:        "l9_stud_bot",
		MetaTitle:    "Бот с раписанием занятий - l9labs",
		Description:  "Расписание занятий и уведомления о парах прямо в твоём мессенджере!",
	}

	tmpl, _ := template.ParseFiles("templates/bot.html")
	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found!")
	}
	var exists bool
	VerificationName, exists = os.LookupEnv("VERIF_NAME")
	if !exists {
		log.Fatal("lost env key: VERIF_NAME")
	}
	VerificationContent, exists = os.LookupEnv("VERIF_CONT")
	if !exists {
		log.Fatal("lost env key: VERIF_CONT")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/bot", botHandler)
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "media/favicon.ico")
	})

	server := &http.Server{
		Addr:         "localhost:5000",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 1 * time.Second, //10ms Redundant time
		IdleTimeout:  15 * time.Second,
	}

	fmt.Println("Server is listening...")
	log.Fatal(server.ListenAndServe())
}
