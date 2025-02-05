package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	inlinetmpls "github.com/rajath002/File-Server/internal/handlers/inline-tmpls"
	"github.com/rajath002/File-Server/internal/models"
	"github.com/rajath002/File-Server/internal/utils"
)

func StreamVideo(w http.ResponseWriter, r *http.Request) {
	// get the media file path from the request URL
	mediaName := r.URL.Path[len("/video-stream/"):]

	fmt.Println("Media file path:", mediaName)

	// Open the media file
	mediaPath := fmt.Sprintf("./%s", mediaName) // Path to your media file

	// check if the media file exists
	if _, err := os.Stat(mediaPath); os.IsNotExist(err) {
		fmt.Println("Media file not found:", mediaPath)
		http.Error(w, "Media not found.", http.StatusNotFound)
		return
	}

	file, err := os.Open(mediaPath)
	if err != nil {
		fmt.Println("Error opening media file:", err)
		http.Error(w, "Media not found.", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Get the file information
	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file information:", err)
		http.Error(w, "Could not obtain media information.", http.StatusInternalServerError)
		return
	}

	// Serve media content with range support
	http.ServeContent(w, r, filepath.Base(mediaPath), fileStat.ModTime(), file)
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

	// Parse and execute the template
	t, err := template.New("videoList").Parse(inlinetmpls.VideoPlayerTmpl)
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
