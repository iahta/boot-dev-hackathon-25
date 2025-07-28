package internal

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func (cfg *FileConfig) UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB max
	if err != nil {
		http.Error(w, "File too big", http.StatusBadRequest)
		return
	}
	log.Printf("file path is: %s", cfg.filePath)

	for key := range r.MultipartForm.File {
		log.Printf("Found file key: %s", key)
	}

	files := r.MultipartForm.File["uploadFile"]
	if len(files) == 0 {
		log.Printf("No files received")
		http.Error(w, "No files", http.StatusBadRequest)
		return
	}

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			log.Printf("❌ Failed to open file: %s", err)
			http.Error(w, fmt.Sprintf("Failed to open file: Error: %s File: %s", err, file), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		dstPath := filepath.Join(cfg.filePath, fileHeader.Filename)
		dst, err := os.Create(dstPath)
		if err != nil {
			log.Printf("❌ Failed to create path: %s", err)
			http.Error(w, fmt.Sprintf("Failed to create destination path: %s", err), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			log.Printf("❌ Failed to copy file: %s", err)
			http.Error(w, fmt.Sprintf("Failed to copy file to destination: %s", err), http.StatusInternalServerError)
			return
		}
		log.Printf("File uploaded to %s", dstPath)
	}
	log.Printf("File uploaded is ok")
	w.WriteHeader(http.StatusOK)
}
