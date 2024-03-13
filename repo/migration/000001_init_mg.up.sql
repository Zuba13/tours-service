CREATE TABLE IF NOT EXISTS tours (
    id          INTEGER PRIMARY KEY,
    author_id   INTEGER NOT NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    price       NUMERIC(10, 2) NOT NULL,
    difficult INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS checkpoints (
    id          INTEGER PRIMARY KEY,
    tour_id     INTEGER NOT NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    latitude    DECIMAL(9, 6) NOT NULL,
    longitude   DECIMAL(9, 6) NOT NULL,
    CONSTRAINT fk_tour
        FOREIGN KEY (tour_id)
        REFERENCES tours (id)
        ON DELETE CASCADE
);