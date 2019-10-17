package config

import(
	"net/http"
	// "path"
	"html/template"
)

func HomeHandler(w http.ResponseWriter,r *http.Request){
	http.FileServer(http.Dir("assets"))
	// var filepath = path.Join("views/home", "home.html")
	// var tmpl, err = template.ParseFiles(filepath)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Kelompok 2",
	}

	var tmpl = template.Must(template.ParseFiles(
		"views/index/index.html",
		"views/tables.html",
		"views/calender.html",
		"views/typography.html",
	))

	var err = tmpl.ExecuteTemplate(w, "home", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	

	// err = tmpl.Execute(w, data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	
}