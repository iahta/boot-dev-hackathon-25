package main

import (
	"github.com/iahta/boot-dev-hackathon-25/internal"
)

/*
func main() {
	err := os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		log.Fatalf("error creating directory: %v", err)
	}

	mux := http.NewServeMux()
	appHandler := http.StripPrefix("/app", http.FileServer(http.Dir("/")))
	mux.Handle("/app/", appHandler)
	mux.HandleFunc("/", handlers.UploadFormHandler)
	mux.HandleFunc("/upload", handlers.UploadFileHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Server started on http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
*/

func main() {
	filePath := internal.Setup()
	internal.Server(filePath)

}
