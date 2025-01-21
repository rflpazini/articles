CREATE TABLE catalog
(
    id    SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    type  VARCHAR(50)  NOT NULL
);

INSERT INTO catalog (title, type)
VALUES ('Filme A', 'Movie'),
       ('Série B', 'TV Show'),
       ('Documentário C', 'Documentary');