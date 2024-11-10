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
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
			<style>
				.main-bg {
					background-image: 
						url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 1440 320'%3E%3Cpath fill='%230099ff' fill-opacity='1' d='M0,288L80,250.7C160,213,320,139,480,96C640,53,800,43,960,64C1120,85,1280,139,1360,165.3L1440,192L1440,0L1360,0C1280,0,1120,0,960,0C800,0,640,0,480,0C320,0,160,0,80,0L0,0Z'%3E%3C/path%3E%3C/svg%3E");
					background-size: cover;
					background-repeat: no-repeat;
					background-position: bottom;
				}
				video {
					border-radius: 20px;
				}
				.video-title {
					color: #0a58ca;
					border-radius: 10px;
					padding: 10px 20px; /* Add some padding for spacing */
					background: rgba(255, 255, 255, 0.2); /* Semi-transparent background */
					backdrop-filter: blur(10px); /* Adjust the blur intensity */
					border-radius: 8px; /* Optional rounded corners */
					box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1); /* Optional shadow for depth */
				}
			</style>
			</head>
				<body class="bg-light main-bg d-flex justify-content-center align-items-center vh-100">
					<div class="container text-center">
						<h1 class="display-6 mb-4 video-title">
							Go Video Streaming
						</h1>
						
						<div class="">
							<div class="ratio mb-3" style="--bs-aspect-ratio: 45%;">
								<video controls playsinline>
									<source src="{{.Path}}" type="video/mp4" >
									Your browser does not support the video tag.
								</video>
							</div>
						</div>
					</div>

			<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous" defer></script>
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
