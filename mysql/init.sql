-- DROP SCHEMA IF EXISTS tutorial_db;
-- CREATE SCHEMA tutorial_db;
USE tutorial_db;

DROP TABLE IF EXISTS book, employee, department;

CREATE TABLE book (
    id int AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL,
    value int NOT NULL
);

CREATE TABLE department (
    id int AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL
);


INSERT INTO department(name) VALUES ('営業部'), ('人事部'), ('企画部');

CREATE TABLE employee (
    id int AUTO_INCREMENT PRIMARY KEY,
    department_id int,
    income int,
    age int,
    gender int NOT NULL,
    name varchar(30),
    CONSTRAINT
        FOREIGN KEY (department_id)
        REFERENCES department (id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

INSERT INTO employee(department_id, income, age, gender, name)
    VALUES
    (1, 1000, 38, 1, '佐藤 A'),
    (1, 900,  27, 0, '佐藤 B'),
    (1, 800,  32, 0, '佐藤 C'),
    (1, 600,  24, 1, '佐藤 D'),
    (2, 1500, 50, 0, '佐藤 E'),
    (2, 400,  25, 1, '佐藤 F'),
    (2, 600,  29, 0, '佐藤 G'),
    (2, 1200, 42, 1, '佐藤 H'),
    (2, 700,  38, 0, '佐藤 I'),
    (2, 500,  26, 1, '佐藤 J'),
    (3, 100,  22, 1, '佐藤 K'),
    (3, 200,  24, 0, '佐藤 L'),
    (3, 150,  23, 1, '佐藤 M'),
    (3, 300,  24, 1, '佐藤 N'),
    (3, 400,  31, 0, '佐藤 O'),
    (3, 210,  26, 0, '佐藤 P'),
    (3, 220,  28, 1, '佐藤 Q'),
    (3, 420,  29, 1, '佐藤 R'),
    (3, 330,  26, 0, '佐藤 S'),
    (3, 280,  31, 1, '佐藤 T');