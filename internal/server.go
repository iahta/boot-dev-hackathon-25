package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Server(filePath string) {
	cfg := FileConfig{
		filePath: filePath,
	}
	fmt.Printf("configure file path is: %s", cfg.filePath)
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		log.Fatalf("error creating directory: %v", err)
	}

	mux := http.NewServeMux()
	appHandler := http.StripPrefix("/app", http.FileServer(http.Dir("/")))
	mux.Handle("/app/", appHandler)
	mux.HandleFunc("/", cfg.UploadFormHandler)
	mux.HandleFunc("/upload", cfg.UploadFileHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Server started on http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
