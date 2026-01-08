"use strict"

function saveResumeObject() {
    let [ResumeObject, temp] = createObject(resume);
    localStorage.setItem("ResumeObject", JSON.stringify(ResumeObject));
    console.log("Saved");
}

function getAddingButton(elem) {
    const child = elem.querySelector(":scope > [data-for-button-add]");
    const btnValue = child.dataset.forButtonAdd;
    return elem.querySelector(`:scope > [data-button-add="${btnValue}"]`);
}

function parseObjectIntoElement(elem, obj) {
    if (obj === undefined || obj === null) {
        return;
    }
    let res;
    switch (elem.dataset.objType) {
        case "tempNode":
            let elemTemp = elem.querySelector(":scope > [data-obj-field]");
            parseObjectIntoElement(elemTemp, obj);
        case "string":
            if (obj !== undefined) {
                elem.value = String(obj);
            }
            break;
        case "array":
            const btn = getAddingButton(elem);
            const ev = new MouseEvent("click", {cancelable: true});
            let nodes = elem.querySelectorAll(":scope > [data-obj-field]");
            for (let i = obj.length; i > nodes.length; i -= 1) {
                btn.dispatchEvent(ev);
            }
            nodes = elem.querySelectorAll(":scope > [data-obj-field]");
            for (let i = 0; i < nodes.length && i < obj.length; i += 1) {
                parseObjectIntoElement(nodes[i], obj[i]);
            }
            break;
        case "object":
            const fields = elem.querySelectorAll(":scope > [data-obj-field]");
            for (const node of fields) {
                parseObjectIntoElement(node, obj[node.dataset.objField]);
            }
            break;
    }
}

function loadResumeButton() {
    const resumeObj = JSON.parse(localStorage.getItem("ResumeObject"));
    parseObjectIntoElement(resume, resumeObj);
    console.log("Loaded");
}

const buttonSaveResume = document.getElementById("buttonSaveResume");
const buttonLoadResume = document.getElementById("buttonLoadResume");

buttonSaveResume.addEventListener("click", saveResumeObject);
buttonLoadResume.addEventListener("click", loadResumeButton);
