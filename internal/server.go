package internal

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func Server(filePath string) {
	ipAddress := getLocalIP()
	url := fmt.Sprintf("http://%s:8080", ipAddress)
	cfg := FileConfig{
		filePath: filePath,
		url:      url,
	}

	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		log.Fatalf("error creating directory: %v", err)
	}

	mux := http.NewServeMux()
	appHandler := http.StripPrefix("/app", http.FileServer(http.Dir("/")))
	staticHanlder := http.StripPrefix("/static", http.FileServer(http.Dir("static")))
	mux.Handle("/static/", staticHanlder)
	mux.Handle("/app/", appHandler)
	mux.HandleFunc("/upload", cfg.UploadFileHandler)
	mux.HandleFunc("/qrcode", cfg.QRHandler)
	mux.HandleFunc("/", cfg.UploadFormHandler)

	srv := &http.Server{
		Addr:    ipAddress + ":8080",
		Handler: mux,
	}
	log.Printf("Server started on %s", url)
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
