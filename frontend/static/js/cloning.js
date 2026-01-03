"use strict"
function cloneElementWithButton(event) {
    const key = event.currentTarget.dataset.buttonAdd;
    const clone = indexCloning[key].cloneNode(true);
    addEventsOnButtons(clone);
    event.currentTarget.before(clone);
    event.preventDefault();
}

function addEventsOnButtons(node) {
    const addButtons = node.querySelectorAll("[data-button-add]");
    for (const btn of addButtons) {
        btn.addEventListener("click", cloneElementWithButton);
    }
}

const resume = document.getElementById("resume");

const indexCloning = {};
resume.querySelectorAll("[data-for-button-add]").forEach(
    el => {
        const key = el.dataset.forButtonAdd;
        indexCloning[key] = el.cloneNode(true);
    }
)

addEventsOnButtons(resume);
