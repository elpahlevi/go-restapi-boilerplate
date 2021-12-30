CREATE TABLE students (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO students(name)
VALUES
('James Damon'),
('Jon Snow'),
('Matt Solar');
