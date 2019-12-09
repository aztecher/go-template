CREATE DATABASE test_template;
USE test_template;

CREATE TABLE person(id INT AUTO_INCREMENT NOT NULL PRIMARY KEY, name VARCHAR(20) NOT NULL);
INSERT INTO person(name) values ("tanaka");
INSERT INTO person(name) values ("tokoro");
INSERT INTO person(name) values ("shimabara");
