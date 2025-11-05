// Create a simple web server with multiple routes
func main() {
    // Handle static files
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Home page handler
    http.HandleFunc("/", homeHandler)
    
    // API endpoint
    http.HandleFunc("/api/data", apiHandler)
    
    fmt.Println("Server starting on :8080")
    http.ListenAndServe(":8080", nil)
}

// homeHandler serves the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    fmt.Fprintf(w, "<h1>Welcome to my Go website!</h1>")
}

// apiHandler returns JSON data
func apiHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"message": "Hello from Go API", "status": "success"}`)
}
