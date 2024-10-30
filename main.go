package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rajath002/File-Server/internal/handlers"
	"github.com/rajath002/File-Server/internal/utils"
)

func main() {
	// Define the directory with the video files
	fs := http.FileServer(http.Dir("."))

	// Serve the files in the "videos" folder at the "/videos" route
	http.Handle("/videos/", http.StripPrefix("/videos/", fs))

	http.HandleFunc("/", handlers.ReadFilesHanderFromRoot)

	address, err := utils.GetWiFiIPv4Address()
	if err != nil {
		fmt.Println("Unable to get WiFi IP address:", err)
		return
	}

	// Define the address and port for the server to run on
	port := 8080
	fmt.Printf("You can access the server at \nhttp://%s:8080\n", address)
	fmt.Println("OR")
	fmt.Printf("http://localhost:%d\n", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
