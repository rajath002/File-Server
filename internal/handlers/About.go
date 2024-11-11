package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	inlinetmpls "github.com/rajath002/File-Server/internal/handlers/inline-tmpls"
	"github.com/rajath002/File-Server/internal/models"
	"github.com/rajath002/File-Server/internal/utils"
)

func AboutHandlerPage(res http.ResponseWriter, req *http.Request) {
	// Get the IP address of the WiFi interface
	address, err := utils.GetWiFiIPv4Address()
	if err != nil {
		fmt.Println("Unable to get WiFi IP address:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	td := models.About{
		IPAddress: address,
	}

	// Parse and execute the template, inlinetmpls.AboutTmpl is the template.
	t, err := template.New("videoList").Parse(inlinetmpls.AboutTmpl)
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
