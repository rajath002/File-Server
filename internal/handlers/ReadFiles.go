package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	inlinetmpls "github.com/rajath002/File-Server/internal/handlers/inline-tmpls"
	"github.com/rajath002/File-Server/internal/models"
)

/*
ReadFilesHander reads the video files from the "videos" directory,
and displays them on the home page using a template file "home.page.tmpl"
*/
func ReadFilesHander(res http.ResponseWriter, req *http.Request) {
	// Read the video files from the "videos" directory
	files, err := os.ReadDir(".")
	if err != nil {
		http.Error(res, "Unable to read video directory", http.StatusInternalServerError)
		return
	}

	// Create a slice to hold the file names
	var videoFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".mp4" {
			videoFiles = append(videoFiles, file.Name())
		}
	}

	td := models.TemplateData{
		VideoFiles: videoFiles,
	}

	// Parse the template file
	t, err := template.ParseFiles("templates/home.page.tmpl")
	if err != nil {
		http.Error(res, "Unable to parse template", http.StatusInternalServerError)
		return
	}

	// Execute the template with the video files data
	err = t.Execute(res, td)
	if err != nil {
		http.Error(res, "Unable to generate HTML", http.StatusInternalServerError)
	}
}

func add(x, y int) int {
	return x + y
}

// check if the string ends with any of the supported extensions
func hasSupportedExtension(fileName string, _supportedExtensions []string) bool {
	if fileName == "" {
		return false
	}
	if strings.HasSuffix(fileName, "/") {
		return true
	}
	for _, extension := range _supportedExtensions {
		if strings.HasSuffix(strings.ToUpper(fileName), extension) {
			return true
		}
	}
	return false
}

// check if the string ends with any of the image extensions
func hasImageExtension(fileName string) bool {
	_supportedExtensions := []string{".JPG", ".JPEG", ".PNG"}
	res := hasSupportedExtension(fileName, _supportedExtensions)
	return res
}

/*
ReadFilesHanderFromRoot reads the video files from the root directory,
and displays them on the home page using an inline HTML template
*/
func ReadFilesHanderFromRoot(res http.ResponseWriter, req *http.Request) {
	// Read the video files from the "videos" directory
	files, err := os.ReadDir(".")
	if err != nil {
		http.Error(res, "Unable to read video directory", http.StatusInternalServerError)
		return
	}

	// Create a slice to hold the file names
	var videoFiles []string
	var supportedExtensions = []string{".MP4", ".MKV", ".JPG", ".JPEG", ".PNG"}
	for _, file := range files {
		if !file.IsDir() {
			ext := strings.ToUpper(filepath.Ext(file.Name()))
			for _, extension := range supportedExtensions {
				if ext == extension {
					videoFiles = append(videoFiles, file.Name())
					break
				}
			}
		}
	}

	td := models.TemplateData{
		VideoFiles: videoFiles,
	}

	// Parse the template with the custom function "add",
	// `inlinetmpls.ReadFilesTmpl` contains the HTML template
	t, err := template.New("videoList").Funcs(
		template.FuncMap{
			"add":               add,
			"hasImageExtension": hasImageExtension,
		},
	).Parse(inlinetmpls.ReadFilesTmpl)
	if err != nil {
		fmt.Println("Unable to parse template:", err)
		http.Error(res, "Unable to read video directory", http.StatusInternalServerError)
		return
	}

	// Execute the template with the video files data
	err = t.Execute(res, td)
	if err != nil {
		http.Error(res, "Unable to generate HTML", http.StatusInternalServerError)
	}
}
