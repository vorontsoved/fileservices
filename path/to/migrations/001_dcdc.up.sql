CREATE SCHEMA files;

CREATE TABLE  files.files (
                                           id SERIAL PRIMARY KEY,
                                           filename VARCHAR NOT NULL,
                                           created_at TIMESTAMP NOT NULL,
                                           modified_at TIMESTAMP NOT NULL
);
