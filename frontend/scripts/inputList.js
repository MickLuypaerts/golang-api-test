class InputList {
    constructor(show, hide) {
        this.show = show;
        this.hide = hide;
    }
    
    Enable() {  
        if (this.show != null) {
            for (let input of this.show) {
                if (input.style.visibility === "hidden") {
                    input.style.visibility = "";
                }
            }
        }
        if (this.hide != null) {
            for (let input of this.hide) {
                if (input.style.visibility === "") {
                    input.style.visibility = "hidden";
                }
            }
        }
    }
    GetInputs() {
        if (this.show != null) {
            let values = [];
            for (let input of this.show) {
                values.push(input.childNodes[3].value);
            }
            if (values.length == 4) {
                return JSON.stringify({"name": values[0], "description": values[1], "price": parseInt(values[2]), "sku": values[3]});
            } else if (values.length == 5) {
                return JSON.stringify({"name": values[1], "description": values[2], "price": parseInt(values[3]), "sku": values[4]});
            }
        }
    }
}