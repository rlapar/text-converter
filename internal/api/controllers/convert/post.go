package convert

import (
	"encoding/json"
	"github.com/golang/gddo/httputil/header"
	"net/http"
)

//TODO format enum

type convertRequest struct {
	SourceFormat string	//sourceFormat in request body
	TargetFormat string //targetFormat in request body
	Content string		//content in request body
}

type convertResponse struct {
	Content string `json:"content"`
}



func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// decode request
	if r.Header.Get("Content-Type") != "" {
		contentType, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if contentType != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
	// force the max size of the body in bytes
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	var requestBody convertRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Reqiest body contains badly-formatted JSON.", http.StatusBadRequest)
		return
	}


	// Implement logic


	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(convertResponse{Content: "Not implement"})
	if err != nil {
		return
	}





	//converter := converter.Converter{}
	//convertedContent, err := converter.ConvertPdf2Text(requestBody.Content)
	//
	//if err != nil {
	//	http.Error(w, "Invalid PDF content.", http.StatusBadRequest)
	//	return
	//}
	//
	w.WriteHeader(http.StatusOK)
	//err = json.NewEncoder(w).Encode(convertResponse{Content: convertedContent})
	err = json.NewEncoder(w).Encode(convertResponse{Content: "not implemented"})
	if err != nil {
		return
	}
}