CREATE TABLE music
(
    id              SERIAL UNIQUE,
    name            VARCHAR(255) NOT NULL UNIQUE,
    performer       TEXT,
    realisr_year    INTEGER,
    genre           VARCHAR(20)  NOT NULL
);