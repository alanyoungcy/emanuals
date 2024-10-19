package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Recevied request for path: ", r.URL.Path)
	//fmt.Printf("Serving: %s\n", r.URL.Path)
	//fmt.Fprintf(w, "<h1>Hello, World! My first go server</h1>")
	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Println("Error parsing template: ", err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, " %v Server error\n", http.StatusNotFound)
		fmt.Fprintf(w, "Error parsing template: %s\n", err)
		return
	}
	pages, _ := scandir("./manuals")
	fmt.Println(pages)
	t.Execute(w, pages)

}

func main() {
	fmt.Println("Hello, World!")

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/manuals/", http.StripPrefix("/manuals/", http.FileServer(http.Dir("manuals"))))
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
