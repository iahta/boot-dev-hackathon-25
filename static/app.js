
document.addEventListener("DOMContentLoaded", function () {
    
    const button = document.getElementById("qrbutton");
    const container = document.getElementById("qrimage");
    if (button && container) {
        button.addEventListener("click", () => {
            console.log("qr button pressed")
            const img = document.createElement("img");
            img.src = "/qrcode";
            img.alt = "QR Code";

            img.onload = () => {
                container.innerHTML = "";
                container.appendChild(img);
            };

            img.onerror = () => {
                console.error("Failed to load image from /qrcode")
            }
        });
    } else {
        console.warn("qr button not found")
    }
});
