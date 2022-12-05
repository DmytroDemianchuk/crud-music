CREATE TABLE music
(
    id              SERIAL UNIQUE,
    name            VARCHAR(255) NOT NULL UNIQUE,
    performer       VARCHAR(255),
    realise_year    INTEGER,
    genre           VARCHAR(20)  NOT NULL
);