function cloneElementWithButton(event) {
    const forCloning = event.currentTarget.previousElementSibling
    const clone = forCloning.cloneNode(true)
    forCloning.after(clone)
}

const btnAddContatcs = document.getElementById("buttonAddContact")
const btnAddPosDesc = document.getElementById("buttonAddPositionDescription")
const btnAddCourse = document.getElementById("buttonAddCourse")
const btnAddProjectDecs = document.getElementById("buttonAddProjectDescription")
const btnAddProject = document.getElementById("buttonAddProject")
const btnAddSkill = document.getElementById("buttonAddSkill")

btnAddContatcs.addEventListener("click", cloneElementWithButton)
btnAddPosDesc.addEventListener("click", cloneElementWithButton)
btnAddCourse.addEventListener("click", cloneElementWithButton)
btnAddProjectDecs.addEventListener("click", cloneElementWithButton)
btnAddProject.addEventListener("click", cloneElementWithButton)
btnAddSkill.addEventListener("click", cloneElementWithButton)