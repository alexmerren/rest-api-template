CREATE DATABASE IF NOT EXISTS contacts;

use contacts;

DROP TABLE IF EXISTS contact;

CREATE TABLE IF NOT EXISTS contact (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255),
    phonenumber varchar(255)
);

INSERT INTO contact (id, name, email, phonenumber) VALUES (
    1,
    "Alex Merren",
    "alexandermerren@gmail.com",
    ""
);
