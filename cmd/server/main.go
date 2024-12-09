package main
import (
	"fmt"
	"html/template"
	"net/http"
)




func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleRoot)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Health check")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})


	http.ListenAndServe("localhost:8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
    // Parse the HTML template
    tmpl, err := template.ParseFiles("public/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Execute the template
    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
