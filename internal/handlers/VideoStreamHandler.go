package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/rajath002/File-Server/internal/models"
	"github.com/rajath002/File-Server/internal/utils"
)

func StreamVideo(w http.ResponseWriter, r *http.Request) {
	// get the video file path from the request URL
	videoName := r.URL.Path[len("/video-stream/"):]

	fmt.Println("Video file path:", videoName)

	// Open the video file
	videoPath := fmt.Sprintf("./%s", videoName) // Path to your video file

	// check if the video file exists
	if _, err := os.Stat(videoPath); os.IsNotExist(err) {
		fmt.Println("Video file not found:", videoPath)
		http.Error(w, "Video not found.", http.StatusNotFound)
		return
	}

	file, err := os.Open(videoPath)
	if err != nil {
		fmt.Println("Error opening video file:", err)
		http.Error(w, "Video not found.", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Get the file information
	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file information:", err)
		http.Error(w, "Could not obtain video information.", http.StatusInternalServerError)
		return
	}

	// fmt.Println("Streaming video file:", videoPath)
	// Serve video content with range support
	http.ServeContent(w, r, filepath.Base(videoPath), fileStat.ModTime(), file)
}

// video player page
func VideoPlayerPage(res http.ResponseWriter, req *http.Request) {

	videoName := req.URL.Path[len("/video-stream/"):]
	fmt.Println("Video file path:", videoName)

	// Get the IP address of the WiFi interface
	address, err := utils.GetWiFiIPv4Address()
	if err != nil {
		fmt.Println("Unable to get WiFi IP address:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	td := models.Video{
		Path: fmt.Sprintf("http://%s:8080/video-stream/%s", address, videoName),
	}

	tpl := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Go Video Streaming</title>
		</head>
		<body>
			<h1>Go Video Streaming</h1>
			<video controls autoplay width="640" height="360">
				<source src="{{.Path}}" type="video/mp4">
				Your browser does not support the video tag.
			</video>
		</body>
		</html>
	`

	// Parse and execute the template
	t, err := template.New("videoList").Parse(tpl)
	if err != nil {
		fmt.Println("Unable to parse template:", err)
		http.Error(res, "Unable to read video directory", http.StatusInternalServerError)
		return
	}

	// Execute the template
	err = t.Execute(res, td)
	if err != nil {
		http.Error(res, "Unable to generate HTML", http.StatusInternalServerError)
	}

}
