CREATE TABLE IF NOT exists products (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    price       DECIMAL(20,2) NOT NULL,
    SKU         VARCHAR(255) NOT NULL,
    createdOn   timestamp DEFAULT now(),
    updatedOn   timestamp DEFAULT now(),
    deletedOn   TIMESTAMP
);

INSERT INTO products (name, description, price, SKU)
VALUES 
    ('luxs', 'Dit is een amber kleurig bier zacht van afdronk.', 1.70, 'TODO'),
    ('Luxs Classics', 'TODO', 1.70, 'TODO');