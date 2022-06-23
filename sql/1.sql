DROP TABLE IF EXISTS textnote;

CREATE TABLE textnote (
    id bigserial NOT NULL,
    title varchar(255) NOT NULL,
    note text,
    finished boolean NOT NULL,
    due_date timestamp,
    CONSTRAINT textnote_pkc PRIMARY KEY(id)
);
