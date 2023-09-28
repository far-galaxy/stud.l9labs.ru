package main

import (
	"html/template"
	"log"
	"net/http"
)

// Главная страница
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
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
func BotHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
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
