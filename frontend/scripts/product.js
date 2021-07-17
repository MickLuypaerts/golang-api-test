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
        tableRow.setAttribute("id", "prod");
        let keys = Object.keys(this);
        for (let key of keys) {
            let row = document.createElement("td");
            row.textContent = this[key];
            tableRow.appendChild(row);
        }
        return tableRow;
    }
}