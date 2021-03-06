package config

import(
	"net/http"
	// "path"
	"html/template"
)

func TablesHandler(w http.ResponseWriter,r *http.Request){
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
		"views/tables/tables.html",
	))

	var err = tmpl.ExecuteTemplate(w, "tables", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	

	// err = tmpl.Execute(w, data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	
}