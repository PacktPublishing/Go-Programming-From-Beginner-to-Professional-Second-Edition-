package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFileUpload_ServeHTTP(t *testing.T) {
	server := server{}

	fileDataBuffer := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&fileDataBuffer)

	file, err := os.Open("./go.mod")
	if err != nil {
		log.Fatal(err)
	}

	formFile, err := multipartWriter.CreateFormFile("myFile", file.Name())
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		log.Fatal(err)
	}

	multipartWriter.Close()

	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(fileDataBuffer.Bytes()))

	r.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	w := httptest.NewRecorder()

	server.ServeHTTP(w, r)
	if w.Body.String() != "go.mod Uploaded!" {
		t.Errorf("Expected 'go.mod Uploaded!' string but received: '%s'", w.Body.String())
	}
}
