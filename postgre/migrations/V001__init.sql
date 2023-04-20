CREATE SCHEMA IF NOT EXISTS olympics;

create table olympics.t_athletes
(
    name    TEXT    NOT NULL,
    age     INTEGER,
    country TEXT    NOT NULL,
    year    INTEGER NOT NULL,
    date    TEXT    NOT NULL,
    sport   TEXT    NOT NULL,
    gold    INTEGER NOT NULL,
    silver  INTEGER NOT NULL,
    bronze  INTEGER NOT NULL,

    PRIMARY KEY (name, sport, year)
);