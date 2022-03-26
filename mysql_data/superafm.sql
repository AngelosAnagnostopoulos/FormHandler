CREATE DATABASE golang_test;
USE golang_test;

CREATE TABLE Person (
    ID int NOT NULL PRIMARY KEY AUTO_INCREMENT, 
    fname varchar(255) NOT NULL,
    email varchar(255),
    phone varchar(255),
    city varchar(255)
);

INSERT INTO Person VALUES (1, "Kwstas", "test@gmail.com", "testPhone", "testCity")