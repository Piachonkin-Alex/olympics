CREATE TABLE olympics.t_role (
    name TEXT PRIMARY KEY,
    role INTEGER NOT NULL
);

INSERT INTO olympics.t_role(name, role) VALUES ('admin', '3');
