package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define the directory with the video files
	fs := http.FileServer(http.Dir("./videos"))

	// Define the route for the home page, serve the index.html file
	http.HandleFunc("/", RootHandler)

	// Serve the files in the "videos" folder at the "/videos" route
	http.Handle("/videos/", http.StripPrefix("/videos/", fs))

	// Define the address and port for the server to run on
	port := 8080
	fmt.Printf("Server started at http://localhost:%d/videos\n", port)

	// Start the server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func RootHandler(res http.ResponseWriter, req *http.Request) {
	file, _ := os.ReadFile("public/index.html")
	t := template.New("")
	t, _ = t.Parse(string(file))

	t.Execute(res, nil)
}
