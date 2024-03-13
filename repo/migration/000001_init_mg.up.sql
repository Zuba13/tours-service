CREATE TABLE tours (
    id          UUID PRIMARY KEY,
    author_id   UUID NOT NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    price       NUMERIC(10, 2) NOT NULL,
    difficult   INTEGER NOT NULL
);
