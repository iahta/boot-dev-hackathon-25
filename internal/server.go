package internal

import (
	"log"
	"net"
	"net/http"
	"os"
)

func Server(filePath string) {
	cfg := FileConfig{
		filePath: filePath,
	}

	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		log.Fatalf("error creating directory: %v", err)
	}

	mux := http.NewServeMux()
	appHandler := http.StripPrefix("/app", http.FileServer(http.Dir("/")))
	mux.Handle("/app/", appHandler)
	mux.HandleFunc("/", cfg.UploadFormHandler)
	mux.HandleFunc("/upload", cfg.UploadFileHandler)

	ipAddress := getLocalIP()
	srv := &http.Server{
		Addr:    ipAddress + ":8080",
		Handler: mux,
	}
	log.Printf("Server started on http://%s:8080", ipAddress)
	log.Fatal(srv.ListenAndServe())
}

func getLocalIP() string {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	return "localhost"
}
