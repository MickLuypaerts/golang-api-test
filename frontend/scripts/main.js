
const table = document.querySelector("table");

class Product {
    constructor(id, name, desc, price, sku) {
        this.id = id;
        this.name = name;
        this.desc = desc;
        this. price = price;
        this.sku = sku;
    }

    ToTableRow() {
        let tableRow = document.createElement("tr");
        let keys = Object.keys(this);
        for (let key of keys) {
            let row = document.createElement("td");
            row.textContent = this[key];
            tableRow.appendChild(row);
        }
        return tableRow;
    }
}


window.addEventListener("load", () => {
    sendToServer("http://localhost:8080/products", "GET", null)
    .then(response => {
        if (response != null) {
            createFirstRow()
            response.map(prod => addProdToTable(prod));
        }
    });
});

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