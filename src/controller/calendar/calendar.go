package config

import(
	"net/http"
	// "path"
	"html/template"
)

func CalendarHandler(w http.ResponseWriter,r *http.Request){
	http.FileServer(http.Dir("assets"))

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Kelompok 2",
	}

	var tmpl = template.Must(template.ParseFiles(
		"views/calendar/calendar.html",
	))

	var err = tmpl.ExecuteTemplate(w, "calendar", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
}