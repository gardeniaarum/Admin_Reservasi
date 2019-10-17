package main

import(
	"fmt"
	"net/http"
	i "reservasi_admin/src/controller/index"
	t "reservasi_admin/src/controller/tables"
	c "reservasi_admin/src/controller/calendar"
)

func main(){
	http.HandleFunc("/", i.indexHandler)
	http.HandleFunc("/index", i.indexHandler)
	http.HandleFunc("/tables", t.tablesHandler)
	http.HandleFunc("/calendar", c.calendarHandler)

	http.Handle("/css/", 
	http.StripPrefix("/css/", 
		http.FileServer(http.Dir("assets/css"))))
	
	http.Handle("/js/", 
	http.StripPrefix("/js/", 
		http.FileServer(http.Dir("assets/js"))))
	
	http.Handle("/image/", 
	http.StripPrefix("/image/", 
		http.FileServer(http.Dir("assets/images"))))

	fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}
