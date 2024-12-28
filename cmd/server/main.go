package main
import (
	"fmt"
	"html/template"
	"net/http"
)



func main() {
	// focus on cto, cio, architect, tech lead, senior dev, dev
	// golang + ai
	// teaching + courses
	// golang saas boilerplate
	// podcast like htmx-pod (small bits of go)
	// Handle static files (CSS, JS, images, etc.)
	// Serve static files from the "public" directory

	mux := http.NewServeMux()

	// Serve static files (CSS, JS, images, etc.) from the 'public' directory
	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	mux.Handle("GET /images/", http.StripPrefix("/images/", http.FileServer(http.Dir("public/images/"))))
	mux.Handle("GET /styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("public/styles/"))))

	mux.HandleFunc("GET /", handleRoot)
	mux.HandleFunc("GET /ai", handleGoLangAI)
	mux.HandleFunc("GET /health", handleHealth)


	http.ListenAndServe(":8080", mux)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Health check")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func handleGoLangAI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to Golang AI")
	fmt.Println("AI Golang Libraries / Content and more")
	// this is an endpoint everything around golang + ai
	// list of libraries, projects, articles, etc
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
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
