
let selectedFiles = [];

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

    document.getElementById('uploadFile').addEventListener('change', () => {
        previewFiles();
    })
});

function previewFiles() {
    const fileInput = document.getElementById("uploadFile");
    const previewContainer = document.getElementById("preview");

    previewContainer.innerHTML = "";
    selectedFiles = [];
    
    Array.from(fileInput.files).forEach((file, index) => {
        const fileId = `file-${Date.now()}-${index}`;
        selectedFiles.push({ file, name: file.name, id: fileId });

        const reader = new FileReader();
        reader.onload = (e) => {
            const container = document.createElement("div");
            container.className = "file-preview";
            container.style.border = "1px solid #ccc";
            container.style.padding = "10px";
            container.style.textAlign = "center";
            container.style.maxWidth = "150px";

            const img = document.createElement("img");
            img.src = e.target.result;
            img.alt = file.name;
            img.style.maxWidth = "100%";
            img.style.height = "auto";

            const nameText = document.createElement("p");
            nameText.textContent = file.name;
            nameText.id = `name-${fileId}`;

            const renameBtn = document.createElement("button");
            renameBtn.textContent = "Rename"
            renameBtn.onclick = () => {
                const newName = prompt("New filename (without extension):", file.name.split(".")[0]);
                if (newName) {
                    const ext = file.name.split(".").pop();
                    const updatedName = `${newName}.${ext}`;
                    nameText.textContent = updatedName;

                    const f = selectedFiles.find(f => f.id === fileId);
                    if (f) f.name = updatedName;
                }
            };

            const removeBtn = document.createElement("button");
            removeBtn.textContent = "Remove";
            removeBtn.onclick = () => {
                selectedFiles = selectedFiles.filter(f => f.id !== fileId);
                container.remove();
            };

            container.appendChild(img);
            container.appendChild(nameText);
            container.appendChild(renameBtn);
            container.appendChild(removeBtn);
            previewContainer.appendChild(container);
        };
        reader.readAsDataURL(file);
    });
}

async function uploadFiles() {
    const uploadStatus = document.getElementById("uploadStatus");
    uploadStatus.innerHTML = "";

    if (selectedFiles.length === 0) {
        alert("no files to upload");
        return;
    }

    const formData = new FormData();
    selectedFiles.forEach(f => {
        const renamedFile = new File([f.file], f.name, { type: f.file.type });
        formData.append("uploadFile", renamedFile);
    });

    const progressBar = document.createElement("progress");
    progressBar.max = 100;
    progressBar.value = 0;
    uploadStatus.appendChild(progressBar);

    try {
        const res = await fetch(window.location.origin + "/upload", {
            method: 'POST',
            body: formData,
        });
        if (!res.ok) {
            console.log("Upload Failed");
            throw new Error(`Faild to upload files: ${data.error}`);
        }

        let val = 0;
        const interval = setInterval(() => {
            if (val < 90) {
                val += 5;
                progressBar.value = val;
            } else {
                clearInterval(interval);
            }
        }, 100);

        const text = await res.text();
        progressBar.value = 100;
        uploadStatus.innerHTML = "Upload Complete!"
    } catch (error) {
        uploadStatus.innerHTML = `Upload Failed: ${error.message}`;
    }
}
