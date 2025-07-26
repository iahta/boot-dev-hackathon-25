package internal

import (
	"log"
	"net/http"

	"github.com/skip2/go-qrcode"
)

func (cfg *FileConfig) QRHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("QR Handler")
	qImg, err := qrcode.Encode(cfg.url, qrcode.Medium, 256)
	if err != nil {
		log.Fatalf("error generating qr code")
	}

	log.Printf("📱 QR code generated for: %s", cfg.url)

	w.Header().Set("Content-Type", "image/png")
	w.Write(qImg)

}
