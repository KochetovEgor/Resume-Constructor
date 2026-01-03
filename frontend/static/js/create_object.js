"use strict"

function createObject(elem) {
    let res;
    const fields = elem.querySelectorAll(":scope > [data-obj-field]");
    switch (elem.dataset.objType) {
        case "tempNode":
            res = createObject(fields[0]);
            break;
        case "string":
            res = elem.value;
            break;
        case "array":
            res = [];
            for (const node of fields) {
                res.push(createObject(node));
            }
            break;
        case "object":
            res = {};
            for (const node of fields) {
                res[node.dataset.objField] = createObject(node);
            }
            break;
    }
    return res;
}
