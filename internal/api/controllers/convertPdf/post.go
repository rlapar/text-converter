package convertPdf

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/gddo/httputil/header"
	"io/ioutil"
	"net/http"
	"text-converter/internal/cfg"
	"text-converter/internal/converter"
)

type convertRequest struct {
	Content string		//content in request body
}

type convertResponse struct {
	Content string `json:"content"`
}



func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// force the max size of the body in bytes
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	//var []byte content
	var content []byte

	// decode request
	if r.Header.Get("Content-Type") != "" {
		contentType, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if contentType == "application/json" {
			cfg.Logger.Info("Content-Type: application/json")
			var requestBody convertRequest
			decoder := json.NewDecoder(r.Body)
			decoder.DisallowUnknownFields()
			err := decoder.Decode(&requestBody)
			if err != nil {
				cfg.Logger.Warning("Badly formatted JSON request body")
				http.Error(w, "Request body contains badly-formatted JSON", http.StatusBadRequest)
				return
			}
			content, err = base64.StdEncoding.DecodeString(requestBody.Content)
			if err != nil {
				http.Error(w, "Content is not a base64 encoded string", http.StatusBadRequest)
				return
			}
		} else if contentType == "application/octet-stream" {
			var err error
			cfg.Logger.Info("Content-Type: application/octet-stream")
			content, err = ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Can't read binary data", http.StatusBadRequest)
				return
			}
		} else {
			cfg.Logger.Warning(
				fmt.Sprintf("Unsupported Content-Type %s", contentType),
			)
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	converter := converter.Converter{}
	convertedContent, err := converter.ConvertPdf2Text(content)
	if err != nil {
		cfg.Logger.Warning("Invalid PDF content")
		http.Error(w, "Invalid PDF content.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(convertResponse{Content: convertedContent})
	if err != nil {
		return
	}
}