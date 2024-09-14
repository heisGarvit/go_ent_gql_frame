package upload

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

var StaticDir = "/var/www/html"

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Parse the multipart form.
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error parsing form"))
		slog.Error("Error parsing form", "err", err)
		return
	}

	// Get the file from the form.
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error retrieving file"))
		slog.Error("Error retrieving file", "err", err)
		return
	}
	defer file.Close()

	fileName := "/u/" + uuid.New().String() + filepath.Ext(header.Filename)

	// Create a new file on the server.
	dst, err := os.Create(StaticDir + fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating file"))
		slog.Error("Error creating file", "err", err)
		return
	}
	defer dst.Close()

	// Copy the file to the server.
	_, err = io.Copy(dst, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error copying file"))
		slog.Error("Error copying file", "err", err)
	}

	jsonResponse := fmt.Sprintf(`{"url": "%s"}`, fileName)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonResponse))

}
