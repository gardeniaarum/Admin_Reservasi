package main

import(
	"fmt"
	"net/http"
	i "reservasi_admin/src/controller/index"
	t "reservasi_admin/src/controller/tables"
	// k "reservasi_admin/src/controller/ketersediaan"
	// c "reservasi_admin/src/controller/comment"
	// conn "project_reservasi/src/config"
)

func main(){
	http.HandleFunc("/", i.indexHandler)
	http.HandleFunc("/index", i.indexHandler)
	http.HandleFunc("/tables", t.tablesHandler)

	http.Handle("/css/", 
	http.StripPrefix("/css/", 
		http.FileServer(http.Dir("assets/css"))))
	
	http.Handle("/js/", 
	http.StripPrefix("/js/", 
		http.FileServer(http.Dir("assets/js"))))
	
	http.Handle("/image/", 
	http.StripPrefix("/image/", 
		http.FileServer(http.Dir("assets/images"))))
	
	http.Handle("/icon/", 
	http.StripPrefix("/icon/", 
		http.FileServer(http.Dir("assets/icon"))))

	http.Handle("/vidio/", 
	http.StripPrefix("/vidio/", 
		http.FileServer(http.Dir("assets/vidio"))))

	fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}
