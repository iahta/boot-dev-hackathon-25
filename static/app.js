
document.addEventListener("DOMContentLoaded", async () => {
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

    document.getElementById('uploadForm').addEventListener('submit', async (event) => {
        event.preventDefault();
        await uploadFiles();
    });
});

async function uploadFiles() {
    const uploadFile = document.getElementById("uploadFile");

    const formData = new FormData();
    for (const file of uploadFile.files) {
        formData.append("uploadFile", file);
    }

    try {
        const res = await fetch(window.location.origin + "/upload", {
            method: 'POST',
            body: formData,
        });
        if (!res.ok) {
            console.log("Upload Failed");
            throw new Error(`Faild to upload files: ${data.error}`);
        }
        console.log("Upload Successful!");
    } catch (error) {
        alert(`Error: ${error.message}`);
    };    
}
