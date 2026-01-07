"use strict"

function createObject(elem) {
    let res;
    let isEmpty = true;
    let check;
    const fields = elem.querySelectorAll(":scope > [data-obj-field]");
    switch (elem.dataset.objType) {
        case "tempNode":
            const [value, check] = createObject(fields[0]);
            isEmpty = check;
            res = value;
            break;
        case "string":
            res = elem.value;
            if (res) {
                isEmpty = false;
            }
            break;
        case "array":
            res = [];
            for (const node of fields) {
                const [val, check] = createObject(node);
                if (!check) {
                    res.push(val);
                }
                isEmpty &&= check;
            }
            break;
        case "object":
            res = {};
            for (const node of fields) {
                const [val, check] = createObject(node);
                if (!check) {
                    res[node.dataset.objField] = val;
                }
                isEmpty &&= check;
            }
            break;
    }
    return [res, isEmpty];
}

console.log(223)
