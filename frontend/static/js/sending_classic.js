"use strict"

/*function debugObject() {
    console.log(createObject(resume));
}*/

async function sendResume() {
    const response = await fetch("http://localhost:8100/resume", {
        method: "POST",
        body: JSON.stringify(createObject(resume))
    })
    if (response.ok) {
        console.log("ok")
    } else {
        console.log("not ok")
    }
    return response
}

function previewPDF(response) {
    console.log(response)
}

async function buttonGeneratePDF() {
    const response = await sendResume()
    const PDFblob = await response.blob()
    console.log(PDFblob)
    pdfViewNode.src = URL.createObjectURL(PDFblob)
}

const pdfViewNode = document.getElementById("pdfViewNode")

buttonSendResume.addEventListener("click", buttonGeneratePDF);
console.log(25)
//buttonSendResume.addEventListener("click", debugObject);