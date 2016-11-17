CREATE TABLE users (
    id         varchar(36),
    email      varchar(255),
    password   varchar(64),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    role       integer,
    PRIMARY KEY(id),
    CONSTRAINT email UNIQUE(email)
);