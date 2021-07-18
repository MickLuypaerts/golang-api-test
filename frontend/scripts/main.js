const baseUrl = "http://localhost:8080/products";

const table = document.querySelector("table");
const postDataForm  = document.getElementById("putDataForm");
const methodsSelector = document.getElementById("methodsSelector");

const idCon = document.getElementById("idCon");
const nameCon = document.getElementById("nameCon");
const descCon = document.getElementById("descCon");
const priceCon = document.getElementById("priceCon");
const skuCon = document.getElementById("skuCon");

const getInputs = new InputList([idCon], [nameCon, descCon, priceCon, skuCon]);
const getAllInputs = new InputList(null, [idCon, nameCon, descCon, priceCon, skuCon]);
const postInputs = new InputList([nameCon, descCon, priceCon, skuCon],[idCon]);
const putInputs = new InputList([idCon, nameCon, descCon, priceCon, skuCon], null);
const deleteInputs = new InputList([idCon], [nameCon, descCon, priceCon, skuCon]);

// TODO: handle 404 product not found error 
window.addEventListener("load", () => {
    sendToServer(baseUrl, "GET", null)
    .then(response => {
        idCon.style.visibility = "";
        nameCon.style.visibility = "hidden";
        descCon.style.visibility = "hidden";
        priceCon.style.visibility = "hidden";
        skuCon.style.visibility = "hidden";
        methodsSelector.value = "GET";
        if (response != null) {
            createFirstRow();
            response.map(prod => addProdToTable(prod));
        }
    });
});


methodsSelector.addEventListener("change", () => {
    switch(methodsSelector.value) {
        case "GET":
            getInputs.Enable();
            break;
        case "GET-all":
            getAllInputs.Enable();
            break;
        case "POST":
            postInputs.Enable();
            break;
        case "PUT":
            putInputs.Enable();
            break;
        case "DELETE":
            deleteInputs.Enable();
        default:
            console.log("default");
      }
})

postDataForm.addEventListener("submit", (event) => {
    event.preventDefault();
    let id = event.target[0].value;
    let data = null;

    switch(methodsSelector.value) {
        case "GET":
            sendToServer(baseUrl + "/" + id, "GET",data)
            .then(response => {
                if (response != null) {
                    clearTable();
                    addProdToTable(response);
                }
            });
            break;
        case "GET-all":
            sendToServer(baseUrl, "GET", null)
            .then(response => {
                if (response != null) {
                    clearTable();
                    response.map(prod => addProdToTable(prod));
                }
            });
            break;
        case "POST":
            data = postInputs.GetInputs();
            sendToServer(baseUrl, "POST", data)
            .then(response => {
                console.log(response)
            })
            break;
        case "PUT":
            data = putInputs.GetInputs();
            sendToServer(baseUrl + "/" + id, "PUT", data)
            .then(response => {
                clearTable();
                addProdToTable(response)  
            })
            break;
        case "DELETE":
            sendToServer(baseUrl + "/" + id, "DELETE", null)
            .then(response => {
                console.log(response);
            })
        default:
            console.log("default");
      }
})

function clearTable() {
    products = document.querySelectorAll("tr#prod");
    for (product of products) {
        product.remove();
    }
}

function createFirstRow() {
    let tableRow = document.createElement("tr");
    let rowsValue = Object.getOwnPropertyNames(new Product);
    for (value of rowsValue) {
        let row = document.createElement("th");
        row.textContent = value;
        tableRow.appendChild(row);
    }
    table.appendChild(tableRow);
}

function addProdToTable(product) {
    const prod = new Product(product.id, product.name, product.description, product.price, product.sku);
    let tableRow = prod.ToTableRow();
    table.appendChild(tableRow);
    

}

async function sendToServer(url, method, data) {
    try {
        const response = await fetch(url, {
            method: method,
            headers: {
                "Content-Type": "application/json",
            },
            body: data
        });
        if (response.status != 200 && response.status != 201 && response.status != 204) {
            throw response.statusText
        }
        try {
            const jsonData = await response.json();
            return jsonData
        }
        catch {
            return response
        }

    } catch (error) {
        console.error(error);
    }
}